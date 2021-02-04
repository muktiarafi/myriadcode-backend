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
	Password string `json:"-"`
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
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
