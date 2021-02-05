package handler

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/muktiarafi/myriadcode-backend/internal/configs"
	"github.com/muktiarafi/myriadcode-backend/internal/driver"
	"github.com/muktiarafi/myriadcode-backend/internal/helpers"
	"github.com/muktiarafi/myriadcode-backend/internal/logs"
	"github.com/muktiarafi/myriadcode-backend/internal/repository"
	"github.com/muktiarafi/myriadcode-backend/internal/router"
	"github.com/muktiarafi/myriadcode-backend/internal/service"
	"github.com/ory/dockertest/v3"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	pool     *dockertest.Pool
	resource *dockertest.Resource
)

var mux *chi.Mux

func TestMain(m *testing.M) {
	app := configs.NewAppConfig()
	app.Logger = logs.NewLogger()

	mux = router.SetRouter()

	db := driver.DB{
		SQL: SetTestDatabase(),
	}

	helpers.NewHelper(app)
	userRepository := repository.NewUserRepository(&db)
	userService := service.NewUserService(&userRepository)
	userHandler := NewUserHandler(&userService)
	userHandler.Route(mux)

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func SetTestDatabase() *sql.DB {
	var err error
	pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err = pool.Run("postgres", "latest", []string{"POSTGRES_PASSWORD=secret", "POSTGRES_DB=postgres"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	var db *sql.DB
	if err = pool.Retry(func() error {
		db, err = sql.Open(
			"pgx",
			fmt.Sprintf("host=localhost port=%s dbname=postgres user=postgres password=secret", resource.GetPort("5432/tcp")))
		if err != nil {
			return err
		}

		migrationFilePath := filepath.Join("..", "..", "database", "migrations")
		return driver.Migration(migrationFilePath, db)
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return db
}

func createUser(formData map[string]string) []byte {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for k, v := range formData {
		writer.WriteField(k, v)
	}
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "/users/register", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	response := httptest.NewRecorder()

	mux.ServeHTTP(response, request)

	responseBody, _ := ioutil.ReadAll(response.Body)

	return responseBody
}
