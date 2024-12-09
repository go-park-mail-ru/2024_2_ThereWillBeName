package outbox

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	"fmt"
)

func InsertOutboxRecord(ctx context.Context, db *sql.DB, eventType, payload string) error {
	query := `INSERT INTO outbox (event_type, payload, status, created_at) 
              VALUES ($1, $2, 'pending', NOW())`

	_, err := db.ExecContext(ctx, query, eventType, payload)
	if err != nil {
		return fmt.Errorf("failed to insert outbox record: %w", err)
	}

	return nil
}

func GetPendingOutboxRecords(ctx context.Context, db *sql.DB) ([]models.OutboxRecord, error) {
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

func UpdateOutboxRecordStatus(ctx context.Context, db *sql.DB, id int, status string) error {
	query := `UPDATE outbox SET status = $1, processed_at = NOW() WHERE id = $2`

	_, err := db.ExecContext(ctx, query, status, id)
	if err != nil {
		return fmt.Errorf("failed to update outbox record status: %w", err)
	}

	return nil
}
