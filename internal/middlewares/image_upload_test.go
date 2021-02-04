package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/muktiarafi/myriadcode-backend/internal/helpers"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestImageUpload(t *testing.T) {
	const testFileDir = "../../static/test-file"
	t.Run("with image", func(t *testing.T) {
		file, err := os.Open(testFileDir + "/gambar.png")
		if err != nil {
			t.Error(err)
		}
		defer file.Close()

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", filepath.Base("gambar.png"))
		if err != nil {
			t.Error(err)
		}
		_, err = io.Copy(part, file)
		writer.Close()

		request := httptest.NewRequest(http.MethodPost, "/image", body)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		mux.ServeHTTP(response, request)

		responseBody, _ := ioutil.ReadAll(response.Body)

		responseFileName := strings.Split(string(responseBody), "-")
		want := "gambar.png"
		got := responseFileName[len(responseFileName) - 1]

		if want != got {
			t.Errorf("Expecting to get filename %q, but got %q instead", want, got)
		}
	})

	t.Run("without image", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		writer.WriteField("a", "a")
		writer.Close()

		request := httptest.NewRequest(http.MethodPost, "/image", body)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		mux.ServeHTTP(response, request)

		responseBody, _ := ioutil.ReadAll(response.Body)

		want := withoutImageResponseMessage
		got := string(responseBody)

		if want != got {
			t.Errorf("Expecting to get response message %q, but got %q instead", want, got)
		}
	})

	t.Run("with unsupported image format", func(*testing.T) {
		file, err := os.Open(testFileDir + "/test.txt")
		if err != nil {
			t.Error(err)
		}
		defer file.Close()

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", filepath.Base("test.txt"))
		if err != nil {
			t.Error(err)
		}
		_, err = io.Copy(part, file)
		writer.Close()

		request := httptest.NewRequest(http.MethodPost, "/image", body)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		mux.ServeHTTP(response, request)

		responseBody, _ := ioutil.ReadAll(response.Body)

		var apiResponse helpers.ErrorResponse
		json.Unmarshal(responseBody, &apiResponse)

		want := http.StatusBadRequest
		got := apiResponse.Status

		if want != got {
			t.Errorf("Expecting to get status %d, but got %d instead", want, got)
		}
	})
}
