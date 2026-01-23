package slf

import (
	"context"
	"fmt"
	"strings"
)

// A Level is a logging priority. Higher levels are more important.
type Level int

const (
	// LevelUnknown unknown logging level
	LevelUnknown Level = -1
	// LevelTrace logs are typically voluminous, and are usually disabled in
	// production.
	LevelTrace Level = iota
	// LevelDebug logs are typically voluminous, and are usually disabled in
	// production.
	LevelDebug
	// LevelInfo is the default logging priority.
	LevelInfo
	// LevelWarn logs are more important than Info, but don't need individual human review.
	LevelWarn
	// LevelError logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	LevelError
	// LevelFatal logs a message and then exit.
	LevelFatal

	levelUnknownString = "unknown"
	levelTraceString   = "trace"
	levelDebugString   = "debug"
	levelInfoString    = "info"
	levelWarnString    = "warn"
	levelWarningString = "warning"
	levelErrorString   = "error"
	levelFatalString   = "fatal"
)

// String returns literal level.
func (l Level) String() string {
	switch l {
	case LevelTrace:
		return levelTraceString
	case LevelDebug:
		return levelDebugString
	case LevelInfo:
		return levelInfoString
	case LevelWarn:
		return levelWarnString
	case LevelError:
		return levelErrorString
	case LevelFatal:
		return levelFatalString
	default:
		return levelUnknownString
	}
}

// ToLevel converts string to slf Level
func ToLevel(s string) (Level, error) {
	switch strings.ToLower(s) {
	case levelTraceString:
		return LevelTrace, nil
	case levelDebugString:
		return LevelDebug, nil
	case levelInfoString:
		return LevelInfo, nil
	case levelWarnString, levelWarningString:
		return LevelWarn, nil
	case levelErrorString:
		return LevelError, nil
	case levelFatalString:
		return LevelFatal, nil
	default:
		return LevelUnknown, fmt.Errorf("unknown level string: %s", s)
	}
}

// ContextHook extracts fields from context
type ContextHook func(context.Context) []Field

// ContextHooks is the collection of ContextHook
type ContextHooks []ContextHook

type Logger interface {
	GetName() string
	SetLevel(l Level)
	SetLevelString(l string) error
	GetLevel() Level
	Close() error

	IsEnableDebug() bool
	IsEnableInfo() bool
	IsEnableWarn() bool
	IsEnableError() bool
	IsEnableFatal() bool

	Debugln(v ...interface{})
	Infoln(v ...interface{})
	Warnln(v ...interface{})
	Errorln(v ...interface{})
	Fatalln(v ...interface{})

	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})

	AddContextHook(h ContextHook)
	ClearContextHooks()
	AddCallerSkip(skip int) Logger

	WithError(err error) Logger
	WithField(key string, val interface{}) Logger
	WithFields(fields ...Field) Logger
	WithContext(ctx context.Context) Logger

	Infow(msg string, fields ...Field)
	Warnw(msg string, fields ...Field)
	Errorw(msg string, fields ...Field)
	Fatalw(msg string, fields ...Field)
	Debugw(msg string, fields ...Field)

	Debugc(ctx context.Context, msg string, fields ...Field)
	Infoc(ctx context.Context, msg string, fields ...Field)
	Warnc(ctx context.Context, msg string, fields ...Field)
	Errorc(ctx context.Context, msg string, fields ...Field)
	Fatalc(ctx context.Context, msg string, fields ...Field)
}

// LoggerFactory is the interface that wraps Getlogger method.
type LoggerFactory interface {
	GetLogger(name string) (Logger, error)
	SetLogger(name string, l Logger) error
}
