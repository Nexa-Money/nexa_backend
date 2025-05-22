package utils

import (
	"nexa/internal/model"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ValidateUser(c *fiber.Ctx, u model.User) error {

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if strings.TrimSpace(u.Name) == "" || strings.TrimSpace(u.Email) == "" || strings.TrimSpace(u.Password) == "" {
		return HttpError(c, ErrorStructure{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Campos obrigatórios faltando",
			Error:      "nome, email e senha são campos obrigatórios!",
		})
	}

	if len(u.Password) < 6 {
		return HttpError(c, ErrorStructure{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Campo senha incompleto",
			Error:      "Sua senha deve ter pelo menos 6 dígitos!",
		})
	}

	if !emailRegex.MatchString(u.Email) {
		return HttpError(c, ErrorStructure{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Campo email incompleto",
			Error:      "A estrutura do seu e-mail provavelmente está errada!",
		})
	}

	return nil
}
