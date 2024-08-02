package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapOptions struct {
	Encoder zapcore.Encoder
	Writer  zapcore.WriteSyncer
	Level   LogLevel
}

// ZapLogger uber.zap implements the Logger.
type ZapLogger struct {
	logger *zap.SugaredLogger
	level  *zap.AtomicLevel
}

func NewZapLogger(cfg ZapOptions, opts ...zap.Option) *ZapLogger {
	l, err := zapcore.ParseLevel(string(cfg.Level))
	if err != nil {
		Fatalf("failed to parse level '%s': %v", cfg.Level, err)
		l = zapcore.InfoLevel
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(l)
	core := zap.New(zapcore.NewCore(cfg.Encoder, cfg.Writer, &atomicLevel), opts...)
	return &ZapLogger{
		logger: core.Sugar(),
		level:  &atomicLevel,
	}
}

// Logger return the zap.Logger
func (z *ZapLogger) Logger() *zap.Logger {
	return z.logger.Desugar()
}

// WithFields adds new fields to log.
func (z *ZapLogger) WithFields(fields map[string]any) Logger {
	f := make([]any, 0, len(fields))
	for key, value := range fields {
		f = append(f, key, value)
	}
	return &ZapLogger{logger: z.logger.With(f...)}
}

// Trace logs a message at level Trace.
func (z *ZapLogger) Trace(args ...any) {
	z.logger.Debug(args...)
}

// Tracef logs a message at level Trace.
func (z *ZapLogger) Tracef(format string, args ...any) {
	z.logger.Debugf(format, args...)
}

// Debug logs a message at level Debug.
func (z *ZapLogger) Debug(args ...any) {
	z.logger.Debug(args...)
}

// Debugf logs a message at level Debug.
func (z *ZapLogger) Debugf(format string, args ...any) {
	z.logger.Debugf(format, args...)
}

// Info logs a message at level Info.
func (z *ZapLogger) Info(args ...any) {
	z.logger.Info(args...)
}

// Infof logs a message at level Info.
func (z *ZapLogger) Infof(format string, args ...any) {
	z.logger.Infof(format, args...)
}

// Warn logs a message at level Warn.
func (z *ZapLogger) Warn(args ...any) {
	z.logger.Warn(args...)
}

// Warnf logs a message at level Warn.
func (z *ZapLogger) Warnf(format string, args ...any) {
	z.logger.Warnf(format, args...)
}

// Error logs a message at level Error.
func (z *ZapLogger) Error(args ...any) {
	z.logger.Error(args...)
}

// Errorf logs a message at level Error.
func (z *ZapLogger) Errorf(format string, args ...any) {
	z.logger.Errorf(format, args...)
}

// Fatal logs a message at level Fatal.
func (z *ZapLogger) Fatal(args ...any) {
	z.logger.Fatal(args...)
}

// Fatalf logs a message at level Fatal.
func (z *ZapLogger) Fatalf(format string, args ...any) {
	z.logger.Fatalf(format, args...)
}

func (z *ZapLogger) GetLevel() LogLevel {
	return LogLevel(z.logger.Level().String())
}

func (z *ZapLogger) SetLevel(level LogLevel) {
	if l, err := zapcore.ParseLevel(string(level)); err == nil {
		z.level.SetLevel(l)
	}
}

func (z *ZapLogger) IsLevelEnabled(level LogLevel) bool {
	l, _ := zapcore.ParseLevel(string(level))
	return z.logger.Level().Enabled(l)
}

// NewJSONEncoder create the json encoder
func NewJSONEncoder() zapcore.Encoder {
	enc := zap.NewProductionEncoderConfig()
	enc.EncodeTime = TimeEncoder
	enc.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(enc)
}

// NewConsoleEncoder create the console encoder
func NewConsoleEncoder() zapcore.Encoder {
	enc := zap.NewProductionEncoderConfig()
	enc.EncodeTime = TimeEncoder
	enc.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(enc)
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	encodeTimeLayout(t, "2006-01-02 15:04:05", enc)
}

func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}

// NewFileWriter create a file writer
func NewFileWriter(file string, maxSize, maxBackups, maxAge int, compress bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    maxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: maxBackups, // 保留旧文件的最大个数
		MaxAge:     maxAge,     // 保留旧文件的最大天数
		Compress:   compress,   // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

// NewConsoleWriter create a console writer
func NewConsoleWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}
