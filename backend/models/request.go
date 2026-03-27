package models

import "time"

type Request struct {
	Rid        string    `json:"rid"`
	SenderId   string    `json:"sender_id"`
	ReceiverId string    `json:"receiver_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Friend struct {
	Rid       string    `json:"rid"`
	FriendId  string    `json:"friend_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
