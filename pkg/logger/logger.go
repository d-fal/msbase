package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logPath string
)

func SetLogPath(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	logPath = path
}

func GetZapLogger(logfile string) *zap.Logger {
	logfile = strings.ReplaceAll(logfile, "/", "_")
	logger, exists := zapLoggersMap[logfile]

	if !exists {

		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s.log", logPath, logfile),
			MaxSize:    30, // megabytes
			MaxBackups: 3,
			MaxAge:     1, // days
		})
		coreConfig := zap.NewProductionEncoderConfig()
		coreConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(coreConfig),
			w,
			zap.InfoLevel,
		)

		zapLoggersMap[logfile] = zap.New(core)

		return zapLoggersMap[logfile]
	}
	return logger
}
