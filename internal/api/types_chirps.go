package api

import (
	"time"

	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID	`json:"id"`
	UserID    uuid.UUID	`json:"user_id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	Body      string		`json:"body"`
}
