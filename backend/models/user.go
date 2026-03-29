package models

import (
	"time"
)

type User struct {
	Uid        string    `json:"uid"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Bio        string    `json:"bio"`
	ProfileUrl string    `json:"profile_url"`
	IsOnline   bool      `json:"is_online"`
	LastSeen   time.Time `json:"last_seen"`
	IsActive   bool      `json:"is_active"`
	IsDeleted  bool      `json:"is_deleted"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserInfo struct {
	Uid        string    `json:"uid"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Bio        string    `json:"bio"`
	ProfileUrl string    `json:"profile_url"`
	IsOnline   bool      `json:"is_online"`
	LastSeen   time.Time `json:"last_seen"`
}

type UpdateProfileInput struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
}

type RegisterInput struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Bio        string `json:"bio"`
	ProfileUrl string `json:"profile_url"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
