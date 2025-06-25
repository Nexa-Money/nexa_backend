package factory

import (
	"nexa/internal/model"
	"time"

	"github.com/google/uuid"
)

type CategoryFactory struct{}

func NewCategoryFactory() *CategoryFactory {
	return &CategoryFactory{}
}

func (cf *CategoryFactory) CreateCategory(category model.Category) *model.Category {
	category.ID = uuid.New()
	category.CreatedAt = time.Now().UTC().Truncate(time.Second)
	return &category
}
