package repository

import (
	"context"
	"nexa/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository struct {
	Conn *pgxpool.Pool
}

func NewCategoryRepository(conn *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{Conn: conn}
}

func (r *CategoryRepository) InsertCategory(category model.Category) error {
	_, err := r.Conn.Exec(context.Background(),
		`INSERT INTO categories (id, user_id, name, type, budget, created_at) VALUES ($1, $2, $3, $4, $5, $6)`,
		category.ID, category.UserID, category.Name, category.Type, category.Budget, category.CreatedAt)
	return err
}

func (r *CategoryRepository) GetAllCategories(userID uuid.UUID) ([]model.Category, error) {
	rows, err := r.Conn.Query(context.Background(),
		`SELECT id, user_id, name, type, budget, created_at FROM categories WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		err := rows.Scan(&c.ID, &c.UserID, &c.Name, &c.Type, &c.Budget, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *CategoryRepository) GetCategoryByID(id uuid.UUID) (model.Category, error) {
	var c model.Category
	err := r.Conn.QueryRow(context.Background(),
		`SELECT id, user_id, name, type, budget, created_at FROM categories WHERE id = $1`, id).
		Scan(&c.ID, &c.UserID, &c.Name, &c.Type, &c.Budget, &c.CreatedAt)
	return c, err
}

func (r *CategoryRepository) UpdateCategory(id uuid.UUID, category model.Category) error {
	_, err := r.Conn.Exec(context.Background(),
		`UPDATE categories SET name=$1, type=$2, budget=$3 WHERE id=$4`,
		category.Name, category.Type, category.Budget, id)
	return err
}

func (r *CategoryRepository) DeleteCategory(id uuid.UUID) error {
	_, err := r.Conn.Exec(context.Background(), `DELETE FROM categories WHERE id=$1`, id)
	return err
}
