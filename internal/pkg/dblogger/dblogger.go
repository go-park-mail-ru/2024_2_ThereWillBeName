package dblogger

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"reflect"
	"time"
)

type DB struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewDB(db *sql.DB, logger *slog.Logger) *DB {
	return &DB{
		db:     db,
		logger: logger,
	}
}

func (d *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	rows, err := d.db.QueryContext(ctx, query, args...)
	duration := time.Since(start)

	d.logger.DebugContext(ctx, "Executing QueryContext",
		slog.String("query", query),
		slog.Any("args_info", argsInfo(args)),
		slog.Duration("duration", duration),
		slog.String("error", errToString(err)),
	)

	return rows, err
}

func (d *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	result, err := d.db.ExecContext(ctx, query, args...)
	duration := time.Since(start)

	d.logger.DebugContext(ctx, "Executing ExecContext",
		slog.String("query", query),
		slog.Any("args", argsInfo(args)),
		slog.Duration("duration", duration),
		slog.String("error", errToString(err)),
	)

	return result, err
}

func (d *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	start := time.Now()
	row := d.db.QueryRowContext(ctx, query, args...)
	duration := time.Since(start)

	d.logger.DebugContext(ctx, "Executing QueryRowContext",
		slog.String("query", query),
		slog.Any("args", argsInfo(args)),
		slog.Duration("duration", duration),
	)

	return row
}

func (d *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	start := time.Now()
	stmt, err := d.db.PrepareContext(ctx, query)
	duration := time.Since(start)

	d.logger.DebugContext(ctx, "Preparing statement",
		slog.String("query", query),
		slog.Duration("duration", duration),
		slog.String("error", errToString(err)),
	)

	return stmt, err
}

func (d *DB) Close() error {
	return d.db.Close()
}

func errToString(err error) string {
	if err != nil {
		return err.Error()
	}
	return "nil"
}

func argsInfo(args []interface{}) string {
	if len(args) == 0 {
		return "no args"
	}

	types := make([]string, len(args))
	for i, arg := range args {
		types[i] = reflect.TypeOf(arg).String()
	}

	return fmt.Sprintf("%d args: [%s]", len(args), fmt.Sprintf("%s", types))
}
