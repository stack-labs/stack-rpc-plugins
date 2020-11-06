package logrus

import (
	"github.com/sirupsen/logrus"
	"github.com/stack-labs/stack-rpc/logger"
)

type Options struct {
	logger.Options
	Formatter logrus.Formatter
	// Flag for whether to log caller info (off by default)
	ReportCaller bool
	SplitLevel   bool
	// Exit Function to call when FatalLevel log
	ExitFunc func(int)
}

type formatterKey struct{}
type splitLevelKey struct{}
type reportCallerKey struct{}
type exitKey struct{}

func TextFormatter(formatter *logrus.TextFormatter) logger.Option {
	return logger.SetOption(formatterKey{}, formatter)
}

func JSONFormatter(formatter *logrus.JSONFormatter) logger.Option {
	return logger.SetOption(formatterKey{}, formatter)
}

func ExitFunc(exit func(int)) logger.Option {
	return logger.SetOption(exitKey{}, exit)
}

func SplitLevel(s bool) logger.Option {
	return logger.SetOption(splitLevelKey{}, s)
}

// warning to use this option. because logrus doest not open CallerDepth option
// this will only print this package
func ReportCaller() logger.Option {
	return logger.SetOption(reportCallerKey{}, true)
}
