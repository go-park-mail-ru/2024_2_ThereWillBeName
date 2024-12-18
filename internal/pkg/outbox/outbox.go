package outbox

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type OutboxListener struct {
	db *dblogger.DB
}

func NewOutboxListener(db *dblogger.DB) *OutboxListener {
	return &OutboxListener{db: db}
}

func (o *OutboxListener) StartListening(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second) // Интервал проверки outbox
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := o.processOutboxEvents(ctx); err != nil {
				log.Printf("Error processing outbox events: %v", err)
			}
		}
	}
}

func (o *OutboxListener) processOutboxEvents(ctx context.Context) error {
	records, err := GetPendingOutboxRecords(ctx, o.db)
	if err != nil {
		return fmt.Errorf("failed to get pending outbox records: %w", err)
	}

	for _, record := range records {
		if err := o.handleEvent(ctx, record); err != nil {
			log.Printf("Failed to handle event ID %d: %v", record.ID, err)
			continue
		}

		if err := UpdateOutboxRecordStatus(ctx, o.db, record.ID, "processed"); err != nil {
			log.Printf("Failed to update status for event ID %d: %v", record.ID, err)
		}
	}

	return nil
}

func (o *OutboxListener) handleEvent(ctx context.Context, record models.OutboxRecord) error {
	switch record.EventType {
	case "review_created", "review_updated", "review_deleted":
		return HandleReviewEvent(ctx, o.db, record.Payload, record.EventType)
	case "avatar_uploaded":
		return HandleUploadAvatarEvent(ctx, o.db, record.Payload)
	default:
		return fmt.Errorf("unknown event type: %s", record.EventType)
	}
}

func InsertOutboxRecord(ctx context.Context, db *dblogger.DB, eventType, payload string) error {
	query := `INSERT INTO outbox (event_type, payload, "status", created_at) 
				VALUES ($1, $2, 'pending', NOW())`

	_, err := db.ExecContext(ctx, query, eventType, payload)
	if err != nil {
		return fmt.Errorf("failed to insert outbox record: %w", err)
	}

	return nil
}

func GetPendingOutboxRecords(ctx context.Context, db *dblogger.DB) ([]models.OutboxRecord, error) {
	query := `SELECT id, event_type, payload, status, created_at, processed_at 
				FROM outbox WHERE status = 'pending'`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve outbox records: %w", err)
	}
	defer rows.Close()

	var records []models.OutboxRecord
	for rows.Next() {
		var record models.OutboxRecord
		if err := rows.Scan(&record.ID, &record.EventType, &record.Payload, &record.Status, &record.CreatedAt, &record.ProcessedAt); err != nil {
			return nil, fmt.Errorf("failed to scan outbox record: %w", err)
		}
		records = append(records, record)
	}

	return records, nil
}

func UpdateOutboxRecordStatus(ctx context.Context, db *dblogger.DB, id int, status string) error {
	query := `UPDATE outbox SET status = $1, processed_at = NOW() WHERE id = $2`

	_, err := db.ExecContext(ctx, query, status, id)
	if err != nil {
		return fmt.Errorf("failed to update outbox record status: %w", err)
	}

	return nil
}

func HandleReviewEvent(ctx context.Context, db *dblogger.DB, payload string, eventType string) error {
	var data struct {
		UserID  int `json:"user_id"`
		PlaceID int `json:"place_id"`
	}
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		return fmt.Errorf("failed to parse payload: %w", err)
	}

	if err := Recalculate(ctx, db, data.PlaceID); err != nil {
		log.Printf("failed to recalculate average rating or reviews for place ID %d", data.PlaceID)
	}

	switch eventType {
	case "review_created":
		if err := CheckAndInsertReviewAchievements(ctx, db, data.UserID); err != nil {
			return fmt.Errorf("failed to check achievements for user %d: %w", data.UserID, err)
		}
	case "review_deleted":
		if err := CheckAndRemoveReviewAchievements(ctx, db, data.UserID); err != nil {
			return fmt.Errorf("failed to remove achievements for user %d: %w", data.UserID, err)
		}
	}

	log.Printf("Successfully recalculated average rating and reviews' number for place ID: %d\n", data.PlaceID)
	return nil
}

func Recalculate(ctx context.Context, db *dblogger.DB, placeID int) error {
	query := `
        UPDATE place
        SET rating = (
            SELECT COALESCE(AVG(rating), 0)
            FROM review
            WHERE place_id = $1	
        ),
		 number_of_reviews = (
            SELECT COUNT(*)
            FROM review
            WHERE place_id = $1
        )
        WHERE id = $1
    `

	_, err := db.ExecContext(ctx, query, placeID)
	if err != nil {
		return fmt.Errorf("failed to recalculate average rating or reviews' number: %w", err)
	}

	return nil
}

func CheckAndInsertReviewAchievements(ctx context.Context, db *dblogger.DB, userID int) error {
	var reviewCount int
	err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM review WHERE user_id = $1", userID).Scan(&reviewCount)
	if err != nil {
		return fmt.Errorf("failed to count reviews for user %d: %w", userID, err)
	}

	// Достижения и их условия
	achievements := map[int]int{
		1:  1, //ID достижения за 1 отзыв
		5:  2, // ID достижения за 5 отзывов
		10: 3, // ID достижения за 10 отзывов
	}

	for count, achievementID := range achievements {
		if reviewCount >= count {
			// Проверяем, есть ли уже запись о достижении
			var exists bool
			err := db.QueryRowContext(ctx, `
				SELECT EXISTS (
					SELECT 1 FROM user_achievement
					WHERE user_id = $1 AND achievement_id = $2
				)`, userID, achievementID).Scan(&exists)
			if err != nil {
				return fmt.Errorf("failed to check achievement existence: %w", err)
			}

			if !exists {
				_, err := db.ExecContext(ctx, `
					INSERT INTO user_achievement (user_id, achievement_id) 
					VALUES ($1, $2)`, userID, achievementID)
				if err != nil {
					return fmt.Errorf("failed to insert achievement: %w", err)
				}
				log.Printf("Achievement %d unlocked for user %d\n", achievementID, userID)
			}
		}
	}

	return nil
}

func CheckAndRemoveReviewAchievements(ctx context.Context, db *dblogger.DB, userID int) error {
	var reviewCount int
	err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM review WHERE user_id = $1", userID).Scan(&reviewCount)
	if err != nil {
		return fmt.Errorf("failed to count reviews for user %d: %w", userID, err)
	}

	achievements := map[int]int{
		1:  1, //ID достижения за 1 отзыв
		5:  2, // ID достижения за 5 отзывов
		10: 3, // ID достижения за 10 отзывов
	}

	for count, achievementID := range achievements {
		if reviewCount < count {
			// Удаляем достижение, если условия больше не выполняются
			_, err := db.ExecContext(ctx, `
				DELETE FROM user_achievement
				WHERE user_id = $1 AND achievement_id = $2`, userID, achievementID)
			if err != nil {
				return fmt.Errorf("failed to remove achievement: %w", err)
			}
			log.Printf("Achievement %d removed for user %d\n", achievementID, userID)
		}
	}

	return nil
}

func HandleUploadAvatarEvent(ctx context.Context, db *dblogger.DB, payload string) error {
	var data struct {
		UserID int `json:"user_id"`
	}

	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		return fmt.Errorf("failed to parse avatar upload payload: %w", err)
	}

	avatarUploadAchievementID := 4

	var exists bool
	err := db.QueryRowContext(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM user_achievement
			WHERE user_id = $1 AND achievement_id = $2
		)`, data.UserID, avatarUploadAchievementID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check achievement existence: %w", err)
	}

	if !exists {
		_, err := db.ExecContext(ctx, `
			INSERT INTO user_achievement (user_id, achievement_id) 
			VALUES ($1, $2)`, data.UserID, avatarUploadAchievementID)
		if err != nil {
			return fmt.Errorf("failed to insert achievement for user %d: %w", data.UserID, err)
		}
		log.Printf("Achievement for uploading avatar unlocked for user %d\n", data.UserID)
	}

	return nil
}
