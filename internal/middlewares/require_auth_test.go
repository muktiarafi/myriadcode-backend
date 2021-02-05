package middlewares

import (
	"encoding/json"
	"github.com/muktiarafi/myriadcode-backend/internal/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireAuth(t *testing.T) {
	t.Run("request with cookie", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/cookie", nil)
		response := httptest.NewRecorder()

		mux.ServeHTTP(response, request)

		cookies := response.Result().Cookies()
		if len(cookies) == 0 {
			t.Error("Should get cookie but get none")
		}

		request = httptest.NewRequest(http.MethodPost, "/auth", nil)
		request.AddCookie(cookies[0])
		response = httptest.NewRecorder()

		mux.ServeHTTP(response, request)

		responseBody, _ := ioutil.ReadAll(response.Body)

		responsePayload := struct {
			Data models.UserPayload `json:"data"`
		}{models.UserPayload{}}
		json.Unmarshal(responseBody, &responsePayload)

		want := payload.Nickname
		got := responsePayload.Data.Nickname

		if got != want {
			t.Errorf("Expected nickname %q, but got %q instead", got, want)
		}
	})

	t.Run("request without cookie", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/auth", nil)
		response := httptest.NewRecorder()

		mux.ServeHTTP(response, request)

		got := response.Result().StatusCode
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("Expected to get status code %d, but got %d instead", got, want)
		}
	})
}
