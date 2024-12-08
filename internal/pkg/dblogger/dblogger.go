package dblogger

import (
	"2024_2_ThereWillBeName/internal/pkg/config"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgconn" // Импорт pgconn для CommandTag
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
func SetupDBPool(cfg *config.Config, logger *slog.Logger) (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Database.DbUser,
		cfg.Database.DbPass,
		cfg.Database.DbHost,
		cfg.Database.DbPort,
		cfg.Database.DbName,
	)

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %v", err)
	}

	// Настройка пула соединений
	config.MaxConns = int32(cfg.Database.MaxConnections) // Максимальное количество соединений
	config.MinConns = 2                                  // Минимальное количество соединений
	config.HealthCheckPeriod = 1 * time.Minute           // Период проверки соединений
	config.MaxConnIdleTime = 5 * time.Minute             // Максимальное время простоя соединения
	config.ConnConfig.PreferSimpleProtocol = true        // Упрощенный протокол для производительности

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %v", err)
	}

	return pool, nil
}

func (d *DB) Acquire(ctx context.Context) (*pgxpool.Conn, error) {
	return d.pool.Acquire(ctx)
}

func (d *DB) Prepare(ctx context.Context, queryName string, query string) (*pgxpool.Conn, error) {
	start := time.Now()

	conn, err := d.pool.Acquire(ctx)
	if err != nil {
		d.logger.ErrorContext(ctx, "Failed to acquire connection",
			slog.String("queryName", queryName),
			slog.String("error", errToString(err)),
		)
		return nil, err
	}

	_, err = conn.Conn().Prepare(ctx, queryName, query)
	duration := time.Since(start)

	d.logger.DebugContext(ctx, "Preparing statement",
		slog.String("queryName", queryName),
		slog.String("query", query),
		slog.Duration("duration", duration),
		slog.String("error", errToString(err)),
	)

	if err != nil {
		// Освобождаем соединение, если произошла ошибка
		conn.Release()
		return nil, err
	}

	return conn, nil
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

func (d *DB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
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

// Закрытие пула соединений
func (d *DB) Close() {
	d.pool.Close()
}

func errToString(err error) string {
	if err != nil {
		return err.Error()
	}
	return "nil"
}
