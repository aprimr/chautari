package repository

import (
	"context"
	"fmt"

	"github.com/aprimr/chautari/db"
	"github.com/aprimr/chautari/models"
	"github.com/jackc/pgx/v5"
)

func CreateUser(ctx context.Context, registerInput models.RegisterInput, username string) error {
	query := "INSERT INTO users (name, username, email, password, bio, profile_url) VALUES($1, $2, $3, $4, $5, $6)"
	_, err := db.Pool.Exec(ctx, query, registerInput.Name, username, registerInput.Email, registerInput.Password, registerInput.Bio, registerInput.ProfileUrl)
	return err
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := "SELECT uid, name, username, email, password, bio, profile_url, is_online, last_seen, is_active, is_deleted, created_at, updated_at FROM users WHERE email=$1"
	row := db.Pool.QueryRow(ctx, query, email)

	user := models.User{}
	err := row.Scan(
		&user.Uid,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Bio,
		&user.ProfileUrl,
		&user.IsOnline,
		&user.LastSeen,
		&user.IsActive,
		&user.IsDeleted,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
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
