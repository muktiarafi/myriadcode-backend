package handler

import (
	"encoding/json"
	"fmt"
	"github.com/muktiarafi/myriadcode-backend/internal/helpers"
	"net/http"
	"testing"
)

func TestUserHandler_CreateUser(t *testing.T) {

	t.Run("Send valid registration data to the handler", func(t *testing.T) {
		formData := map[string]string{
			"name":     "paijo",
			"nickname": "bambank",
			"password": "12345678",
		}

		responseBody := createUser(formData)

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

		responseBody := createUser(formData)

		apiResponse := helpers.ErrorResponse{}
		json.Unmarshal(responseBody, &apiResponse)

		got := apiResponse.Status
		want := http.StatusBadRequest

		fmt.Println(apiResponse.Error)

		if got != want {
			t.Errorf("Expected status code %d got %d instead", want, got)
		}
	})

	t.Run("Register with duplicate nickname", func(t *testing.T) {
		formData := map[string]string{
			"name":     "abcd",
			"nickname": "paijo",
			"password": "12345678",
		}

		responseBody := createUser(formData)

		apiResponse := helpers.SuccessResponse{}
		json.Unmarshal(responseBody, &apiResponse)

		if apiResponse.Status != http.StatusOK {
			t.Errorf("Expected %d status code, got %d instead", http.StatusOK, apiResponse.Status)
		}

		responseBody = createUser(formData)

		errorResponse := helpers.ErrorResponse{}
		json.Unmarshal(responseBody, &errorResponse)

		if errorResponse.Status != http.StatusBadRequest {
			t.Errorf("Expected %d status code, got %d instead", http.StatusOK, errorResponse.Status)
		}

		want := "nickname already taken"
		got := errorResponse.Error

		if got != want {
			t.Errorf("Expected %q, but got %q instead", want, got)
		}
	})
}
