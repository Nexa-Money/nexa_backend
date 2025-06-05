package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID `json:"id,omitempty" db:"id"`
	UserID      uuid.UUID `json:"user_id,omitempty" db:"user_id"`
	Category    string    `json:"category,omitempty" db:"category"` // TODO: change to UUID
	Amount      float32   `json:"amount,omitempty" db:"amount"`
	Date        string    `json:"date,omitempty" db:"date"`
	Description string    `json:"description,omitempty" db:"description"`
	Type        string    `json:"type,omitempty" db:"type"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
}
