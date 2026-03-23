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
