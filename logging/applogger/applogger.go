package applogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var sugaredLogger *zap.SugaredLogger

func init()  {
	encoderCfg := zapcore.EncoderConfig {
		TimeKey:		"time",
		MessageKey:     "msg",
		LevelKey:       "level",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
	}

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zap.DebugLevel)
	sugaredLogger = zap.New(core).Sugar()
}


func UpdateLog(level zapcore.LevelEnabler)  {
	encoderCfg := zapcore.EncoderConfig {
		TimeKey:		"time",
		MessageKey:     "msg",
		LevelKey:       "level",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
	}

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, level)
	sugaredLogger = zap.New(core).Sugar()
}

func Fatal(template string, args ...interface{}) {
	sugaredLogger.Fatalf(template, args...)
}

func Error(template string, args ...interface{}) {
	sugaredLogger.Errorf(template, args...)
}

func Panic(template string, args ...interface{}) {
	sugaredLogger.Panicf(template, args...)
}

func Warn(template string, args ...interface{}) {
	sugaredLogger.Warnf(template, args...)
}

func Info(template string, args ...interface{}) {
	sugaredLogger.Infof(template, args...)
}

func Debug(template string, args ...interface{}) {
	sugaredLogger.Debugf(template, args...)
}