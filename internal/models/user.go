package models

import "time"

type UserBaseProperty struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	ImagePath string    `json:"imagePath"`
	Nickname  string    `json:"nickname"`
	IsAdmin   bool      `json:"isAdmin"`
	JoinedAt  time.Time `json:"joinedAt"`
}

type User struct {
	UserBaseProperty
	Password string `json:"password"`
}

type UserPayload struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	IsAdmin  bool   `json:"isAdmin"`
}

type CurrentUser struct {
	UserBaseProperty
}

type RegisterUser struct {
	Name     string `json:"name" validate:"required,min=4,max=32"`
	Nickname string `json:"nickname" validate:"required,min=1,max=10"`
	Password string `json:"password" validate:"required,min=8"`
}
