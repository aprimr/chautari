package repository

import (
	"context"
	"fmt"

	"github.com/aprimr/chautari/db"
	"github.com/jackc/pgx/v5"
)

// Check if Request is already exists - return (exists, error)
func RequestExists(ctx context.Context, userId, contactId string) (bool, error) {
	var id string
	query := "SELECT id FROM contacts WHERE ((user_id=$1 AND contact_id=$2) OR (user_id=$2 AND contact_id=$1))"
	err := db.Pool.QueryRow(ctx, query, userId, contactId).Scan(&id)
	if err == pgx.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// Create new contact request
func CreateContactRequest(ctx context.Context, userId, contactId string) error {
	query := "INSERT INTO contacts (user_id, contact_id) VALUES($1, $2)"
	cmdTag, err := db.Pool.Exec(ctx, query, userId, contactId)
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("failed to send request")
	}
	return err
}
