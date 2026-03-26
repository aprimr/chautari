package services

import (
	"context"
	"fmt"

	"github.com/aprimr/chautari/models"
	"github.com/aprimr/chautari/repository"
)

func GetMe(ctx context.Context, uid string) (*models.User, error) {
	// check if user is active
	isActive, err := repository.IsUserActive(ctx, uid)
	if err != nil {
		return nil, err
	}
	if !isActive {
		return nil, fmt.Errorf("user is inactive")
	}

	// check if user is deleted
	isDeleted, err := repository.IsUserDeleted(ctx, uid)
	if err != nil {
		return nil, err
	}
	if isDeleted {
		return nil, fmt.Errorf("user is deleted")
	}

	// update last seen
	err = repository.UpdateLastSeen(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("error updating last_seen")
	}

	// get me
	user, err := repository.GetUserByUid(ctx, uid)
	if err != nil {
		return nil, err
	}

	return user, err
}

func SearchUser(ctx context.Context, searchText string, uid string) ([]models.UserInfo, error) {
	users, err := repository.SearchUser(ctx, searchText, uid)
	if err != nil {
		return nil, err
	}

	return users, nil
}
