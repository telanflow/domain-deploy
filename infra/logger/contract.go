package logger

// LogLevel is Logger Level type
type LogLevel string

const (
	// TraceLevel has more verbose message than debug level
	TraceLevel LogLevel = "trace"
	// DebugLevel has verbose message
	DebugLevel LogLevel = "debug"
	// InfoLevel is default log level
	InfoLevel LogLevel = "info"
	// WarnLevel is for logging messages about possible issues
	WarnLevel LogLevel = "warn"
	// ErrorLevel is for logging errors
	ErrorLevel LogLevel = "error"
	// FatalLevel is for logging fatal messages. The system shuts down after logging the message.
	FatalLevel LogLevel = "fatal"
)

type Logger interface {
	WithFields(map[string]any) Logger
	Trace(args ...any)
	Tracef(format string, args ...any)
	Debug(args ...any)
	Debugf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	GetLevel() LogLevel
	SetLevel(LogLevel)
	IsLevelEnabled(level LogLevel) bool
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `json:"level,omitempty"`      // 日志等级
	Output     string `json:"output,omitempty"`     // file console
	Encoder    string `json:"encoder,omitempty"`    // json console
	File       string `json:"file,omitempty"`       // 文件名
	MaxSize    int    `json:"maxSize,omitempty"`    // 日志文件分隔大小 单位：MB
	MaxBackups int    `json:"maxBackups,omitempty"` // 保留旧日志文件的最大个数
	MaxAge     int    `json:"maxAge,omitempty"`     // 保留旧日志文件的最大天数
	Compress   bool   `json:"compress,omitempty"`   // 是否压缩/归档旧日志文件
}
