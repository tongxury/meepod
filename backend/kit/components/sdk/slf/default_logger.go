package slf

import (
	"context"
)

var (
	defLogger      Logger
	defFactory     LoggerFactory
	definedFactory LoggerFactory
)

const builtinLogger = "__builtin_logger__"

func init() {
	c := NewProductionConfig().SetOutput(Stdout)
	c.callerSkip = 2
	defFactory = NewLogrusLoggerFactory(c)
	defLogger, _ = defFactory.GetLogger(builtinLogger)
}

// Init initialize the default logger
func InitLogger(conf Config) (err error) {

	c, err := toConfiguration(conf)
	if err != nil {
		return
	}

	c.callerSkip = 2
	defFactory = NewLogrusLoggerFactory(c)
	defLogger, err = defFactory.GetLogger(builtinLogger)
	return err
}

// NewLogger create a new logger
func NewLogger(name string, c *Configuration) (Logger, error) {
	f := NewLogrusLoggerFactory(c)
	l, err := f.GetLogger(name)
	if err != nil {
		return nil, err
	}
	if definedFactory != nil {
		if err := definedFactory.SetLogger(name, l); err != nil {
			return nil, err
		}
	} else {
		if err := defFactory.SetLogger(name, l); err != nil {
			return nil, err
		}
	}
	return l, nil
}

// SetLoggerFactory set a logger factory
func SetLoggerFactory(factory LoggerFactory) {
	if factory == nil {
		panic("LoggerFactory can't be nil")
	}
	definedFactory = factory
}

// GetDefaultLogger get the default logger.
func GetDefaultLogger() (Logger, error) {
	return GetLogger(builtinLogger)
}

// GetLogger get a logger by name. Return the corresponading logger if name already exists,
// otherwise create a new one and save it.
func GetLogger(name string) (Logger, error) {
	if definedFactory != nil {
		return definedFactory.GetLogger(name)
	}
	return defFactory.GetLogger(name)
}

// Close close the logger
func Close() error {
	return defLogger.Close()
}

// SetLevel set logging level
func SetLevel(level Level) {
	defLogger.SetLevel(level)
}

// SetLevelString set logging level by literal names
func SetLevelString(level string) error {
	return defLogger.SetLevelString(level)
}

// AddContextHook adds a context hook.
func AddContextHook(h ContextHook) {
	defLogger.AddContextHook(h)
}

// ClearContextHooks clear all context hooks.
func ClearContextHooks() {
	defLogger.ClearContextHooks()
}

// WithError adds a error field
// It doesn't log unitl any of Debug(f/w)/Info(f/w)/Warn(f/w)/Error(f/w)/Fatal(f/w) methods is called.
func WithError(err error) Logger {
	return defLogger.AddCallerSkip(-1).WithError(err)
}

// WithField adds a field into the logger.
// It doesn't log unitl any of Debug(f/w)/Info(f/w)/Warn(f/w)/Error(f/w)/Fatal(f/w) methods is called.
func WithField(key string, val interface{}) Logger {
	return defLogger.AddCallerSkip(-1).WithField(key, val)
}

// WithFields adds multiple fields into the logger.
// It doesn't log unitl any of Debug(f/w)/Info(f/w)/Warn(f/w)/Error(f/w)/Fatal(f/w) methods is called.
func WithFields(fields ...Field) Logger {
	return defLogger.AddCallerSkip(-1).WithFields(fields...)
}

// WithContext extracts fields from context and adds that into the logger.
// Note that extracting fields are done through calling context hooks.
// It doesn't log unitl any of Debug(f/w)/Info(f/w)/Warn(f/w)/Error(f/w)/Fatal(f/w) methods is called.
func WithContext(ctx context.Context) Logger {
	return defLogger.AddCallerSkip(-1).WithContext(ctx)
}

// Infow logs a message at Info level. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Infow(msg string, fields ...Field) {
	defLogger.Infow(msg, fields...)
}

// Warnw logs a message at Warn level. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Warnw(msg string, fields ...Field) {
	defLogger.Warnw(msg, fields...)
}

// Errorw logs a message at Error level. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Errorw(msg string, fields ...Field) {
	defLogger.Errorw(msg, fields...)
}

// Fatalw logs a message at Fatal level. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Fatalw(msg string, fields ...Field) {
	defLogger.Fatalw(msg, fields...)
}

// Debugw logs a message at Debug level. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func Debugw(msg string, fields ...Field) {
	defLogger.Debugw(msg, fields...)
}

// Infoln logs arguments and any fields accumulated on the logger at Info level.
func Infoln(v ...interface{}) {
	defLogger.Infoln(v...)
}

// Infof logs arguments and any fields accumulated on the logger at Info level.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	defLogger.Infof(format, v...)
}

// Warnln logs arguments and any fields accumulated on the logger at Warn level.
func Warnln(v ...interface{}) {
	defLogger.Warnln(v...)
}

// Warnf logs arguments and any fields accumulated on the logger at Warn level.
// Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, v ...interface{}) {
	defLogger.Warnf(format, v...)
}

// Errorln logs arguments and any fields accumulated on the logger at Error level.
func Errorln(v ...interface{}) {
	defLogger.Errorln(v...)
}

// Errorf logs arguments and any fields accumulated on the logger at Error level.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	defLogger.Errorf(format, v...)
}

// Fatalln logs arguments and any fields accumulated on the logger at Fatal level.
func Fatalln(v ...interface{}) {
	defLogger.Fatalln(v...)
}

// Fatalf logs arguments and any fields accumulated on the logger at Fatal level.
// Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, v ...interface{}) {
	defLogger.Fatalf(format, v...)
}

// Debugln logs arguments and any fields accumulated on the logger at Debug level.
func Debugln(v ...interface{}) {
	defLogger.Debugln(v...)
}

// Debugf logs arguments and any fields accumulated on the logger at Debug level.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	defLogger.Debugf(format, v...)
}

// Debugc logs arguments and fields extracted from context by hooks and accumulated on the logger at Debug level.
func Debugc(ctx context.Context, msg string, fields ...Field) {
	defLogger.Debugc(ctx, msg, fields...)
}

// Infoc logs arguments and fields extracted from context by hooks and accumulated on the logger at Info level.
func Infoc(ctx context.Context, msg string, fields ...Field) {
	defLogger.Infoc(ctx, msg, fields...)
}

// Warnc logs arguments and fields extracted from context by hooks and accumulated on the logger at Warn level.
func Warnc(ctx context.Context, msg string, fields ...Field) {
	defLogger.Warnc(ctx, msg, fields...)
}

// Errorc logs arguments and fields extracted from context by hooks and accumulated on the logger at Error level.
func Errorc(ctx context.Context, msg string, fields ...Field) {
	defLogger.Errorc(ctx, msg, fields...)
}

// Fatalc logs arguments and fields extracted from context by hooks and accumulated on the logger at Fatal level.
func Fatalc(ctx context.Context, msg string, fields ...Field) {
	defLogger.Fatalc(ctx, msg, fields...)
}
