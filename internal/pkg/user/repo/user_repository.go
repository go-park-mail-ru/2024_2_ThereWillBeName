package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user models.User) (uint, error) {
	var userID uint
	query := "INSERT INTO users (login, email, password, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, user.Login, user.Email, user.Password).Scan(&userID)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return 0, models.ErrAlreadyExists
		}
		return 0, err
	}

	return userID, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	query := "SELECT id, login, email, password, created_at FROM users WHERE email = $1"
	row := r.db.QueryRowContext(ctx, query, email)
	err := row.Scan(&user.ID, &user.Login, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found with email: %s", email)
		}
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, user models.User) error {
	query := "UPDATE users SET login = $1, email=$2, password = $3 WHERE id = $4"
	_, err := r.db.ExecContext(ctx, query, user.Login, user.Email, user.Password, user.ID)
	return err
}

func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *UserRepositoryImpl) GetUsers(ctx context.Context, count, offset int64) ([]models.User, error) {
	query := "SELECT id, login, email, created_at FROM users LIMIT $1 OFFSET $2"
	rows, err := r.db.QueryContext(ctx, query, count, offset)
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

func (r *UserRepositoryImpl) GetAvatarPathByUserId(ctx context.Context, userID uint) (string, error) {
	queryBuilder := squirrel.Select("avatarPath").
		From("users").
		Where(squirrel.Eq{"id": userID}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		log.Println("1: %w", err)
		return "", fmt.Errorf("failed to build query: %w", models.ErrInternal)
	}

	var avatarPath string
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&avatarPath)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found: %w", models.ErrNotFound)
		}
		log.Println("2:", err)
		return "", fmt.Errorf("failed to retrieve avatar path: %w", models.ErrInternal)
	}

	return avatarPath, nil
}

func (r *UserRepositoryImpl) UpdateAvatarPathByUserId(ctx context.Context, userID uint, avatarPath string) error {
	queryBuilder := squirrel.Update("users").
		Set("avatarpath", avatarPath).
		Where(squirrel.Eq{"id": userID}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build update query: %w", models.ErrInternal)
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", models.ErrInternal)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %w", models.ErrInternal)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated: %w", models.ErrNotFound)
	}

	return nil
}

func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, userID uint) (models.UserProfile, error) {
	queryBuilder := squirrel.Select("login, email, avatarPath").
		From("users").
		Where(squirrel.Eq{"id": userID}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return models.UserProfile{}, fmt.Errorf("failed to build query: %w", models.ErrInternal)
	}

	row := r.db.QueryRowContext(ctx, query, args...)

	var userProfile models.UserProfile

	if err := row.Scan(&userProfile.Login, &userProfile.Email, &userProfile.AvatarPath); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.UserProfile{}, models.ErrNotFound
		}
		return models.UserProfile{}, fmt.Errorf("failed to scan user profile: %w", err)
	}

	return userProfile, nil
}