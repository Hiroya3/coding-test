package log

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Infof(ctx context.Context, msg string, fields ...any)
	Warnf(ctx context.Context, msg string, fields ...any)
	Errorf(ctx context.Context, msg string, fields ...any)
}

type logger struct {
	Logrus *logrus.Logger
}

func NewLogger() Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	l.SetOutput(os.Stdout)
	return &logger{
		Logrus: l,
	}
}

func (l *logger) Infof(ctx context.Context, msg string, fields ...any) {
	l.Logrus.WithFields(logFields(ctx)).Infof(msg, fields...)
}

func (l *logger) Warnf(ctx context.Context, msg string, fields ...any) {
	l.Logrus.WithFields(logFields(ctx)).Warnf(msg, fields...)
}

func (l *logger) Errorf(ctx context.Context, msg string, fields ...any) {
	l.Logrus.WithFields(logFields(ctx)).Errorf(msg, fields...)
}

const (
	ContextKeyUserID = iota
)

var contextMapKey = map[int]string{
	ContextKeyUserID: "user_id",
}

func logFields(ctx context.Context) logrus.Fields {
	fields := map[string]any{}

	for k, v := range contextMapKey {
		ctxValue := ctx.Value(k)
		if ctxValue != nil {
			fields[v] = ctxValue
		}
	}
	return fields
}

func ContextWithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, ContextKeyUserID, userID)
}
