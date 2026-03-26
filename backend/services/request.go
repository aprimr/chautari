package services

import (
	"context"
	"fmt"

	"github.com/aprimr/chautari/repository"
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
	err := repository.AcceptRequest(ctx, requestId, currentUserId)
	return err
}

func CancelRequest(ctx context.Context, requestId, senderId string) error {
	err := repository.CancelRequest(ctx, requestId, senderId)
	return err
}


func RejectRequest(ctx context.Context, requestId, receiverId string) error {
	err := repository.RejectRequest(ctx, requestId, receiverId)
	return err
}