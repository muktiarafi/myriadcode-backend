package service

import (
	"database/sql"
	"errors"
	"github.com/muktiarafi/myriadcode-backend/internal/apierror"
	"github.com/muktiarafi/myriadcode-backend/internal/models"
	"github.com/muktiarafi/myriadcode-backend/internal/repository"
	"github.com/muktiarafi/myriadcode-backend/internal/validation"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: *userRepository,
	}
}

func (ur *UserServiceImpl) CreateUser(userPostData *models.RegisterUser, imagePath string) (*models.CurrentUser, error) {
	err := validation.ValidateCreateUser(userPostData)
	if err != nil {
		return nil, apierror.NewBadRequestError(err, err.Error())
	}

	user, err := ur.userRepository.FindUserByNickname(userPostData.Nickname)
	if err != sql.ErrNoRows {
		return nil, err
	}
	if len(user.Nickname) != 0 {
		return nil, errors.New("nickname already taken")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userPostData.Password), 12)
	if err != nil {
		return nil, err
	}

	userPostData.Password = string(hash)
	currentUser, err := ur.userRepository.CreateUser(userPostData, imagePath)
	if err != nil {
		return nil, err
	}

	return currentUser, nil
}

func (ur *UserServiceImpl) UpdateUser(user *models.CurrentUser) (*models.CurrentUser, error) {
	//stmt := `UPDATE users
	//SET name = $1, image_path = $2
	//WHERE id = $3
	//RETURNING id, name, image_path, nickname, is_admin, joined_at`
	//
	//var currentUser models.CurrentUser
	//err := ur.DB.SQL.QueryRow(
	//	stmt,
	//	user.Name, user.ImagePath, user.ID).
	//	Scan(
	//		&currentUser.ID,
	//		&currentUser.Name,
	//		&currentUser.ImagePath,
	//		&currentUser.Nickname,
	//		&currentUser.IsAdmin,
	//		&currentUser.JoinedAt)
	//if err != nil {
	//	return nil, err
	//}

	return &models.CurrentUser{}, nil
}
