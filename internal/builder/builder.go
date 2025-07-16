package builder

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapLogger 是一个内部函数，用于构建 zap.Logger 实例。
func NewZapLogger(level zapcore.Level) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(level) // 设置日志级别

	// 构建 logger，并应用初始字段
	return cfg.Build(zap.AddCaller(), zap.AddCallerSkip(1))
}
