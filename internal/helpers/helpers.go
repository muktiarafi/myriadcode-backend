package helpers

import (
	"github.com/muktiarafi/myriadcode-backend/internal/configs"
	"os"
)

var app *configs.AppConfig

func NewHelper(a *configs.AppConfig) {
	key := os.Getenv("JWT_KEY")
	if len(key) != 0 {
		jwtKey = key
	}

	app = a
}

