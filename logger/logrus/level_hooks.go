package logrus

import (
	"github.com/stack-labs/stack-rpc/logger"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"

	ls "github.com/sirupsen/logrus"
)

func prepareLevelHooks(opts logger.PersistenceOptions, l ls.Level) ls.LevelHooks {
	hooks := make(ls.LevelHooks)
	for _, level := range ls.AllLevels {
		if level >= l {
			hooks[level] = []ls.Hook{
				&PersistenceLevelHook{
					Writer: &lumberjack.Logger{
						Filename:   opts.Dir + pathSeparator + level.String(),
						MaxSize:    500, // megabytes
						MaxBackups: 3,
						MaxAge:     28,   //days
						Compress:   true, // disabled by default
					},
					Fired:  true,
					levels: []ls.Level{level},
				},
			}
		}
	}

	return hooks
}

type PersistenceLevelHook struct {
	Writer io.Writer
	Fired  bool
	levels []ls.Level
}

func (hook *PersistenceLevelHook) Levels() []ls.Level {
	return hook.levels
}

func (hook *PersistenceLevelHook) Fire(entry *ls.Entry) error {
	hook.Fired = true
	return nil
}
