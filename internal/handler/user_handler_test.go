package handler

import (
	"encoding/json"
	"github.com/muktiarafi/myriadcode-backend/internal/helpers"
	"github.com/muktiarafi/myriadcode-backend/internal/models"
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

		assertResponseCode(t, apiResponse.Status, http.StatusOK)
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

		assertResponseCode(t, apiResponse.Status, http.StatusBadRequest)
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

		assertResponseCode(t, apiResponse.Status, http.StatusOK)

		responseBody = createUser(formData)

		errorResponse := helpers.ErrorResponse{}
		json.Unmarshal(responseBody, &errorResponse)

		assertResponseCode(t, errorResponse.Status, http.StatusBadRequest)

		want := "nickname already taken"
		got := errorResponse.Error

		if got != want {
			t.Errorf("Expected %q, but got %q instead", want, got)
		}
	})
}

func TestUserHandler_Authenticate(t *testing.T) {
	formData := map[string]string{
		"name":     "abcd",
		"nickname": "paimin",
		"password": "12345678",
	}

	createUser(formData)

	t.Run("login with already created user", func(t *testing.T) {
		loginRequest := models.LoginRequest{
			Nickname: formData["nickname"],
			Password: formData["password"],
		}

		response, responseBody := login(&loginRequest)

		apiResponse := helpers.SuccessResponse{}
		json.Unmarshal(responseBody, &apiResponse)

		assertResponseCode(t, apiResponse.Status, http.StatusOK)

		if len(response.Result().Cookies()) == 0 {
			t.Error("Should get cookie but get none")
		}
	})

	t.Run("login with not exist user", func(t *testing.T) {
		loginRequest := models.LoginRequest{
			Nickname: "budi",
			Password: formData["password"],
		}

		_, responseBody := login(&loginRequest)

		apiResponse := helpers.ErrorResponse{}
		json.Unmarshal(responseBody, &apiResponse)

		assertResponseCode(t, apiResponse.Status, http.StatusBadRequest)
	})

	t.Run("login with not valid password", func(t *testing.T) {
		loginRequest := models.LoginRequest{
			Nickname: formData["nickname"],
			Password: "12345",
		}

		_, responseBody := login(&loginRequest)

		apiResponse := helpers.ErrorResponse{}
		json.Unmarshal(responseBody, &apiResponse)

		assertResponseCode(t, apiResponse.Status, http.StatusBadRequest)
	})
}
