package services

import (
	"context"
	"fmt"

	"github.com/aprimr/chautari/models"
	"github.com/aprimr/chautari/repository"
	"github.com/aprimr/chautari/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(ctx context.Context, registerInput models.RegisterInput) error {
	// Check if user already exists with that email
	exists, err := repository.UserExists(ctx, registerInput.Email)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("email already exists")
	}

	// Generate a username and check if it is already taken
	var username string
	for i := range 10 {
		username = utils.GenerateUsername()

		exists, err := repository.UsernameAlreadyTaken(ctx, username)
		if err != nil {
			return err
		}
		if !exists {
			break // unique username found
		}

		if i == 9 {
			return fmt.Errorf("could not generate unique username")
		}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Replace real password with hashed password
	registerInput.Password = string(hashedPassword)

	// Save user
	err = repository.CreateUser(ctx, registerInput, username)
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(ctx context.Context, loginInput models.LoginInput) (string, error) {
	// Check if user exists
	exists, err := repository.UserExists(ctx, loginInput.Email)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", fmt.Errorf("user not found")
	}

	// Get user from DB
	user, err := repository.GetUserByEmail(ctx, loginInput.Email)
	if err != nil {
		if err.Error() == "user not found" {
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}

	// Match password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
	if err != nil {
		return "", fmt.Errorf("incorrect password")
	}

	// Generate JWT token
	jwtToken, err := utils.CreateToken(user.Uid, user.Username, user.Email)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func UpdatePassword(ctx context.Context, uid string, passwordInput models.UpdatePasswordInput) error {
	// get password hash
	hash, err := repository.GetPasswordHash(ctx, uid)
	if err != nil {
		return err
	}

	// compare hash and password
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordInput.Password))
	if err != nil {
		return fmt.Errorf("incorrect password")
	}

	// hash new password
	newHash, err := bcrypt.GenerateFromPassword([]byte(passwordInput.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update password
	err = repository.UpdatePassword(ctx, uid, string(newHash))
	if err != nil {
		return err
	}

	return nil
}
