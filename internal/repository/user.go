package repository

import (
	"context"
	"fmt"
	"nexa/internal/model"
	"time"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	Conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

func (ur *UserRepository) InsertUser(user *model.User) error {
	query := `INSERT INTO "user" (id, name, email, password, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := ur.Conn.Exec(
		context.Background(),
		query,
		user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

func (ur *UserRepository) GetAllUsers() ([]model.User, error) {
	rows, err := ur.Conn.Query(context.Background(), `SELECT id, name, email, created_at, updated_at FROM "user"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User

	query := `SELECT id, name, email, created_at, updated_at FROM "user" WHERE id = $1`
	err := ur.Conn.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) UpdateUser(id string, user model.User) error {
	currentUser, err := ur.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Name == "" {
		user.Name = currentUser.Name
	}

	if user.Email == "" {
		user.Email = currentUser.Email
	}

	if user.Password == "" {
		user.Password = currentUser.Password
	}

	user.UpdatedAt = time.Now().UTC().Truncate(time.Second)

	query := `UPDATE "user" SET name = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5`
	_, err = ur.Conn.Exec(context.Background(), query, user.Name, user.Email, user.Password, user.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUser(id string) error {
	query := `DELETE FROM "user" WHERE id = $1`

	_, err := ur.Conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}
