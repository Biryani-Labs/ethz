package logs

import (
	"fmt"
	"os"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Initilize() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	config.EncoderConfig.EncodeName = zapcore.FullNameEncoder

	var err error
	Logger, err = config.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing logger: %v\n", err)
		os.Exit(1)
	}
}

func Error(err error, message string, fields ...zap.Field) error {
	_, file, line, _ := runtime.Caller(1)
	Logger.Error(message, append([]zap.Field{
		zap.String("file", file),
		zap.Int("line", line),
	}, fields...)...)
	return err
}

func Warn(message string, fields ...zap.Field) {
	_, file, line, _ := runtime.Caller(1)
	Logger.Warn(message, append([]zap.Field{
		zap.String("file", file),
		zap.Int("line", line),
	}, fields...)...)
}

func Info(message string, fields ...zap.Field) {
	_, file, line, _ := runtime.Caller(1)
	Logger.Info(message, append([]zap.Field{
		zap.String("file", file),
		zap.Int("line", line),
	}, fields...)...)
}

func Debug(message string, fields ...zap.Field) {
	_, file, line, _ := runtime.Caller(1)
	Logger.Debug(message, append([]zap.Field{
		zap.String("file", file),
		zap.Int("line", line),
	}, fields...)...)
}
