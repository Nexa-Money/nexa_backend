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

func (tr *TransactionRepository) InsertTransaction(transaction *model.Transaction) error {
	query := `INSERT INTO "transaction" (id, user_id, category, amount, date, description, type, created_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := tr.Conn.Exec(
		context.Background(),
		query,
		&transaction.ID, &transaction.UserID, &transaction.CategoryID, &transaction.Amount, &transaction.Date, &transaction.Description, &transaction.Type, &transaction.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %w", err)
	}

	return nil
}

func (tr *TransactionRepository) GetTransactions(id string) ([]model.Transaction, error) {
	rows, err := tr.Conn.Query(context.Background(), `SELECT id, user_id, category, amount, date, description, type, created_at FROM "transaction" WHERE user_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var transaction model.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.CategoryID, &transaction.Amount, &transaction.Date, &transaction.Description, &transaction.Type, &transaction.CreatedAt); err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
