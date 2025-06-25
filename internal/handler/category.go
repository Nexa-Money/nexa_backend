package handler

import (
	"nexa/internal/factory"
	"nexa/internal/model"
	"nexa/internal/repository"
	"nexa/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	CategoryRepository *repository.CategoryRepository
	CategoryFactory    *factory.CategoryFactory
}

func NewCategoryHandler(conn *pgxpool.Pool) *CategoryHandler {
	return &CategoryHandler{
		CategoryRepository: repository.NewCategoryRepository(conn),
		CategoryFactory:    factory.NewCategoryFactory(),
	}
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var body model.Category
	if err := c.BodyParser(&body); err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Dados inválidos",
			Error:      err,
		})
	}

	if err := utils.ValidateCategory(c, h.CategoryRepository, body); err != nil {
		return err
	}

	category := h.CategoryFactory.CreateCategory(body)
	if err := h.CategoryRepository.InsertCategory(*category); err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 500,
			Message:    "Erro ao criar categoria",
			Error:      err,
		})
	}

	return utils.HttpSuccess(c, utils.SuccessStructure{
		StatusCode: 201,
		Message:    "Categoria criada com sucesso",
		Data:       category,
	})
}

func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 400,
			Message:    "ID de usuário inválido",
			Error:      err,
		})
	}
	categories, err := h.CategoryRepository.GetAllCategories(userID)
	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 500,
			Message:    "Erro ao buscar categorias",
			Error:      err,
		})
	}
	return c.JSON(categories)
}

func (h *CategoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 400,
			Message:    "ID inválido",
			Error:      err,
		})
	}
	category, err := h.CategoryRepository.GetCategoryByID(id)
	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 404,
			Message:    "Categoria não encontrada",
			Error:      err,
		})
	}
	return c.JSON(category)
}

func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 400,
			Message:    "ID inválido",
			Error:      err,
		})
	}
	var body model.Category
	if err := c.BodyParser(&body); err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Dados inválidos",
			Error:      err,
		})
	}
	if err := h.CategoryRepository.UpdateCategory(id, body); err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 500,
			Message:    "Erro ao atualizar categoria",
			Error:      err,
		})
	}
	return utils.HttpSuccess(c, utils.SuccessStructure{
		StatusCode: 200,
		Message:    "Categoria atualizada com sucesso",
		Data:       body,
	})
}

func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 400,
			Message:    "ID inválido",
			Error:      err,
		})
	}
	if err := h.CategoryRepository.DeleteCategory(id); err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 500,
			Message:    "Erro ao deletar categoria",
			Error:      err,
		})
	}
	return c.JSON(fiber.Map{"message": "Categoria excluída com sucesso"})
}
