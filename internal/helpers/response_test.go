package helpers

import (
	"encoding/json"
	"errors"
	"github.com/muktiarafi/myriadcode-backend/internal/apierror"
	"github.com/muktiarafi/myriadcode-backend/internal/configs"
	"github.com/muktiarafi/myriadcode-backend/internal/logs"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendJSON(t *testing.T) {
	a := struct {
		Name string `json:"name"`
	}{"Udin"}
	response := httptest.NewRecorder()

	SendJSON(response, http.StatusOK, a)

	responseBody, _ := ioutil.ReadAll(response.Body)
	apiResponse := SuccessResponse{}
	json.Unmarshal(responseBody, &apiResponse)

	want := http.StatusOK
	got := apiResponse.Status

	if want != got {
		t.Errorf("Expected %d status code, got %d status code instead", want, got)
	}
}

func TestSendError(t *testing.T) {
	a := configs.NewAppConfig()
	a.Logger = logs.NewLogger()

	NewHelper(a)

	t.Run("standard error", func(t *testing.T) {
		err := errors.New("fwerwer")
		response := httptest.NewRecorder()

		SendError(response, err)

		responseBody, _ := ioutil.ReadAll(response.Body)
		apiResponse := ErrorResponse{}
		json.Unmarshal(responseBody, &apiResponse)

		want := http.StatusText(http.StatusInternalServerError)
		got := apiResponse.Error

		if got != want {
			t.Errorf("expected %q, got %q instead", want, got)
		}
	})

	t.Run("ApiError type", func(t *testing.T) {
		err := errors.New("invalid")
		response := httptest.NewRecorder()

		const warning = "Invalid Validation"

		SendError(response, apierror.NewBadRequestError(err, warning))

		responseBody, _ := ioutil.ReadAll(response.Body)
		apiResponse := ErrorResponse{}
		json.Unmarshal(responseBody, &apiResponse)

		want := warning
		got := apiResponse.Error

		if got != want {
			t.Errorf("expected %q, got %q instead", want, got)
		}
	})
}
