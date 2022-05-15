package entity

import "time"

type ClientDiscount struct {
	ClientID     int64     `json:"client_id"`
	ClientName   string    `json:"client_name"`
	ClientNumber string    `json:"client_number"`
	Sale         int8      `json:"sale"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
