package handler

import (
	"github.com/go-chi/chi"
	"github.com/muktiarafi/myriadcode-backend/internal/helpers"
	"github.com/muktiarafi/myriadcode-backend/internal/middlewares"
	"github.com/muktiarafi/myriadcode-backend/internal/models"
	"github.com/muktiarafi/myriadcode-backend/internal/service"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: *userService,
	}
}

func (uh *UserHandler) Route(mux *chi.Mux) {
	mux.Route("/users", func(r chi.Router) {
		r.With(middlewares.ImageUpload).Post("/register", uh.CreateUser)
	})
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	register := &models.RegisterUser{
		Name:     r.FormValue("name"),
		Nickname: r.FormValue("nickname"),
		Password: r.FormValue("password"),
	}

	defaultImageName := "anonim.jpg"
	imageName, ok := r.Context().Value("image").(string)
	if ok {
		defaultImageName = imageName
	}

	currentUser, err := uh.userService.CreateUser(register, defaultImageName)
	if err != nil {
		helpers.SendError(w, err)
		return
	}

	helpers.SendJSON(w, http.StatusOK, currentUser)
}
