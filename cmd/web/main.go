package main

import (
	"github.com/joho/godotenv"
	"github.com/muktiarafi/myriadcode-backend/internal/configs"
	"github.com/muktiarafi/myriadcode-backend/internal/driver"
	"github.com/muktiarafi/myriadcode-backend/internal/handler"
	"github.com/muktiarafi/myriadcode-backend/internal/helpers"
	"github.com/muktiarafi/myriadcode-backend/internal/logs"
	"github.com/muktiarafi/myriadcode-backend/internal/repository"
	"github.com/muktiarafi/myriadcode-backend/internal/router"
	"github.com/muktiarafi/myriadcode-backend/internal/service"
	"log"
	"net/http"
)

const portNumber = ":8000"

func main() {
	app := configs.NewAppConfig()
	app.Logger = logs.NewLogger()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := driver.ConnectSQL(
		configs.PostgresDSN(),
		app,
	)
	if err != nil {
		panic(err)
	}
	defer db.SQL.Close()

	r := router.SetRouter()

	helpers.NewHelper(app)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(&userRepository)
	userHandler := handler.NewUserHandler(&userService)
	userHandler.Route(r)

	log.Println("connected to port 8000")
	http.ListenAndServe(portNumber, r)
}
