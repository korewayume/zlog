package zlog

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/v2pro/plz/gls"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func getLumberJackWriter(fileName string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		LocalTime:  true,
		MaxBackups: 5,
		MaxAge:     1,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func goidCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d,%d", os.Getpid(), gls.GoID()))
	zapcore.ShortCallerEncoder(caller, enc)
}

func getEncoder(colored bool) zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeCaller = goidCallerEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if colored {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// MustLogger MustLogger
func MustLogger(lvl zap.AtomicLevel) *zap.Logger {
	logger, err := NewLogger(lvl)
	if err != nil {
		panic(err)
	}
	return logger
}

// NewLogger NewLogger
func NewLogger(lvl zap.AtomicLevel) (*zap.Logger, error) {
	name := filepath.Base(os.Args[0])
	core := zapcore.NewTee(
		zapcore.NewCore(
			getEncoder(false),
			getLumberJackWriter(fmt.Sprintf("%s.verbose.log", name)),
			lvl,
		),
		zapcore.NewCore(
			getEncoder(false),
			getLumberJackWriter(fmt.Sprintf("%s.error.log", name)),
			zapcore.ErrorLevel,
		),
		zapcore.NewCore(
			getEncoder(true),
			zapcore.Lock(os.Stdout),
			lvl,
		),
		zapcore.NewCore(
			getEncoder(true),
			zapcore.Lock(os.Stderr),
			zapcore.ErrorLevel,
		),
	)
	wrapCore := func(zapcore.Core) zapcore.Core {
		return core
	}

	config := zap.NewDevelopmentConfig()
	config.Development = false

	return config.Build(
		zap.WithCaller(true),
		zap.WrapCore(wrapCore),
		zap.AddCallerSkip(0),
	)
}
