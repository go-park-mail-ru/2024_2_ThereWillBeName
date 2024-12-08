package dblogger

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	pool   *pgxpool.Pool
	logger *slog.Logger
}

func NewDB(pool *pgxpool.Pool, logger *slog.Logger) *DB {
	return &DB{
		pool:   pool,
		logger: logger,
	}
}

func (d *DB) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	start := time.Now()
	rows, err := d.pool.Query(ctx, query, args...)
	duration := time.Since(start)

	d.logger.DebugContext(ctx, "Executing Query",
		slog.String("query", query),
		slog.Any("args", args),
		slog.Duration("duration", duration),
		slog.String("error", errToString(err)),
	)

	return rows, err
}

func (d *DB) Exec(ctx context.Context, query string, args ...interface{}) (pgx.CommandTag, error) {
	start := time.Now()
	tag, err := d.pool.Exec(ctx, query, args...)
	duration := time.Since(start)

	d.logger.DebugContext(ctx, "Executing Exec",
		slog.String("query", query),
		slog.Any("args", args),
		slog.Duration("duration", duration),
		slog.String("error", errToString(err)),
	)

	return tag, err
}

func (d *DB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	start := time.Now()
	row := d.pool.QueryRow(ctx, query, args...)
	duration := time.Since(start)

	d.logger.DebugContext(ctx, "Executing QueryRow",
		slog.String("query", query),
		slog.Any("args", args),
		slog.Duration("duration", duration),
	)

	return row
}

func (d *DB) Prepare(ctx context.Context, query string) (*pgxpool.Conn, pgx.PreparedStatement, error) {
	start := time.Now()
	conn, err := d.pool.Acquire(ctx)
	if err != nil {
		d.logger.DebugContext(ctx, "Error acquiring connection",
			slog.String("query", query),
			slog.Duration("duration", time.Since(start)),
			slog.String("error", errToString(err)),
		)
		return nil, pgx.PreparedStatement{}, err
	}
	stmt, err := conn.Conn().Prepare(ctx, query, query)
	duration := time.Since(start)

	d.logger.DebugContext(ctx, "Preparing statement",
		slog.String("query", query),
		slog.Duration("duration", duration),
		slog.String("error", errToString(err)),
	)

	return conn, stmt, err
}

func (d *DB) Close() {
	d.pool.Close()
}

func errToString(err error) string {
	if err != nil {
		return err.Error()
	}
	return "nil"
}
