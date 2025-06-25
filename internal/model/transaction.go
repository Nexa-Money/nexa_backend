package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID `json:"id,omitempty" db:"id"`
	UserID      uuid.UUID `json:"user_id,omitempty" db:"user_id"`
	CategoryID  uuid.UUID `json:"category_id,omitempty" db:"category_id"` // TODO: change to UUID
	Amount      float64   `json:"amount,omitempty" db:"amount"`
	Date        time.Time `json:"date,omitempty" db:"date"`
	Description string    `json:"description,omitempty" db:"description"`
	Type        string    `json:"type,omitempty" db:"type"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
}
