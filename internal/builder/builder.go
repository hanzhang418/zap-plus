package builder

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(level zapcore.Level) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(level)
	return cfg.Build()
}
