package handler

import (
	"bytes"
	"encoding/json"
	"github.com/muktiarafi/myriadcode-backend/internal/helpers"
	"github.com/muktiarafi/myriadcode-backend/internal/models"
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

func TestUserHandler_UpdateUser(t *testing.T) {
	t.Run("update user name", func(t *testing.T) {
		formData := map[string]string{
			"name": "udin",
			"nickname": "ud",
			"password": "12345678",
		}

		createUser(formData)

		response, _ := login(&models.LoginRequest{
			Nickname: formData["nickname"],
			Password: formData["password"],
		})

		assertResponseCode(t, response.Result().StatusCode, http.StatusOK)

		cookies := response.Result().Cookies()
		updateUserData := map[string]string{
			"name": "Udin U besar",
		}

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		for k, v := range updateUserData {
			writer.WriteField(k, v)
		}
		writer.Close()

		request := httptest.NewRequest(http.MethodPut, "/users/update", body)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		request.AddCookie(cookies[0])
		response = httptest.NewRecorder()

		mux.ServeHTTP(response, request)

		responseBody, _ := ioutil.ReadAll(response.Body)

		apiResponse := struct {
			Data models.CurrentUser `json:"data"`
		}{models.CurrentUser{}}
		json.Unmarshal(responseBody, &apiResponse)

		assertResponseCode(t, response.Result().StatusCode, http.StatusOK)

		got := apiResponse.Data.Name
		want := updateUserData["name"]

		if got != want {
			t.Errorf("Expected name to be changed to %q, but got %q instead", want, got)
		}
	})

	t.Run("update user image", func(t *testing.T) {
		formData := map[string]string{
			"name": "paijo",
			"nickname": "pai",
			"password": "12345678",
		}

		responseBody := createUser(formData)

		apiResponse := struct {
			Data models.CurrentUser `json:"data"`
		}{models.CurrentUser{}}
		json.Unmarshal(responseBody, &apiResponse)

		got := apiResponse.Data.ImagePath
		want := "anonim.jpg"

		assertImageName(t, got, want)

		response, _ := login(&models.LoginRequest{
			Nickname: formData["nickname"],
			Password: formData["password"],
		})

		assertResponseCode(t, response.Result().StatusCode, http.StatusOK)

		cookies := response.Result().Cookies()
		const fileName = "gambar.png"
		file, err := os.Open(testFileDir + "/" + fileName)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", filepath.Base(fileName))
		if err != nil {
			t.Error(err)
		}
		_, err = io.Copy(part, file)
		writer.Close()

		request := httptest.NewRequest(http.MethodPut, "/users/update", body)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		request.AddCookie(cookies[0])
		response = httptest.NewRecorder()

		mux.ServeHTTP(response, request)

		assertResponseCode(t, response.Result().StatusCode, http.StatusOK)

		responseBody, _ = ioutil.ReadAll(response.Body)

		apiResponse = struct {
			Data models.CurrentUser `json:"data"`
		}{models.CurrentUser{}}
		json.Unmarshal(responseBody, &apiResponse)

		responseFileName := strings.Split(apiResponse.Data.ImagePath, "-")

		got = responseFileName[len(responseFileName) - 1]
		want = fileName

		assertImageName(t, got, want)
	})
}

func assertImageName(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Expected image path to be %q, but got %q instead", want, got)
	}
}
