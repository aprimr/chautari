package repository

import (
	"context"

	"github.com/aprimr/chautari/db"
	"github.com/aprimr/chautari/models"
)

func CreateUser(ctx context.Context, registerInput models.RegisterInput, username string) error {
	query := "INSERT INTO users (name, username, email, password, bio, profile_url) VALUES($1, $2, $3, $4, $5, $6)"
	_, err := db.Pool.Exec(ctx, query, registerInput.Name, username, registerInput.Email, registerInput.Password, registerInput.Bio, registerInput.ProfileUrl)
	return err
}

func UserExists(ctx context.Context, email string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE email=$1"
	err := db.Pool.QueryRow(ctx, query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func UsernameAlreadyTaken(ctx context.Context, username string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE username=$1"
	err := db.Pool.QueryRow(ctx, query, username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
