package slf

import (
	"context"
	"fmt"
	"time"

	"gitee.com/meepo/backend/kit/components/sdk/slf/metrics"
	"github.com/sirupsen/logrus"

	//"gitee.com/meepo/backend/kit/components/sdk/slf/utils/syslog"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"

	//"github.com/sirupsen/logrus"
	//sysloghook "github.com/sirupsen/logrus/hooks/syslog"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
)

// LogrusLoggerAdapter is the adaptor based on logrus
type LogrusLoggerAdapter struct {
	LoggerAdapter
	entry        *logrus.Entry
	contextHooks ContextHooks
	file         *os.File
	callerSkip   int
	metrics      *metrics.LogMetrics
}

func newLogrusLogger(name string, c *Configuration) (*LogrusLoggerAdapter, error) {
	if err := c.Valid(); err != nil {
		return nil, err
	}

	l := new(LogrusLoggerAdapter)
	l.metrics = metrics.GetLogMetrics()
	l.SetName(name)
	logger := logrus.New()
	//logger.SetReportCaller(true)

	var formatter logrus.Formatter
	switch c.Formatter {
	case JSONFormatter:
		formatter = &logrus.JSONFormatter{
			TimestampFormat: defaultTimeFormat,
		}
	case TextFormatter, LogfmtFormatter:
		formatter = &logrus.TextFormatter{
			TimestampFormat: defaultTimeFormat,
		}
	default:
		formatter = &logrus.JSONFormatter{
			TimestampFormat: defaultTimeFormat,
		}
	}
	logger.SetFormatter(formatter)

	switch c.Output {
	case Discard:
		logger.SetOutput(ioutil.Discard)
	case Stdout:
		logger.SetOutput(os.Stdout)
	case Stderr:
		logger.SetOutput(os.Stderr)
	case File:

		if c.File == nil || c.File.FileName == "" {
			return nil, fmt.Errorf("check your log sdkconfig")
		}

		writerMap := lfshook.WriterMap{}

		maxAgeOption := rotatelogs.WithMaxAge(time.Duration(c.File.MaxAge) * time.Second) // 文件最大保存时间
		//rotationCountOption := rotatelogs.WithRotationCount(uint(c.File.MaxBackups))      // 日志最大数量

		levelFormat := "%level"

		if strings.Contains(c.File.FileName, levelFormat) {

			var err error
			if writerMap[logrus.DebugLevel], err = rotatelogs.New(
				strings.ReplaceAll(c.File.FileName, levelFormat, "debug"),
				maxAgeOption,
			); err != nil {
				return nil, err
			}

			if writerMap[logrus.InfoLevel], err = rotatelogs.New(
				strings.ReplaceAll(c.File.FileName, levelFormat, "info"),
				maxAgeOption,
			); err != nil {
				return nil, err
			}

			if writerMap[logrus.WarnLevel], err = rotatelogs.New(
				strings.ReplaceAll(c.File.FileName, levelFormat, "warn"),
				maxAgeOption,
			); err != nil {
				return nil, err
			}

			errorWriter, err := rotatelogs.New(
				strings.ReplaceAll(c.File.FileName, levelFormat, "error"),
				maxAgeOption,
			)
			if err != nil {
				return nil, err
			}
			writerMap[logrus.ErrorLevel] = errorWriter
			writerMap[logrus.FatalLevel] = errorWriter
			writerMap[logrus.PanicLevel] = errorWriter

		} else {
			writer, err := rotatelogs.New(c.File.FileName, maxAgeOption)

			if err != nil {
				return nil, err
			}

			writerMap[logrus.DebugLevel] = writer
			writerMap[logrus.InfoLevel] = writer
			writerMap[logrus.WarnLevel] = writer
			writerMap[logrus.ErrorLevel] = writer
			writerMap[logrus.FatalLevel] = writer
			writerMap[logrus.PanicLevel] = writer
		}

		logger.AddHook(lfshook.NewHook(writerMap, formatter))

	case Syslog:
		if c.Syslog.Tag == "" {
			c.Syslog.Tag = path.Base(os.Args[0])
		}
		logger.SetOutput(ioutil.Discard)
		/*hook, err := sysloghook.NewSyslogHook(
			c.Syslog.Network,
			c.Syslog.Addr,
			syslog.GetPriority(c.Syslog.Facility),
			c.Syslog.Tag)
		if err != nil {
			return nil, err
		}
		logger.Hooks.Add(hook)
		*/
	}
	l.entry = logrus.NewEntry(logger)
	l.SetLevel(c.Level)
	callerSkip := 1
	if c.callerSkip > 1 {
		callerSkip = c.callerSkip
	}
	l.callerSkip = callerSkip
	return l, nil
}

// Close logger can do some cleanup jobs such as flushing any buffered log entries.
func (l *LogrusLoggerAdapter) Close() error {
	//if l.file != nil {
	//	return l.file.Close()
	//}
	return nil
}

// SetLevel set logging level
func (l *LogrusLoggerAdapter) SetLevel(level Level) {
	l.LoggerAdapter.SetLevel(level)
	switch level {
	case LevelTrace:
		l.entry.Logger.SetLevel(logrus.DebugLevel)
	case LevelDebug:
		l.entry.Logger.SetLevel(logrus.DebugLevel)
	case LevelInfo:
		l.entry.Logger.SetLevel(logrus.InfoLevel)
	case LevelWarn:
		l.entry.Logger.SetLevel(logrus.WarnLevel)
	case LevelError:
		l.entry.Logger.SetLevel(logrus.ErrorLevel)
	case LevelFatal:
		l.entry.Logger.SetLevel(logrus.FatalLevel)
	}
}

// AddContextHook adds a context hook.
func (l *LogrusLoggerAdapter) AddContextHook(h ContextHook) {
	l.contextHooks = append(l.contextHooks, h)
}

// ClearContextHooks clear all context hooks.
func (l *LogrusLoggerAdapter) ClearContextHooks() {
	l.contextHooks = nil
}

// AddCallerSkip create a child logger with a incresed caller skip
func (l *LogrusLoggerAdapter) AddCallerSkip(skip int) Logger {
	if skip == 0 {
		return l
	}
	return &LogrusLoggerAdapter{
		LoggerAdapter: l.LoggerAdapter,
		entry:         l.entry.WithFields(nil),
		contextHooks:  append(l.contextHooks[:0:0], l.contextHooks...),
		file:          l.file,
		callerSkip:    l.callerSkip + skip,
		metrics:       l.metrics,
	}
}

func caller(skip int) string {
	const callerOffset = 1
	_, file, line, _ := runtime.Caller(skip + callerOffset)
	idx := strings.LastIndexByte(file, '/')
	if idx == -1 {
		return fmt.Sprintf("%s:%d", file, line)
	}
	idx = strings.LastIndexByte(file[:idx], '/')
	if idx == -1 {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return fmt.Sprintf("%s:%d", file[idx+1:], line)
}

// WithError adds an error as a field into logger.
// It doesn't log unitl anyone of Debug(f/w)/Info(f/w)/Warn(f/w)/Error(f/w)/Fatal(f/w) methods is called.
func (l *LogrusLoggerAdapter) WithError(err error) Logger {
	if err == nil {
		return l
	}
	return &LogrusLoggerAdapter{
		LoggerAdapter: l.LoggerAdapter,
		entry:         l.entry.WithError(err),
		contextHooks:  append(l.contextHooks[:0:0], l.contextHooks...),
		file:          l.file,
		callerSkip:    l.callerSkip,
		metrics:       l.metrics,
	}
}

// WithField adds a field into the logger.
// It doesn't log unitl anyone of Debug(f/w)/Info(f/w)/Warn(f/w)/Error(f/w)/Fatal(f/w) methods is called.
func (l *LogrusLoggerAdapter) WithField(key string, val interface{}) Logger {
	if key == "" {
		return l
	}
	return &LogrusLoggerAdapter{
		LoggerAdapter: l.LoggerAdapter,
		entry:         l.entry.WithField(key, val),
		contextHooks:  append(l.contextHooks[:0:0], l.contextHooks...),
		file:          l.file,
		callerSkip:    l.callerSkip,
		metrics:       l.metrics,
	}
}

func mapifyFields(fields ...Field) logrus.Fields {
	m := make(logrus.Fields, len(fields))
	for _, f := range fields {
		m[f.Key] = Value(&f)
	}
	return m
}

// WithFields adds multiple fields into the logger
// It doesn't log unitl anyone of Debug(f/w)/Info(f/w)/Warn(f/w)/Error(f/w)/Fatal(f/w) methods is called.
func (l *LogrusLoggerAdapter) WithFields(fields ...Field) Logger {
	if len(fields) == 0 {
		return l
	}
	return &LogrusLoggerAdapter{
		LoggerAdapter: l.LoggerAdapter,
		entry:         l.entry.WithFields(mapifyFields(fields...)),
		contextHooks:  append(l.contextHooks[:0:0], l.contextHooks...),
		file:          l.file,
		callerSkip:    l.callerSkip,
		metrics:       l.metrics,
	}
}

func (l *LogrusLoggerAdapter) execContextHooks(ctx context.Context) []Field {
	if ctx == nil || len(l.contextHooks) == 0 {
		return nil
	}
	var fields []Field
	for _, h := range l.contextHooks {
		fields = append(fields, h(ctx)...)
	}
	return fields
}

// WithContext extracts fields from context and adds that into the logger.
// Note that extracting fields are done through calling context hooks.
// It doesn't log unitl anyone of Debug(f/w)/Info(f/w)/Warn(f/w)/Error(f/w)/Fatal(f/w) methods is called.
func (l *LogrusLoggerAdapter) WithContext(ctx context.Context) Logger {
	fields := l.execContextHooks(ctx)
	if len(fields) == 0 {
		return l
	}
	return &LogrusLoggerAdapter{
		LoggerAdapter: l.LoggerAdapter,
		entry:         l.entry.WithFields(mapifyFields(fields...)),
		contextHooks:  append(l.contextHooks[:0:0], l.contextHooks...),
		file:          l.file,
		callerSkip:    l.callerSkip,
		metrics:       l.metrics,
	}
}

// Infow logs a message at Info level. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func (l *LogrusLoggerAdapter) Infow(msg string, fields ...Field) {
	if l.IsEnableInfo() {
		l.metrics.LogEntry(levelInfoString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).WithFields(mapifyFields(fields...)).Infoln(msg)
	}
}

// Warnw logs a message at Warn level. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func (l *LogrusLoggerAdapter) Warnw(msg string, fields ...Field) {
	if l.IsEnableWarn() {
		l.metrics.LogEntry(levelWarnString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).WithFields(mapifyFields(fields...)).Warnln(msg)
	}
}

// Errorw logs a message at Error level. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func (l *LogrusLoggerAdapter) Errorw(msg string, fields ...Field) {
	if l.IsEnableError() {
		l.metrics.LogEntry(levelErrorString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).WithFields(mapifyFields(fields...)).Errorln(msg)
	}
}

// Fatalw logs a message at Fatal level. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func (l *LogrusLoggerAdapter) Fatalw(msg string, fields ...Field) {
	if l.IsEnableFatal() {
		l.metrics.LogEntry(levelFatalString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).WithFields(mapifyFields(fields...)).Fatalln(msg)
	}
}

// Debugw logs a message at Debug level. The message includes any fields passed at the log site,
// as well as any fields accumulated on the logger.
func (l *LogrusLoggerAdapter) Debugw(msg string, fields ...Field) {
	if l.IsEnableDebug() {
		l.metrics.LogEntry(levelDebugString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).WithFields(mapifyFields(fields...)).Debugln(msg)
	}
}

// Infoln logs arguments and any fields accumulated on the logger at Info level.
func (l *LogrusLoggerAdapter) Infoln(v ...interface{}) {
	if l.IsEnableInfo() {
		l.metrics.LogEntry(levelInfoString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).Infoln(v...)
	}
}

// Infof logs arguments and any fields accumulated on the logger at Info level.
// Arguments are handled in the manner of fmt.Printf.
func (l *LogrusLoggerAdapter) Infof(format string, v ...interface{}) {
	if l.IsEnableInfo() {
		l.metrics.LogEntry(levelInfoString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).Infof(format, v...)
	}
}

// Warnln logs arguments and any fields accumulated on the logger at Warn level.
func (l *LogrusLoggerAdapter) Warnln(v ...interface{}) {
	if l.IsEnableWarn() {
		l.metrics.LogEntry(levelWarnString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).Warnln(v...)
	}
}

// Warnf logs arguments and any fields accumulated on the logger at Warn level.
// Arguments are handled in the manner of fmt.Printf.
func (l *LogrusLoggerAdapter) Warnf(format string, v ...interface{}) {
	if l.IsEnableWarn() {
		l.metrics.LogEntry(levelWarnString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).Warnf(format, v...)
	}
}

// Errorln logs arguments and any fields accumulated on the logger at Error level.
func (l *LogrusLoggerAdapter) Errorln(v ...interface{}) {
	if l.IsEnableError() {
		l.metrics.LogEntry(levelErrorString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).Errorln(v...)
	}
}

// Errorf logs arguments and any fields accumulated on the logger at Error level.
// Arguments are handled in the manner of fmt.Printf.
func (l *LogrusLoggerAdapter) Errorf(format string, v ...interface{}) {
	if l.IsEnableError() {
		l.metrics.LogEntry(levelErrorString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).Errorf(format, v...)
	}
}

// Fatalln logs arguments and any fields accumulated on the logger at Fatal level.
func (l *LogrusLoggerAdapter) Fatalln(v ...interface{}) {
	if l.IsEnableFatal() {
		l.metrics.LogEntry(levelFatalString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).Fatalln(v...)
	}
}

// Fatalf logs arguments and any fields accumulated on the logger at Fatal level.
// Arguments are handled in the manner of fmt.Printf.
func (l *LogrusLoggerAdapter) Fatalf(format string, v ...interface{}) {
	if l.IsEnableFatal() {
		l.metrics.LogEntry(levelFatalString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).Fatalf(format, v...)
	}
}

// Debugln logs arguments and any fields accumulated on the logger at Debug level.
func (l *LogrusLoggerAdapter) Debugln(v ...interface{}) {
	if l.IsEnableDebug() {
		l.metrics.LogEntry(levelDebugString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).Debugln(v...)
	}
}

// Debugf logs arguments and any fields accumulated on the logger at Debug level.
// Arguments are handled in the manner of fmt.Printf.
func (l *LogrusLoggerAdapter) Debugf(format string, v ...interface{}) {
	if l.IsEnableDebug() {
		l.metrics.LogEntry(levelDebugString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).Debugf(format, v...)
	}
}

// Debugc logs arguments and fields extracted from context by hooks and accumulated on the logger at Debug level.
func (l *LogrusLoggerAdapter) Debugc(ctx context.Context, msg string, fields ...Field) {
	if l.IsEnableDebug() {
		l.metrics.LogEntry(levelDebugString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).WithFields(mapifyFields(append(fields, l.execContextHooks(ctx)...)...)).Debugln(msg)
	}
}

// Infoc logs arguments and fields extracted from context by hooks and accumulated on the logger at Info level.
func (l *LogrusLoggerAdapter) Infoc(ctx context.Context, msg string, fields ...Field) {
	if l.IsEnableInfo() {
		l.metrics.LogEntry(levelInfoString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).WithFields(mapifyFields(append(fields, l.execContextHooks(ctx)...)...)).Infoln(msg)
	}
}

// Warnc logs arguments and fields extracted from context by hooks and accumulated on the logger at Warning level.
func (l *LogrusLoggerAdapter) Warnc(ctx context.Context, msg string, fields ...Field) {
	if l.IsEnableWarn() {
		l.metrics.LogEntry(levelWarnString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).WithFields(mapifyFields(append(fields, l.execContextHooks(ctx)...)...)).Warnln(msg)
	}
}

// Errorc logs arguments and fields extracted from context by hooks and accumulated on the logger at Error level.
func (l *LogrusLoggerAdapter) Errorc(ctx context.Context, msg string, fields ...Field) {
	if l.IsEnableError() {
		l.metrics.LogEntry(levelErrorString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).WithFields(mapifyFields(append(fields, l.execContextHooks(ctx)...)...)).Errorln(msg)
	}
}

// Fatalc logs arguments and fields extracted from context by hooks and accumulated on the logger at Fatal level.
func (l *LogrusLoggerAdapter) Fatalc(ctx context.Context, msg string, fields ...Field) {
	if l.IsEnableFatal() {
		l.metrics.LogEntry(levelFatalString)
		l.entry.WithField(callerKey, caller(l.callerSkip)).WithFields(mapifyFields(append(fields, l.execContextHooks(ctx)...)...)).Fatalln(msg)
	}
}

// LogrusLoggerFactory is the factory of logrus logger.
type LogrusLoggerFactory struct {
	config  Configuration
	mux     *sync.RWMutex
	loggers map[string]Logger
}

// NewLogrusLoggerFactory creates a LoggerFactory building logrus logger.
func NewLogrusLoggerFactory(c *Configuration) LoggerFactory {
	return &LogrusLoggerFactory{
		config:  *c,
		mux:     new(sync.RWMutex),
		loggers: make(map[string]Logger),
	}
}

// GetLogger create a logrus logger.
func (factory *LogrusLoggerFactory) GetLogger(name string) (Logger, error) {
	factory.mux.RLock()
	l, ok := factory.loggers[name]
	factory.mux.RUnlock()
	if ok {
		return l, nil
	}
	l, err := newLogrusLogger(name, &factory.config)
	if err != nil {
		return nil, err
	}
	factory.mux.Lock()
	factory.loggers[name] = l
	factory.mux.Unlock()
	return l, nil
}

// SetLogger save a logger
func (factory *LogrusLoggerFactory) SetLogger(name string, l Logger) error {
	factory.mux.Lock()
	factory.loggers[name] = l
	factory.mux.Unlock()
	return nil
}
