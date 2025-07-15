package zap_plus

import (
	"github.com/hanzhang418/zap-plus/internal/builder"
	"github.com/hanzhang418/zap-plus/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerImpl struct {
	logger *zap.Logger
}

func (l loggerImpl) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l loggerImpl) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l loggerImpl) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func New(level zapcore.Level) (logger.Logger, error) {
	z, err := builder.New(level)
	if err != nil {
		return nil, err
	}
	return z, nil
}
