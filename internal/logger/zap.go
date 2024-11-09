package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// TimeEncoderRFC3339UTC logs 2019-08-13T04:39:11Z
	TimeEncoderRFC3339UTC = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z0700"))
	})
)

func NewLogger() *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = TimeEncoderRFC3339UTC
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		zapcore.AddSync(os.Stderr),
		Level,
	)

	return zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}
