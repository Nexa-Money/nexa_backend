package handler

import (
	"nexa/internal/factory"
	"nexa/internal/model"
	"nexa/internal/repository"
	"nexa/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	TransactionRepository *repository.TransactionRepository
	TransactionFactory    *factory.TransactionFactory
}

func NewTransactionHandler(conn *pgxpool.Pool) *TransactionHandler {
	return &TransactionHandler{
		TransactionRepository: repository.NewTransactionRepository(conn),
		TransactionFactory:    factory.NewTransactionFactory(),
	}
}

func (th *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var body model.Transaction

	err := c.BodyParser(&body)
	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Dados inválidos",
			Error:      err,
		})
	}

	transaction := th.TransactionFactory.CreateTransaction(body)

	err = th.TransactionRepository.InsertTransaction(transaction)
	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 500,
			Message:    "Erro interno no servidor",
			Error:      err,
		})
	}

	return utils.HttpSuccess(c, utils.SuccessStructure{
		StatusCode: 201,
		Message:    "Transação criado com sucesso",
		Data:       transaction,
	})
}

func (th *TransactionHandler) GetTransactions(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	transactions, err := th.TransactionRepository.GetTransactions(userID)

	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 500,
			Message:    "Erro interno no servidor",
			Error:      err,
		})
	}

	return c.JSON(transactions)
}
