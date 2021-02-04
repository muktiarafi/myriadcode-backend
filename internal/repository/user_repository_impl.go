package repository

import (
	"github.com/muktiarafi/myriadcode-backend/internal/driver"
	"github.com/muktiarafi/myriadcode-backend/internal/models"
)

func NewUserRepository(db *driver.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

type UserRepositoryImpl struct {
	DB *driver.DB
}

func (ur *UserRepositoryImpl) FindCurrentUserByID(id int) (*models.CurrentUser, error) {
	ctx, cancel := newDBContext()
	defer cancel()

	stmt := `SELECT id, name, nickname, image_path, is_admin, joined_at
	FROM users
	WHERE id = $1`

	var currentUser models.CurrentUser
	var isAdmin []uint8
	err := ur.DB.SQL.QueryRowContext(ctx, stmt, id).Scan(
		&currentUser.ID,
		&currentUser.Name,
		&currentUser.Nickname,
		&currentUser.ImagePath,
		&isAdmin,
		&currentUser.JoinedAt,
	)
	if err != nil {
		return nil, err
	}

	return &currentUser, nil
}

func (ur *UserRepositoryImpl) FindUserByNickname(nickname string) (*models.User, error) {
	ctx, cancel := newDBContext()
	defer cancel()

	stmt := `SELECT id, name, nickname, image_path, password, is_admin, joined_at
	FROM users
	WHERE nickname = $1`

	var user models.User
	err := ur.DB.SQL.QueryRowContext(ctx, stmt, nickname).Scan(
		&user.ID,
		&user.Name,
		&user.Nickname,
		&user.ImagePath,
		&user.Password,
		&user.IsAdmin,
		&user.JoinedAt,
	)
	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (ur *UserRepositoryImpl) CreateUser(userPostData *models.RegisterUser, imagePath string) (*models.CurrentUser, error) {
	ctx, cancel := newDBContext()
	defer cancel()

	stmt := `INSERT INTO users (name, nickname, password, image_path)
		VALUES ($1, $2, $3, $4) 
		RETURNING id, name, image_path, nickname, is_admin, joined_at`

	var currentUser models.CurrentUser
	err := ur.DB.SQL.QueryRowContext(
		ctx,
		stmt,
		userPostData.Name,
		userPostData.Nickname,
		userPostData.Password,
		imagePath,
	).Scan(
		&currentUser.ID,
		&currentUser.Name,
		&currentUser.ImagePath,
		&currentUser.Nickname,
		&currentUser.IsAdmin,
		&currentUser.JoinedAt,
	)
	if err != nil {
		return &models.CurrentUser{}, err
	}

	return &currentUser, nil
}

func (ur *UserRepositoryImpl) UpdateUser(user *models.CurrentUser) (*models.CurrentUser, error) {
	ctx, cancel := newDBContext()
	defer cancel()

	stmt := `UPDATE users
	SET name = $1, image_path = $2
	WHERE id = $3
	RETURNING id, name, image_path, nickname, is_admin, joined_at`

	var currentUser models.CurrentUser
	err := ur.DB.SQL.QueryRowContext(
		ctx,
		stmt,
		user.Name, user.ImagePath, user.ID).
		Scan(
			&currentUser.ID,
			&currentUser.Name,
			&currentUser.ImagePath,
			&currentUser.Nickname,
			&currentUser.IsAdmin,
			&currentUser.JoinedAt)
	if err != nil {
		return nil, err
	}

	return &currentUser, nil
}
