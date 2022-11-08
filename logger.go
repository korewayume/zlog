package zlog

import (
	"fmt"
	"os"

	"github.com/v2pro/plz/gls"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger logger
var Logger *zap.Logger

// GoidCallerEncoder caller中增加Goroutine ID
func GoidCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d,%d", os.Getpid(), gls.GoID()))
	zapcore.ShortCallerEncoder(caller, enc)
}

func init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "LEVEL",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 这里可以指定颜色
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   GoidCallerEncoder, // 全路径编码器
	}

	// 设置日志级别
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel), // 日志级别
		Development:      true,                                // 开发模式，堆栈跟踪
		Encoding:         "console",                           // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                       // 编码器配置
		OutputPaths:      []string{"stdout"},                  // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}
	// 构建日志
	var err error
	Logger, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(fmt.Sprintf("Logger 初始化失败: %v", err))
	}
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	Logger.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	Logger.DPanic(msg, fields...)
}

func Debugf(format string, a ...any) {
	Logger.Debug(fmt.Sprintf(format, a...))
}

func Infof(format string, a ...any) {
	Logger.Info(fmt.Sprintf(format, a...))
}

func Warnf(format string, a ...any) {
	Logger.Warn(fmt.Sprintf(format, a...))
}

func Errorf(format string, a ...any) {
	Logger.Error(fmt.Sprintf(format, a...))
}

func Fatalf(format string, a ...any) {
	Logger.Fatal(fmt.Sprintf(format, a...))
}

func Panicf(format string, a ...any) {
	Logger.Panic(fmt.Sprintf(format, a...))
}

func DPanicf(format string, a ...any) {
	Logger.DPanic(fmt.Sprintf(format, a...))
}
