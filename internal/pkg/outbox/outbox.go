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
			// Можно залогировать или обновить статус как `failed`, если нужно
			continue
		}

		// Обновляем статус записи в `outbox` на `processed`
		if err := UpdateOutboxRecordStatus(ctx, o.db, record.ID, "processed"); err != nil {
			log.Printf("Failed to update status for event ID %d: %v", record.ID, err)
		}
	}

	return nil
}

func (o *OutboxListener) handleEvent(ctx context.Context, record models.OutboxRecord) error {
	switch record.EventType {
	case "review_added", "review_updated", "review_deleted":
		return HandleReviewEvent(ctx, o.db, record.Payload)
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

func HandleReviewEvent(ctx context.Context, db *dblogger.DB, payload string) error {
	var data struct {
		PlaceID int `json:"place_id"`
	}
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		return fmt.Errorf("failed to parse payload: %w", err)
	}

	// Вызов функции пересчёта среднего рейтинга
	if err := RecalculateAverageRating(ctx, db, data.PlaceID); err != nil {
		return fmt.Errorf("failed to recalculate average rating for place ID %d: %w", data.PlaceID, err)
	}

	log.Printf("Successfully recalculated average rating for place ID: %d\n", data.PlaceID)
	return nil
}

func RecalculateAverageRating(ctx context.Context, db *dblogger.DB, placeID int) error {
	query := `
        UPDATE places
        SET average_rating = (
            SELECT COALESCE(AVG(rating), 0)
            FROM reviews
            WHERE place_id = $1
        )
        WHERE id = $1
    `

	_, err := db.ExecContext(ctx, query, placeID)
	if err != nil {
		return fmt.Errorf("failed to recalculate average rating: %w", err)
	}

	return nil
}
