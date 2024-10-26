package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
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
	query, args, err := squirrel.Insert(`"user"`).
		Columns("login", "email", "password", "created_at").
		Values(user.Login, user.Email, user.Password, squirrel.Expr("NOW()")).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("couldn't build query: %w", err)
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	// query := "INSERT INTO user (login, email, password, created_at) VALUES ($1, $2, $3, NOW())"
	// _, err := r.db.ExecContext(ctx, query, user.Login, user.Email, user.Password)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return models.ErrAlreadyExists
		}
	}

	return err
}

func (r *RepositoryImpl) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	query, args, err := squirrel.Select("id", "login", "email", "password", "created_at").
		From(`"user"`).
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("couldn't build query: %w", err)
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	// query := "SELECT id, login, email, password, created_at FROM user WHERE email = $1"
	// row := r.db.QueryRowContext(ctx, query, email)
	err = row.Scan(&user.ID, &user.Login, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found with email: %s", email)
		}
		return models.User{}, err
	}
	return user, nil
}

func (r *RepositoryImpl) UpdateUser(ctx context.Context, user models.User) error {
	query, args, err := squirrel.Update(`"user"`).
		Set("login", user.Login).
		Set("email", user.Email).
		Set("password", user.Password).
		Where("id = ?", user.ID).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("couldn't build query: %w", err)
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	// query := "UPDATE user SET login = $1, email=$2, password = $3 WHERE id = $4"
	// _, err := r.db.ExecContext(ctx, query, user.Login, user.Email, user.Password, user.ID)
	return err
}

func (r *RepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	query, args, err := squirrel.Delete(`"user"`).
		Where("id = ?", id).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("couldn't build query: %w", err)
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	// query := "DELETE FROM user WHERE id = $1"
	// _, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *RepositoryImpl) GetUsers(ctx context.Context, count, offset int64) ([]models.User, error) {
	query, args, err := squirrel.Select("id", "login", "email", "created_at").
		From(`"user"`).
		Limit(uint64(count)).
		Offset(uint64(offset)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("couldn't build query: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	// query := "SELECT id, login, email, created_at FROM user LIMIT $1 OFFSET $2"
	// rows, err := r.db.QueryContext(ctx, query, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Login, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
