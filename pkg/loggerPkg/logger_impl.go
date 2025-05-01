package loggerPkg

import (
	"Chat-Websocket/config"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

type loggerImpl struct {
	Cfg    *config.Config
	Logger *zap.Logger
}

func NewLoggerImpl(cfg *config.Config) ILogger {
	return &loggerImpl{Cfg: cfg}
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

var logEmojiMap = map[zapcore.Level]string{
	zapcore.DebugLevel:  "🔥",
	zapcore.InfoLevel:   "ℹ️",
	zapcore.WarnLevel:   "⚠️",
	zapcore.ErrorLevel:  "❌",
	zapcore.DPanicLevel: "😱",
	zapcore.PanicLevel:  "💥",
	zapcore.FatalLevel:  "💀",
}

func (l *loggerImpl) getLoggerLevel(cfg *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}
	return level
}

func (l *loggerImpl) InitLogger() {
	logLevel := l.getLoggerLevel(l.Cfg)

	logFilePath := filepath.Join("logs", "application.log")
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		panic("Failed to create logs directory: " + err.Error())
	}
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic("Failed to create log file: " + err.Error())
	}

	fileWriter := zapcore.AddSync(logFile)

	jsonEncoderCfg := zap.NewProductionEncoderConfig()
	jsonEncoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	jsonEncoder := zapcore.NewJSONEncoder(jsonEncoderCfg)

	fileCore := zapcore.NewCore(jsonEncoder, fileWriter, logLevel)

	consoleEncoderCfg := zap.NewDevelopmentEncoderConfig()
	consoleEncoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderCfg)

	consoleWriter := zapcore.AddSync(os.Stdout)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleWriter, logLevel)

	core := zapcore.NewTee(fileCore, consoleCore)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	l.Logger = logger
}

func (l *loggerImpl) logWithEmoji(level zapcore.Level, args ...interface{}) {
	emoji := logEmojiMap[level]
	message := fmt.Sprintf("%s %v", emoji, args)

	switch level {
	case zapcore.DebugLevel:
		l.Logger.Debug(message)
	case zapcore.InfoLevel:
		l.Logger.Info(message)
	case zapcore.WarnLevel:
		l.Logger.Warn(message)
	case zapcore.ErrorLevel:
		l.Logger.Error(message)
	case zapcore.DPanicLevel:
		l.Logger.DPanic(message)
	case zapcore.PanicLevel:
		l.Logger.Panic(message)
	case zapcore.FatalLevel:
		l.Logger.Fatal(message)
	default:
		l.Logger.Info(message)
	}
}

func (l *loggerImpl) Debug(args ...interface{}) {
	l.logWithEmoji(zapcore.DebugLevel, args...)
}

func (l *loggerImpl) Info(args ...interface{}) {
	l.logWithEmoji(zapcore.InfoLevel, args...)
}

func (l *loggerImpl) Warn(args ...interface{}) {
	l.logWithEmoji(zapcore.WarnLevel, args...)
}

func (l *loggerImpl) Error(args ...interface{}) {
	l.logWithEmoji(zapcore.ErrorLevel, args...)
}

func (l *loggerImpl) DPanic(args ...interface{}) {
	l.logWithEmoji(zapcore.DPanicLevel, args...)
}

func (l *loggerImpl) Panic(args ...interface{}) {
	l.logWithEmoji(zapcore.PanicLevel, args...)
}

func (l *loggerImpl) Fatal(args ...interface{}) {
	l.logWithEmoji(zapcore.FatalLevel, args...)
}
