package validation

import (
	"github.com/muktiarafi/myriadcode-backend/internal/models"
	"testing"
)

func TestValidateCreateUser(t *testing.T) {

	t.Run("form data with empty value", func(t *testing.T) {
		register := models.RegisterUser{
			Name:     "",
			Nickname: "",
			Password: "",
		}

		if err := ValidateCreateUser(&register); err == nil {
			t.Error("Should get error but got none")
		}
	})

	t.Run("form data with short password field", func(t *testing.T) {
		register := models.RegisterUser{
			Name:     "Paijo",
			Nickname: "Yolo",
			Password: "1234",
		}

		if err := ValidateCreateUser(&register); err == nil {
			t.Error("Should get error but got none")
		}
	})

	t.Run("form data with valid values", func(t *testing.T) {
		register := models.RegisterUser{
			Name:     "Paijo",
			Nickname: "Yolo",
			Password: "12345678",
		}

		if err := ValidateCreateUser(&register); err != nil {
			t.Error("Should not get error but got one")
		}
	})
}
