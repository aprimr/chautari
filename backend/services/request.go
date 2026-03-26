package services

import (
	"context"
	"fmt"

	"github.com/aprimr/chautari/repository"
	"github.com/aprimr/chautari/utils"
)

func SendRequest(ctx context.Context, senderId, receiverId string) error {
	// Check if request exists
	exists, err := repository.RequestExists(ctx, senderId, receiverId)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("request exists")
	}

	// Add contact
	err = repository.SendRequest(ctx, senderId, receiverId)
	if err != nil {
		return err
	}

	return nil
}

func AcceptRequest(ctx context.Context, requestId, currentUserId string) error {
	utils.LogDebug("Rid service: " + requestId)

	err := repository.AcceptRequest(ctx, requestId, currentUserId)
	return err
}
