package service

import "github.com/muktiarafi/myriadcode-backend/internal/models"

type UserService interface {
	CreateUser(userPostData *models.RegisterUser, imagePath string) (*models.CurrentUser, error)
	UpdateUser(user *models.UserPayload, userName, imageName string) (*models.CurrentUser, error)
	Authenticate(loginRequest *models.LoginRequest) (string, error)
}
