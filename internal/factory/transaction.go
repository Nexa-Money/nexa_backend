package factory

import (
	"nexa/internal/model"
	"time"

	"github.com/google/uuid"
)

type TransactionFactory struct{}

func NewTransactionFactory() *TransactionFactory {
	return &TransactionFactory{}
}

func (tf *TransactionFactory) CreateTransaction(transaction model.Transaction) *model.Transaction {
	return &model.Transaction{
		ID:          uuid.New(),
		UserID:      transaction.UserID,
		CategoryID:  transaction.CategoryID,
		Amount:      transaction.Amount,
		Date:        transaction.Date,
		Description: transaction.Description,
		Type:        transaction.Type,
		CreatedAt:   time.Now().UTC().Truncate(time.Second),
	}
}
