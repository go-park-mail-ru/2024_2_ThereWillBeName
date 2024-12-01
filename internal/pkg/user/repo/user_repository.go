package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type UserRepositoryImpl struct {
	db *dblogger.DB
}

func NewAuthRepository(db *dblogger.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user models.User) (uint, error) {
	var userID uint
	query := `INSERT INTO "user" (login, email, password_hash, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id`
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
	query := `SELECT id, login, email, password_hash FROM "user" WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email)
	err := row.Scan(&user.ID, &user.Login, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found with email: %s, %s", email, models.ErrNotFound)
		}

		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, user models.User) error {
	query := `UPDATE "user" SET login = $1, email=$2, password_hash = $3, updated_at = NOW() WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, user.Login, user.Email, user.Password, user.ID)
	return err
}

func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM "user" WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *UserRepositoryImpl) GetAvatarPathByUserId(ctx context.Context, userID uint) (string, error) {
	query := `SELECT avatar_path FROM "user" WHERE id=$1`

	row := r.db.QueryRowContext(ctx, query, userID)

	var avatarPath string
	err := row.Scan(&avatarPath)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found: %w", models.ErrNotFound)
		}
		return "", fmt.Errorf("failed to retrieve avatar path: %w", models.ErrInternal)
	}

	return avatarPath, nil
}

func (r *UserRepositoryImpl) UpdateAvatarPathByUserId(ctx context.Context, userID uint, avatarPath string) error {
	query := `UPDATE "user" SET avatar_path = $1 WHERE id = $2`

	result, err := r.db.ExecContext(ctx, query, avatarPath, userID)
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
	queryBuilder := squirrel.Select("login, email, avatar_path").
		From(`"user"`).
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
		return models.UserProfile{}, fmt.Errorf("failed to scan user profile: %w", models.ErrInternal)
	}

	return userProfile, nil
}

func (r *UserRepositoryImpl) UpdatePassword(ctx context.Context, userId uint, newPassword string) error {
	query := "UPDATE user SET password = $1 WHERE id = $2"

	result, err := r.db.ExecContext(ctx, query, newPassword, userId)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", models.ErrInternal)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %w", models.ErrInternal)
	}
	if rowsAffected == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (r UserRepositoryImpl) UpdateProfile(ctx context.Context, userID uint, login, email string) error {
	query := "UPDATE user SET email = $1, login = $2 WHERE id = $3"

	result, err := r.db.ExecContext(ctx, query, email, login, userID)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", models.ErrInternal)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %w", models.ErrInternal)
	}
	if rowsAffected == 0 {
		return models.ErrNotFound
	}
	return nil
}
