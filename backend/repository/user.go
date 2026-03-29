package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/aprimr/chautari/db"
	"github.com/aprimr/chautari/models"
	"github.com/jackc/pgx/v5"
)

func CreateUser(ctx context.Context, registerInput models.RegisterInput, username string) error {
	query := "INSERT INTO users (name, username, email, password, bio, profile_url) VALUES($1, $2, $3, $4, $5, $6)"
	_, err := db.Pool.Exec(ctx, query, registerInput.Name, username, registerInput.Email, registerInput.Password, registerInput.Bio, registerInput.ProfileUrl)
	return err
}

func GetPasswordHash(ctx context.Context, uid string) (string, error) {
	query := "SELECT password FROM users WHERE uid=$1 AND is_deleted=FALSE"

	// fire query and scan row
	var hash string
	row := db.Pool.QueryRow(ctx, query, uid)
	err := row.Scan(&hash)

	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}
	return hash, nil
}

func UpdatePassword(ctx context.Context, uid, hash string) error {
	query := "UPDATE users SET password=$1, updated_at=$2 WHERE uid=$3 AND is_active=TRUE"

	// Fire query
	cmdTag, err := db.Pool.Exec(ctx, query, hash, time.Now(), uid)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("error updating password")
	}

	return nil
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

func GetUserByUid(ctx context.Context, uid string) (*models.User, error) {
	query := "SELECT uid, name, username, email, password, bio, profile_url, is_online, last_seen, is_active, is_deleted, created_at, updated_at FROM users WHERE uid=$1"
	row := db.Pool.QueryRow(ctx, query, uid)

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

func IsUserActive(ctx context.Context, uid string) (bool, error) {
	isActive := false
	query := "SELECT is_active FROM users WHERE uid=$1"
	err := db.Pool.QueryRow(ctx, query, uid).Scan(&isActive)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, fmt.Errorf("invalid uid")
		}
		return isActive, err
	}
	return isActive, nil
}

func IsUserDeleted(ctx context.Context, uid string) (bool, error) {
	isDeleted := false
	query := "SELECT is_deleted FROM users WHERE uid=$1"
	err := db.Pool.QueryRow(ctx, query, uid).Scan(&isDeleted)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, fmt.Errorf("invalid uid")
		}
		return isDeleted, err
	}
	return isDeleted, nil
}

func UpdateLastSeen(ctx context.Context, uid string) error {
	query := "UPDATE users SET last_seen=$1 WHERE uid=$2"
	_, err := db.Pool.Exec(ctx, query, time.Now(), uid)
	return err
}

func SearchUser(ctx context.Context, searchText string, uid string) ([]models.UserInfo, error) {
	query := "SELECT uid, name, username, email, bio, profile_url, is_online, last_seen FROM users WHERE (name ILIKE $1 || '%' OR username ILIKE $1 || '%') AND uid != $2 AND is_active=TRUE AND is_deleted=FALSE LIMIT 20"

	// Fire query and scan rows
	users := []models.UserInfo{}
	rows, err := db.Pool.Query(ctx, query, searchText, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.UserInfo{}
		err := rows.Scan(
			&user.Uid,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.Bio,
			&user.ProfileUrl,
			&user.IsOnline,
			&user.LastSeen,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func UpdateUser(ctx context.Context, uid string, updateData models.UpdateProfileInput) (models.UserInfo, error) {
	query := "UPDATE users SET name=$1, username=$2, bio=$3, updated_at=$4 WHERE uid=$5 AND is_active=TRUE RETURNING uid, name, username, email, bio, profile_url, is_online, last_seen"

	// Fire query and scan
	updatedUser := models.UserInfo{}
	row := db.Pool.QueryRow(ctx, query, updateData.Name, updateData.Username, updateData.Bio, time.Now(), uid)
	err := row.Scan(
		&updatedUser.Uid,
		&updatedUser.Name,
		&updatedUser.Username,
		&updatedUser.Email,
		&updatedUser.Bio,
		&updatedUser.ProfileUrl,
		&updatedUser.IsOnline,
		&updatedUser.LastSeen,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return models.UserInfo{}, fmt.Errorf("user not found or inactive")
		}
		return models.UserInfo{}, err
	}

	return updatedUser, nil
}
