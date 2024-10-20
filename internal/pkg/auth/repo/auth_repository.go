package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type RepositoryImpl struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) CreateUser(ctx context.Context, user models.User) error {
	query := "INSERT INTO users (login, password, created_at) VALUES ($1, $2, NOW())"
	_, err := r.db.ExecContext(ctx, query, user.Login, user.Password)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return models.ErrUserAlreadyExists
		}
	}
	return err
}

func (r *RepositoryImpl) GetUserByLogin(ctx context.Context, login string) (models.User, error) {
	var user models.User
	query := "SELECT id, login, password, created_at FROM users WHERE login = $1"
	row := r.db.QueryRowContext(ctx, query, login)
	err := row.Scan(&user.ID, &user.Login, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found with login: %s", login)
		}
		return models.User{}, err
	}
	return user, nil
}

func (r *RepositoryImpl) UpdateUser(ctx context.Context, user models.User) error {
	query := "UPDATE users SET login = $1, password = $2 WHERE id = $3"
	_, err := r.db.ExecContext(ctx, query, user.Login, user.Password, user.ID)
	return err
}

func (r *RepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *RepositoryImpl) GetUsers(ctx context.Context, count, offset int64) ([]models.User, error) {
	query := "SELECT id, login, created_at FROM users LIMIT $1 OFFSET $2"
	rows, err := r.db.QueryContext(ctx, query, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Login, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
