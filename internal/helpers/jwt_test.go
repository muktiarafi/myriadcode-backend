package helpers

import (
	"github.com/muktiarafi/myriadcode-backend/internal/models"
	"testing"
)

func TestToken(t *testing.T) {
	jwtKey = "example"

	userPayload := models.UserPayload{
		ID:       1,
		Nickname: "bambank",
		IsAdmin:  false,
	}

	tokenString, err := CreateToken(&userPayload)
	if err != nil {
		t.Error(err)
	}

	token, payload, err := ParseToken(tokenString)
	if err != nil {
		t.Error(err)
	}

	if !token.Valid {
		t.Error("Invalid token")
	}

	want := userPayload.Nickname
	got := payload.Nickname

	if got != want {
		t.Errorf("Expected nickname to be %q, but got %q instead", want, got)
	}
}
