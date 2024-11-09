package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Level  = zapcore.InfoLevel
	logger = NewLogger()
	sugar  = logger.Sugar()
)

func Default() *zap.Logger {
	return logger
}

func Debug(args ...any) {
	sugar.Debug(args...)
}

func Info(args ...any) {
	sugar.Info(args...)
}

func Warn(args ...any) {
	sugar.Warn(args...)
}

func Error(args ...any) {
	sugar.Error(args...)
}

func DPanic(args ...any) {
	sugar.DPanic(args...)
}

func Panic(args ...any) {
	sugar.Panic(args...)
}

func Fatal(args ...any) {
	sugar.Fatal(args...)
}

func Debugf(template string, args ...any) {
	sugar.Debugf(template, args...)
}

func Infof(template string, args ...any) {
	sugar.Infof(template, args...)
}

func Warnf(template string, args ...any) {
	sugar.Warnf(template, args...)
}

func Errorf(template string, args ...any) {
	sugar.Errorf(template, args...)
}

func DPanicf(template string, args ...any) {
	sugar.DPanicf(template, args...)
}

func Panicf(template string, args ...any) {
	sugar.Panicf(template, args...)
}

func Fatalf(template string, args ...any) {
	sugar.Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...any) {
	sugar.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...any) {
	sugar.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...any) {
	sugar.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...any) {
	sugar.Errorw(msg, keysAndValues...)
}

func DPanicw(msg string, keysAndValues ...any) {
	sugar.DPanicw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...any) {
	sugar.Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...any) {
	sugar.Fatalw(msg, keysAndValues...)
}

func Debugln(args ...any) {
	sugar.Debugln(args...)
}

func Infoln(args ...any) {
	sugar.Infoln(args...)
}

func Warnln(args ...any) {
	sugar.Warnln(args...)
}

func Errorln(args ...any) {
	sugar.Errorln(args...)
}

func DPanicln(args ...any) {
	sugar.DPanicln(args...)
}

func Panicln(args ...any) {
	sugar.Panicln(args...)
}

func Fatalln(args ...any) {
	sugar.Fatalln(args...)
}
