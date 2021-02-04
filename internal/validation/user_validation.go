package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/muktiarafi/myriadcode-backend/internal/models"
)

func ValidateCreateUser(model *models.RegisterUser) error {
	err := validation.ValidateStruct(model,
		validation.Field(&model.Name, validation.Required, validation.Length(4, 45)),
		validation.Field(&model.Nickname, validation.Required, validation.Length(1, 45)),
		validation.Field(&model.Password, validation.Required, validation.Length(8, 0)))

	return err
}
