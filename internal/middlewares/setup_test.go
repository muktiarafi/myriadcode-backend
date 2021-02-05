package middlewares

import (
	"github.com/go-chi/chi"
	"github.com/muktiarafi/myriadcode-backend/internal/configs"
	"github.com/muktiarafi/myriadcode-backend/internal/helpers"
	"github.com/muktiarafi/myriadcode-backend/internal/logs"
	"github.com/muktiarafi/myriadcode-backend/internal/models"
	"net/http"
	"os"
	"testing"
	"time"
)

var mux *chi.Mux

const withoutImageResponseMessage = "Missing image name context"
const testFileDir = "../../static/test-file"

var payload = models.UserPayload{
	ID:       1,
	Nickname: "paimin",
	IsAdmin:  false,
}

func TestMain(m *testing.M) {
	app := configs.NewAppConfig()
	app.Logger = logs.NewLogger()
	helpers.NewHelper(app)

	saveFileDir = testFileDir
	mux = chi.NewRouter()
	mux.With(ImageUpload).Post("/image", testUploadImageHandler)
	mux.Post("/cookie", setCookieHandler)
	mux.With(RequireAuth).Post("/auth", requireAuthHandler)
	code := m.Run()

	os.Exit(code)
}

func testUploadImageHandler(w http.ResponseWriter, r *http.Request) {
	fileName, ok := r.Context().Value("image").(string)
	if !ok {
		w.Write([]byte(withoutImageResponseMessage))
		return
	}

	w.Write([]byte(fileName))
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {

	token, err := helpers.CreateTokenWithExpire(&payload, time.Now().Add(time.Minute).Unix())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	cookie := http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: false,
		Name: "session",
		Value: token,
		Expires: time.Now().Add(time.Minute),
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}

func requireAuthHandler(w http.ResponseWriter, r *http.Request) {
	payload, ok := r.Context().Value("user").(*models.UserPayload)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	helpers.SendJSON(w, http.StatusOK, payload)
}
