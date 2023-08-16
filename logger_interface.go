package zlog

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZLogger struct {
	zap.Logger
	lvl zap.AtomicLevel
}

func (l *ZLogger) SetLevel(lvl zapcore.Level) {
	l.lvl.SetLevel(lvl)
}

func (l *ZLogger) Debugf(format string, a ...any) {
	l.Debug(fmt.Sprintf(format, a...))
}

func (l *ZLogger) Infof(format string, a ...any) {
	l.Info(fmt.Sprintf(format, a...))
}

func (l *ZLogger) Warnf(format string, a ...any) {
	l.Warn(fmt.Sprintf(format, a...))
}

func (l *ZLogger) Errorf(format string, a ...any) {
	l.Error(fmt.Sprintf(format, a...))
}

func (l *ZLogger) Fatalf(format string, a ...any) {
	l.Fatal(fmt.Sprintf(format, a...))
}

func (l *ZLogger) Panicf(format string, a ...any) {
	l.Panic(fmt.Sprintf(format, a...))
}

func (l *ZLogger) DPanicf(format string, a ...any) {
	l.DPanic(fmt.Sprintf(format, a...))
}

// Logger logger接口实例
var Logger *ZLogger

// MLogger 通过包名直接调用
var MLogger *ZLogger

func init() {
	lvl := zap.NewAtomicLevel()
	mLvl := zap.NewAtomicLevel()
	Logger = &ZLogger{Logger: *MustLogger(lvl), lvl: lvl}
	MLogger = &ZLogger{Logger: *MustLogger(mLvl).WithOptions(zap.AddCallerSkip(1)), lvl: mLvl}
}

func SetLevel(l zapcore.Level) {
	MLogger.SetLevel(l)
}

func Debug(msg string, fields ...zap.Field) {
	MLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	MLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	MLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	MLogger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	MLogger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	MLogger.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	MLogger.DPanic(msg, fields...)
}

func Debugf(format string, a ...any) {
	MLogger.Debugf(format, a...)
}

func Infof(format string, a ...any) {
	MLogger.Infof(format, a...)
}

func Warnf(format string, a ...any) {
	MLogger.Warnf(format, a...)
}

func Errorf(format string, a ...any) {
	MLogger.Errorf(format, a...)
}

func Fatalf(format string, a ...any) {
	MLogger.Fatalf(format, a...)
}

func Panicf(format string, a ...any) {
	MLogger.Panicf(format, a...)
}

func DPanicf(format string, a ...any) {
	MLogger.DPanicf(format, a...)
}
