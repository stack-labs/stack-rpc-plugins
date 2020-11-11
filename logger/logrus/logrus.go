package logrus

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/stack-labs/stack-rpc/logger"
)

var (
	pathSeparator = string(os.PathSeparator)
)

type entryLogger interface {
	WithFields(fields logrus.Fields) *logrus.Entry
	WithError(err error) *logrus.Entry

	Log(level logrus.Level, args ...interface{})
	Logf(level logrus.Level, format string, args ...interface{})
}

type logrusLogger struct {
	Logger entryLogger
	opts   Options
}

func (l *logrusLogger) Init(opts ...logger.Option) error {
	for _, o := range opts {
		o(&l.opts.Options)
	}

	if formatter, ok := l.opts.Context.Value(formatterKey{}).(logrus.Formatter); ok {
		l.opts.Formatter = formatter
	}

	if caller, ok := l.opts.Context.Value(reportCallerKey{}).(bool); ok && caller {
		l.opts.ReportCaller = caller
	}

	if exitFunction, ok := l.opts.Context.Value(exitKey{}).(func(int)); ok {
		l.opts.ExitFunc = exitFunction
	}

	if splitLevel, ok := l.opts.Context.Value(splitLevelKey{}).(bool); ok {
		l.opts.SplitLevel = splitLevel
	}

	if l.opts.Persistence != nil && l.opts.Persistence.Enable && l.opts.Out == nil {
		var dir = l.opts.Persistence.Dir
		if dir == "" {
			// todo error or impossible
			dir, _ = os.Getwd()
			dir += fmt.Sprintf("%s%s", pathSeparator, "logs")
		}

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				logger.Errorf("create logs dir err: %s", err)
			}
		}

		l.opts.Persistence.Dir = dir

		if l.opts.Persistence.BackupDir == "" {
			l.opts.Persistence.BackupDir = fmt.Sprintf("%s%s%s", dir, pathSeparator, "backup")
			if _, err := os.Stat(l.opts.Persistence.BackupDir); os.IsNotExist(err) {
				err = os.MkdirAll(l.opts.Persistence.BackupDir, os.ModePerm)
				if err != nil {
					logger.Errorf("create backup dir err: %s", err)
				}
			}
		}
	}

	log := logrus.New() // defaults
	log.SetLevel(fromStackLevel(l.opts.Level))
	log.SetOutput(l.opts.Out)
	log.SetFormatter(l.opts.Formatter)
	log.SetReportCaller(l.opts.ReportCaller)
	log.ExitFunc = l.opts.ExitFunc
	if l.opts.SplitLevel {
		// Send all logs to nowhere by default
		logger.Infof("split log into different level files")
		log.SetOutput(ioutil.Discard)
		log.ReplaceHooks(prepareLevelHooks(*l.opts.Persistence, log.Level))
	}
	l.Logger = log

	return nil
}

func (l *logrusLogger) String() string {
	// stack-rpc-logrus
	return "slogrus"
}

func (l *logrusLogger) Fields(fields map[string]interface{}) logger.Logger {
	return &logrusLogger{l.Logger.WithFields(fields), l.opts}
}

func (l *logrusLogger) Log(level logger.Level, args ...interface{}) {
	l.Logger.Log(fromStackLevel(level), args...)
}

func (l *logrusLogger) Logf(level logger.Level, format string, args ...interface{}) {
	l.Logger.Logf(fromStackLevel(level), format, args...)
}

func (l *logrusLogger) Options() logger.Options {
	return l.opts.Options
}

// New builds a new logger based on options
func NewLogger(opts ...logger.Option) logger.Logger {
	// Default options
	options := Options{
		Options: logger.Options{
			Level:   logger.InfoLevel,
			Fields:  make(map[string]interface{}),
			Context: context.Background(),
		},
		Formatter:    new(logrus.TextFormatter),
		ReportCaller: false,
		ExitFunc:     os.Exit,
	}

	l := &logrusLogger{opts: options}
	err := l.Init(opts...)
	if err != nil {
		logger.Errorf("init logrus err: %s", err)
	}
	return l
}

func fromStackLevel(level logger.Level) logrus.Level {
	switch level {
	case logger.TraceLevel:
		return logrus.TraceLevel
	case logger.DebugLevel:
		return logrus.DebugLevel
	case logger.InfoLevel:
		return logrus.InfoLevel
	case logger.WarnLevel:
		return logrus.WarnLevel
	case logger.ErrorLevel:
		return logrus.ErrorLevel
	case logger.FatalLevel:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}

func toStackLevel(level logrus.Level) logger.Level {
	switch level {
	case logrus.TraceLevel:
		return logger.TraceLevel
	case logrus.DebugLevel:
		return logger.DebugLevel
	case logrus.InfoLevel:
		return logger.InfoLevel
	case logrus.WarnLevel:
		return logger.WarnLevel
	case logrus.ErrorLevel:
		return logger.ErrorLevel
	case logrus.FatalLevel:
		return logger.FatalLevel
	default:
		return logger.InfoLevel
	}
}
