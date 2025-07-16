package zapplus

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	// 导入我们自己的包
	"github.com/hanzhang418/zap-plus/internal/builder"
	"github.com/hanzhang418/zap-plus/pkg/logger"
)

// loggerImpl 是 logger.Logger 接口的具体实现，是未导出的私有类型。
type loggerImpl struct {
	zap *zap.Logger
}

func (l loggerImpl) Debug(msg string, fields ...zap.Field) {
	l.zap.Debug(msg, fields...)
}

func (l loggerImpl) Info(msg string, fields ...zap.Field) {
	l.zap.Info(msg, fields...)
}

func (l loggerImpl) Error(msg string, fields ...zap.Field) {
	l.zap.Error(msg, fields...)
}

func (l loggerImpl) Panic(msg string, fields ...zap.Field) {
	l.zap.Panic(msg, fields...)
}

// New 是库的唯一入口，用于创建一个新的 Logger 实例。
func New(level zapcore.Level) (logger.Logger, error) {
	zl, err := builder.NewZapLogger(level)
	if err != nil {
		return nil, err
	}

	// 返回的是接口，而不是具体类型，实现了信息隐藏
	return &loggerImpl{zap: zl}, nil
}
