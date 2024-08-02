package logger

import (
	"go.uber.org/zap"
)

var defaultLogger Logger

func Init() {
	zapLogger := NewZapLogger(
		ZapOptions{
			Encoder: NewConsoleEncoder(),
			Writer:  NewConsoleWriter(),
			Level:   DebugLevel,
		},
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	)
	SetDefault(zapLogger)
}

func InitForConfig(cfg *LogConfig) {
	encoder := NewConsoleEncoder()
	if cfg.Encoder == "json" {
		encoder = NewJSONEncoder()
	}

	writer := NewConsoleWriter()
	if cfg.Output == "file" {
		writer = NewFileWriter(cfg.File, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge, cfg.Compress)
	}

	// logLevel, err := zapcore.ParseLevel(cfg.Level)
	// if err != nil {
	// 	logLevel = zapcore.InfoLevel
	// }

	defaultLogger = NewZapLogger(
		ZapOptions{
			Encoder: encoder,
			Writer:  writer,
			Level:   LogLevel(cfg.Level),
		},
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	)
}

func Default() Logger {
	return defaultLogger
}

func SetDefault(logger Logger) {
	defaultLogger = logger
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	defaultLogger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	defaultLogger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	defaultLogger.Fatalf(template, args...)
}
