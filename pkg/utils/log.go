/*
Copyright © 2022 Ike Murami <murami.ike@gmail.com>
*/
package utils

import (
	"context"

	"go.uber.org/zap"
)

type ctxLogger struct{}

// Add and get logger to/from context

func ContextWithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, logger)
}

func LoggerFromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(ctxLogger{}).(*zap.Logger); ok {
		return logger
	}
	return zap.L()
}
