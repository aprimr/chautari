package services

import (
	"context"
	"fmt"

	"github.com/aprimr/chautari/repository"
)

func SendContactRequest(ctx context.Context, senderId, receiverId string) error {
	// Check if request exists
	exists, err := repository.RequestExists(ctx, senderId, receiverId)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("request exists")
	}

	// Add contact
	err = repository.CreateContactRequest(ctx, senderId, receiverId)
	if err != nil {
		return err
	}

	return nil
}
