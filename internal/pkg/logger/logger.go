package logger

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"

	"github.com/fatih/color"
)

type ctxKey string

const (
	slogFields ctxKey = "slog_fields"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(slogFields).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}

	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	// Подготовка записи лога в формате JSON
	// logEntry := map[string]interface{}{
	// 	"time":    r.Time.Format("2006-01-02 15:04:05"),
	// 	"level":   r.Level.String(),
	// 	"message": r.Message,
	// 	"fields":  fields,
	// }
	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}
	timeStr := r.Time.Format("2006-01-02 15:04:05")

	msg := color.CyanString(r.Message)

	h.l.Println(timeStr, level, msg, string(b))

	return nil
}

func NewPrettyHandler(out io.Writer, opts PrettyHandlerOptions) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts), // Используем JSON обработчик
		l:       log.New(out, "", 0),
	}

	return h
}

func AppendCtx(parent context.Context, attr slog.Attr) context.Context {
	if parent == nil {
		parent = context.Background()
	}

	if v, ok := parent.Value(slogFields).([]slog.Attr); ok {
		v = append(v, attr)
		return context.WithValue(parent, slogFields, v)
	}

	v := []slog.Attr{}
	v = append(v, attr)
	return context.WithValue(parent, slogFields, v)
}

func LogRequestStart(ctx context.Context, method, uri string) context.Context {
	logCtx := context.Background()
	logCtx = AppendCtx(logCtx, slog.String("method", method))
	logCtx = AppendCtx(logCtx, slog.String("uri", uri))
	return logCtx
}
