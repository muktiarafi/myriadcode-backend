package repository

import "github.com/muktiarafi/myriadcode-backend/internal/models"

type UserRepository interface {
	FindCurrentUserByID(id int) (*models.CurrentUser, error)
	FindUserByNickname(nickname string) (*models.User, error)
	CreateUser(userPostData *models.RegisterUser, imagePath string) (*models.CurrentUser, error)
	UpdateUser(user *models.CurrentUser) (*models.CurrentUser, error)
}
