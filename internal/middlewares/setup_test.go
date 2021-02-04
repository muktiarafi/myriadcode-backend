package middlewares

import (
	"github.com/go-chi/chi"
	"github.com/muktiarafi/myriadcode-backend/internal/configs"
	"github.com/muktiarafi/myriadcode-backend/internal/helpers"
	"github.com/muktiarafi/myriadcode-backend/internal/logs"
	"net/http"
	"os"
	"testing"
)

var mux *chi.Mux

const withoutImageResponseMessage = "Missing image name context"


func TestMain(m *testing.M) {
	app := configs.NewAppConfig()
	app.Logger = logs.NewLogger()
	helpers.NewHelper(app)

	saveFileDir = "."

	mux = chi.NewRouter()
	mux.With(ImageUpload).Post("/image", testUploadImageHandler)
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
