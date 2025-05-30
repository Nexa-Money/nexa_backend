package utils

import (
	"fmt"
	"nexa/internal/model"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ValidateUser(c *fiber.Ctx, u model.User) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

	if strings.TrimSpace(u.Name) == "" {
		fmt.Println("NOME INVÁLIDO DETECTADO")

		return HttpError(c, ErrorStructure{
			StatusCode: 400,
			Message:    "Campo nome obrigatório",
			Error:      "O campo 'name' não pode estar vazio.",
		})
	}

	if len(strings.TrimSpace(u.Name)) < 3 {
		return HttpError(c, ErrorStructure{
			StatusCode: 400,
			Message:    "Nome inválido",
			Error:      "O nome deve ter pelo menos 3 caracteres.",
		})
	}

	if strings.TrimSpace(u.Email) == "" {
		return HttpError(c, ErrorStructure{
			StatusCode: 400,
			Message:    "Campo email obrigatório",
			Error:      "O campo 'email' não pode estar vazio.",
		})
	}

	if !emailRegex.MatchString(u.Email) {
		return HttpError(c, ErrorStructure{
			StatusCode: 400,
			Message:    "Email inválido",
			Error:      "Formato de e-mail inválido.",
		})
	}

	if strings.TrimSpace(u.Password) == "" {
		return HttpError(c, ErrorStructure{
			StatusCode: 400,
			Message:    "Campo senha obrigatório",
			Error:      "O campo 'password' não pode estar vazio.",
		})
	}

	if len(u.Password) < 6 {
		return HttpError(c, ErrorStructure{
			StatusCode: 400,
			Message:    "Senha fraca",
			Error:      "A senha deve conter pelo menos 6 caracteres.",
		})
	}

	return nil
}
