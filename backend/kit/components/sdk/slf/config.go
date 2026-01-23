package slf

import (
	"errors"
	"fmt"
	"os"
	"path"
)

// Formatter defines the format of logging line.
//type Formatter string

const (
	// JSONFormatter prints logs in JSON format which looks like:
	// {"level":"info","mtime":"2019-06-03T20:27:11.748+0800","caller":"foo/bar.go:39","msg":"hello", "foo":"bar"}
	JSONFormatter string = "json"

	// TextFormatter prints logs in plain text format which looks like:
	// 2019-06-03T20:27:11.748+0800 info foo/bar.go:39 hello {"foo": "bar"}
	TextFormatter string = "text"

	// LogfmtFormatter prints logs in logfmt format which looks like:
	// mtime=2019-09-20T14:48:30.017+0800 level=info caller=foo/bar.go:39  msg="this a context log" traceid=xxxx-1234 uid=5678
	LogfmtFormatter string = "logfmt"

	defaultSyslogNetwork  = "udp"
	defaultSyslogAddr     = "127.0.0.1:514"
	defaultSyslogFacility = "local6"

	flagWithoutTime   = 1
	flagWithoutLevel  = (1 << 1)
	flagWithoutCaller = (1 << 2)
	flagWithoutMsg    = (1 << 3)

	callerKey = "caller"
	levelKey  = "level"
	timeKey   = "mtime"
	msgKey    = "msg"

	defaultTimeFormat = "2006-01-02T15:04:05.000Z0700"
)

// Output defines where logs are written to
//type Output string

const (
	// Discard discard all logs
	Discard string = "discard"
	// Stdout output logs to stdout
	Stdout string = "stdout"
	// Stderr output logs to stderr
	Stderr string = "stderr"
	// File output logs to files
	File string = "file"
	// Syslog output logs to syslog
	Syslog string = "syslog"
)

// SyslogConfig syslog configuration
type SyslogConfig struct {
	// Network defines the protocol used to talk to syslogd which should be
	// one of ["udp","tcp","unix"]. The default is "udp".
	Network string
	// Addr should be like "ip:port". The default is "127.0.0.1:514".
	Addr string

	// Facility is used to specify the type of program that is logging the message.
	// local0 - local7 are locally used facilities. The default is "local5".
	Facility string

	// Tag should be the name of the program or process that is logging the message.
	Tag string
}

// NewSyslogConfig creates a default syslog sdkconfig.
func NewSyslogConfig(tag string) *SyslogConfig {
	return &SyslogConfig{
		Network:  defaultSyslogNetwork,
		Addr:     defaultSyslogAddr,
		Facility: defaultSyslogFacility,
		Tag:      tag,
	}
}

// SetNetwork set protocol used to connect to syslogd
func (c *SyslogConfig) SetNetwork(net string) *SyslogConfig {
	c.Network = net
	return c
}

// SetAddr set syslogd address. e.g. "127.0.0.1:514"
func (c *SyslogConfig) SetAddr(addr string) *SyslogConfig {
	c.Addr = addr
	return c
}

// SetFacility set syslog facility.
func (c *SyslogConfig) SetFacility(f string) *SyslogConfig {
	c.Facility = f
	return c
}

// SetTag set syslog tag which should be the name of the program or process that
// is logging the message.
func (c *SyslogConfig) SetTag(t string) *SyslogConfig {
	c.Tag = t
	return c
}

// Valid checks if syslog sdkconfig is valid.
func (c *SyslogConfig) Valid() error {
	switch c.Network {
	case "udp", "tcp", "unix":
	default:
		return fmt.Errorf("invalid network: %s", c.Network)
	}
	switch c.Facility {
	case "kern", "user", "mail", "daemon", "auth", "syslog":
	case "lpr", "news", "uucp", "authpriv", "ftp", "cron":
	case "local0", "local1", "local2", "local3":
	case "local4", "local5", "local6", "local7":
	default:
		return fmt.Errorf("invalid facility: %s", c.Facility)
	}

	return nil
}

// FileConfig sdkconfig for logging into a file
type FileConfig struct {
	// FileBaseDir is the file dir to write log file to.
	//FileBaseDir string
	// FileName is the file to write logs to.
	// e.g. /some/dir/%Y%m%d/some.%level-log.%Y%m%d%H%M
	FileName string

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 MB.
	//MaxSize int

	// MaxBackups is the maximum number of old log files to retain. The default
	// is to retain all old log files though MaxAge may still cause them to get deleted.
	MaxBackups int

	// MaxAge is the maximum number of days to retain old log files based on the timestamp
	// encoded in their filename. The default is not to remove old log files based on age.
	MaxAge int64 // 设置文件清理前的最长保存时间
	//RotationDuration time.Duration // 日志切割时间间隔
}

// NewFileConfig creates a file sdkconfig.
func NewFileConfig(fileName string) *FileConfig {
	return &FileConfig{FileName: fileName}
}

// Valid checks if file sdkconfig is valid
func (c *FileConfig) Valid() error {
	if c.FileName == "" {
		return errors.New("invalid file name: nil")
	}
	return nil
}

// SetFileName set file name
func (c *FileConfig) SetFileName(name string) *FileConfig {
	c.FileName = name
	return c
}

type Config struct {
	Level  string
	Output string
	//Syslog    SyslogConfig
	File      FileConfig
	Formatter string

	callerSkip int
	stripFlags int
}

func toConfiguration(conf Config) (*Configuration, error) {
	c := &Configuration{}

	level, err := ToLevel(conf.Level)
	if err != nil {
		return nil, err
	}
	c.Level = level

	c.Output = conf.Output
	if c.Output == "" {
		c.Output = Syslog
	}

	//c.Syslog = &conf.Syslog
	c.File = &conf.File

	c.Formatter = conf.Formatter
	if c.Formatter == "" {
		c.Formatter = JSONFormatter
	}

	return c, nil
}

// Configuration defines logging sdkconfig.
type Configuration struct {
	Level     Level
	Output    string
	Syslog    *SyslogConfig
	File      *FileConfig
	Formatter string

	callerSkip int
	stripFlags int
}

// NewProductionConfig creates a production environment sdkconfig.
func NewProductionConfig() *Configuration {
	return &Configuration{
		Level:     LevelInfo,
		Output:    Syslog,
		Syslog:    NewSyslogConfig(path.Base(os.Args[0])),
		File:      nil,
		Formatter: JSONFormatter,

		callerSkip: 1,
	}
}

// NewDevelopmentConfig creates a development environment sdkconfig.
func NewDevelopmentConfig() *Configuration {
	return &Configuration{
		Level:     LevelDebug,
		Output:    Stdout,
		Syslog:    nil,
		File:      nil,
		Formatter: LogfmtFormatter,

		callerSkip: 1,
	}
}

// SetLevel set level
func (c *Configuration) SetLevel(l Level) *Configuration {
	c.Level = l
	return c
}

// SetOutput set output
func (c *Configuration) SetOutput(out string) *Configuration {
	c.Output = out
	return c
}

// SetSyslog set syslog sdkconfig.
func (c *Configuration) SetSyslog(conf *SyslogConfig) *Configuration {
	c.Syslog = conf
	return c
}

// SetFile set file sdkconfig.
func (c *Configuration) SetFile(conf *FileConfig) *Configuration {
	c.File = conf
	return c
}

// SetFormatter set formatter.
func (c *Configuration) SetFormatter(f string) *Configuration {
	c.Formatter = f
	return c
}

// WithoutTime makes log entry without mtime field.
func (c *Configuration) WithoutTime() *Configuration {
	c.stripFlags |= flagWithoutTime
	return c
}

func (c *Configuration) getTimeKey() string {
	if c.stripFlags&flagWithoutTime != 0 {
		return ""
	}
	return timeKey
}

// WithoutLevel makes log entry without level field.
func (c *Configuration) WithoutLevel() *Configuration {
	c.stripFlags |= flagWithoutLevel
	return c
}

func (c *Configuration) getLevelKey() string {
	if c.stripFlags&flagWithoutLevel != 0 {
		return ""
	}
	return levelKey
}

// WithoutCaller makes log entry without caller field.
func (c *Configuration) WithoutCaller() *Configuration {
	c.stripFlags |= flagWithoutCaller
	return c
}

func (c *Configuration) getCallerKey() string {
	if c.stripFlags&flagWithoutCaller != 0 {
		return ""
	}
	return callerKey
}

// WithoutMsg makes log entry without msg field.
func (c *Configuration) WithoutMsg() *Configuration {
	c.stripFlags |= flagWithoutMsg
	return c
}

func (c *Configuration) getMsgKey() string {
	if c.stripFlags&flagWithoutMsg != 0 {
		return ""
	}
	return msgKey
}

// Valid checks if sdkconfig is valid.
func (c *Configuration) Valid() error {
	switch c.Level {
	case LevelTrace, LevelDebug, LevelInfo, LevelWarn, LevelError, LevelFatal:
	default:
		return fmt.Errorf("invalid level: %+v", c.Level)
	}

	switch c.Output {
	case Discard, Stdout, Stderr, File, Syslog:
	default:
		return fmt.Errorf("invalid output: %s", c.Output)
	}

	switch c.Formatter {
	case JSONFormatter, TextFormatter, LogfmtFormatter:
	default:
		return fmt.Errorf("invalid formatter: %s", c.Formatter)
	}

	if c.Output == Syslog {
		if c.Syslog == nil {
			return fmt.Errorf("specifiy syslog as output but missing syslog sdkconfig")
		}
		if err := c.Syslog.Valid(); err != nil {
			return err
		}
	}

	if c.Output == File {
		if c.File == nil {
			return fmt.Errorf("specify file as output but missing file sdkconfig")
		}
		if err := c.File.Valid(); err != nil {
			return err
		}
	}

	return nil
}
