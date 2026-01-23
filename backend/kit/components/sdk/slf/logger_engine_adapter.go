package slf

import (
	"fmt"
	"strings"
)

type LoggerAdapter struct {
	name  string
	level Level
}

func (a *LoggerAdapter) SetName(name string) {
	a.name = name
}

func (a *LoggerAdapter) GetName() string {
	return a.name
}

func (a *LoggerAdapter) GetLevel() Level {
	return a.level
}

func (a *LoggerAdapter) SetLevel(l Level) {
	a.level = l
}

func (a *LoggerAdapter) SetLevelString(l string) error {
	l = strings.ToLower(l)
	switch l {
	case levelTraceString:
		a.SetLevel(LevelTrace)
	case levelDebugString:
		a.SetLevel(LevelDebug)
	case levelInfoString:
		a.SetLevel(LevelInfo)
	case levelWarnString, levelWarningString:
		a.SetLevel(LevelWarn)
	case levelErrorString:
		a.SetLevel(LevelError)
	case levelFatalString:
		a.SetLevel(LevelFatal)
	default:
		return fmt.Errorf("unknown level string: %s", l)
	}
	return nil
}

func (a *LoggerAdapter) IsEnableTrace() bool {
	return a.level <= LevelTrace
}

func (a *LoggerAdapter) IsEnableDebug() bool {
	return a.level <= LevelDebug
}

func (a *LoggerAdapter) IsEnableInfo() bool {
	return a.level <= LevelInfo
}

func (a *LoggerAdapter) IsEnableWarn() bool {
	return a.level <= LevelWarn
}

func (a *LoggerAdapter) IsEnableError() bool {
	return a.level <= LevelError
}

func (a *LoggerAdapter) IsEnableFatal() bool {
	return a.level <= LevelFatal
}
