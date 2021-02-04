package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/muktiarafi/myriadcode-backend/internal/helpers"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_CreateUser(t *testing.T) {

	t.Run("Send valid registration data to the handler", func(t *testing.T) {
		formData := map[string]string{
			"name":     "paijo",
			"nickname": "bambank",
			"password": "12345678",
		}

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

		apiResponse := helpers.SuccessResponse{}
		json.Unmarshal(responseBody, &apiResponse)

		got := apiResponse.Status
		want := http.StatusOK

		if got != want {
			t.Errorf("Expected status code %d got %d instead", want, got)
		}
	})

	t.Run("Send invalid registration data to the user", func(t *testing.T) {
		formData := map[string]string{
			"name":     "ab",
			"nickname": "bambank",
			"password": "12345678",
		}

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

		apiResponse := helpers.ErrorResponse{}
		json.Unmarshal(responseBody, &apiResponse)

		got := apiResponse.Status
		want := http.StatusBadRequest

		fmt.Println(apiResponse.Error)

		if got != want {
			t.Errorf("Expected status code %d got %d instead", want, got)
		}
	})
}
