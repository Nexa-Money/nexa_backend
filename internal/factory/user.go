package factory

import (
	"nexa/internal/model"
	"time"

	"github.com/google/uuid"
)

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (uf *UserFactory) CreateUser(user model.User) *model.User {
	user.ID = uuid.New()
	user.CreatedAt = time.Now().UTC().Truncate(time.Second)
	user.UpdatedAt = time.Now().UTC().Truncate(time.Second)

	return &user
}
