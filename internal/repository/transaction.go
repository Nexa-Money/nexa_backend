package repository

import (
	"context"
	"fmt"
	"nexa/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	Conn *pgxpool.Pool
}

func NewTransactionRepository(conn *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{
		Conn: conn,
	}
}

func (ur *TransactionRepository) InsertTransaction(transaction *model.Transaction) error {
	query := `INSERT INTO "transaction" (id, user_id, category, amount, date, description, type, created_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := ur.Conn.Exec(
		context.Background(),
		query,
		transaction.ID, transaction.UserID, transaction.Category, transaction.Amount, transaction.Date, transaction.Description, transaction.Type, transaction.CreatedAt,
	)
	if err != nil {
		fmt.Printf("Deu pau aqui: %v", err)
		return fmt.Errorf("failed to insert transaction: %w", err)
	}

	return nil
}
