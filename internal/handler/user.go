package handler

import (
	"fmt"
	"nexa/internal/factory"
	"nexa/internal/model"
	"nexa/internal/repository"
	"nexa/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserRepository repository.UserRepository
	UserFactory    factory.UserFactory
}

func NewUserHandler(conn *pgxpool.Pool) *UserHandler {
	return &UserHandler{
		UserRepository: *repository.NewUserRepository(conn),
		UserFactory:    *factory.NewUserFactory(),
	}
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	var body model.User

	err := c.BodyParser(&body)
	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Dados inválidos",
			Error:      err,
		})
	}

	err = utils.ValidateUser(c, body)
	if err != nil {
		return err
	}

	user := uh.UserFactory.CreateUser(body)

	err = uh.UserRepository.InsertUser(user)
	if err != nil {
		fmt.Println("Erro ao inserir usuario: ", err)
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 500,
			Message:    "Erro interno no servidor",
			Error:      err,
		})
	}

	return utils.HttpSuccess(c, utils.SuccessStructure{
		StatusCode: 201,
		Message:    "Usuário criado com sucesso",
		Data:       user,
	})
}

func (uh *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := uh.UserRepository.GetAllUsers()
	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 500,
			Message:    "Erro interno no servidor",
			Error:      err,
		})
	}

	return c.JSON(users)
}

func (uh *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := uh.UserRepository.GetUserByID(id)
	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 500,
			Message:    "Erro interno no servidor",
			Error:      err,
		})
	}

	return c.JSON(user)
}

func (uh *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Error",
			Error:      err,
		})
	}

	if err := uh.UserRepository.UpdateUser(id, user); err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 500,
			Message:    "Error",
			Error:      err,
		})
	}

	return utils.HttpSuccess(c, utils.SuccessStructure{
		StatusCode: 200,
		Message:    "Usuário atualizado com sucesso",
		Data:       user,
	})
}

func (uh *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := uh.UserRepository.DeleteUser(id); err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 500,
			Message:    "Error",
			Error:      err,
		})
	}

	return c.JSON(fiber.Map{"message": "Usuário excluído com sucesso"})
}

func (uh *UserHandler) LoginUser(c *fiber.Ctx) error {
	var body model.User

	err := c.BodyParser(&body)
	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Dados inválidos",
			Error:      err,
		})
	}
	email := body.Email

	user, err := uh.UserRepository.GetUserByEmail(email)
	if err != nil {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 404,
			Message:    "Email e/ou senha incorretos",
			Error:      err,
		})
	}

	if body.Password != user.Password {
		return utils.HttpError(c, utils.ErrorStructure{
			StatusCode: 401,
			Message:    "Email e/ou senha incorretos",
			Error:      err,
		})
	} else {
		return utils.HttpSuccess(c, utils.SuccessStructure{
			StatusCode: 200,
			Message:    "Login bem sucedido",
			Data:       user,
		})
	}

}
