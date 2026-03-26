package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/aprimr/chautari/db"
	"github.com/aprimr/chautari/utils"
	"github.com/jackc/pgx/v5"
)

// RequestExists - check if any request exists between users (any status)
func RequestExists(ctx context.Context, senderId, receiverId string) (bool, error) {
	var rid string
	query := "SELECT rid FROM requests WHERE (sender_id = $1 AND receiver_id = $2) OR (sender_id = $2 AND receiver_id = $1)"
	err := db.Pool.QueryRow(ctx, query, senderId, receiverId).Scan(&rid)
	if err == pgx.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// SendRequest - send new request
func SendRequest(ctx context.Context, senderId, receiverId string) error {
	query := "INSERT INTO requests (sender_id, receiver_id, status) VALUES($1, $2, 'pending')"
	_, err := db.Pool.Exec(ctx, query, senderId, receiverId)
	return err
}

// AcceptRequest - accept pending request
func AcceptRequest(ctx context.Context, requestId, currentUserId string) error {
	utils.LogDebug("Rid repo: " + requestId)

	query := "UPDATE requests SET status = 'accepted', updated_at = $1 WHERE rid = $2 AND receiver_id = $3 AND status = 'pending'"
	cmdTag, err := db.Pool.Exec(ctx, query, time.Now(), requestId, currentUserId)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("request not found or already processed")
	}
	return nil
}
