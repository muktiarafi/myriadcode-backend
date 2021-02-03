package main

import (
	"github.com/muktiarafi/myriadcode-backend/internal/configs"
	"github.com/muktiarafi/myriadcode-backend/internal/logs"
	"github.com/muktiarafi/myriadcode-backend/internal/router"
	"net/http"
)


func main() {
	app := configs.NewAppConfig()

	logs.SetupLogs(app)

	r := router.SetRouter()

	http.ListenAndServe(":8000", r)
}
