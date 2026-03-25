package models

import "time"

type Contact struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	ContactId string    `json:"contact_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
