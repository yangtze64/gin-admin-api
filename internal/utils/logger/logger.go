package logger

import (
	"context"
	"gin-admin-api/pkg/logx"
)

func WithContext(ctx context.Context) logx.Logger {
	return logx.WithContext(ctx)
}
