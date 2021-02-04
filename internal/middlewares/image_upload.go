package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/muktiarafi/myriadcode-backend/internal/apierror"
	"github.com/muktiarafi/myriadcode-backend/internal/helpers"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var saveFileDir = "./static/images"

func ImageUpload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.ParseMultipartForm(10 << 20)

		// not handling error because uploading image is optional
		file, handler, _ := r.FormFile("image")
		if file != nil {
			defer file.Close()
			format := strings.Split(handler.Filename, ".")
			if !isAllowedFormat(format[len(format)-1]) {
				helpers.SendError(w, apierror.NewBadRequestError(
					errors.New("uploading unsupported file format"),
					"File format not supported"))
				return
			}
			f, err := os.OpenFile(
				fmt.Sprintf("%s/%d-%s", saveFileDir, time.Now().Unix(), handler.Filename),
				os.O_WRONLY|os.O_CREATE,
				0666)
			if err != nil {
				helpers.SendError(w, err)
				return
			}
			defer f.Close()

			_, err = io.Copy(f, file)
			if err != nil {
				helpers.SendError(w, err)
				return
			}

			ctx := context.WithValue(r.Context(), "image", f.Name())
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}

	})
}

func isAllowedFormat(format string) bool {
	formats := map[string]struct{}{
		"jpg":  {},
		"png":  {},
		"jpeg": {},
	}

	_, ok := formats[format]

	return ok
}
