package services

import (
	"context"
	"fmt"

	"github.com/aprimr/chautari/repository"
)

func SendContactRequest(ctx context.Context, userId, contactId string) error {
	// Check if request exists
	exists, err := repository.RequestExists(ctx, userId, contactId)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("request exists")
	}

	// Add contact
	err = repository.CreateContactRequest(ctx, userId, contactId)
	if err != nil {
		return err
	}

	return nil
}
