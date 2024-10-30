package logadapter

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/goware/logger"
)

func LogAdapter(logger *slog.Logger) logger.Logger {
	return &logAdapter{logger}
}

type logAdapter struct {
	log *slog.Logger
}

var _ logger.Logger = &logAdapter{}

func (s *logAdapter) With(args ...interface{}) logger.Logger {
	return &logAdapter{s.log.With(args...)}
}

func (s *logAdapter) Debug(v ...interface{}) {
	msg, args := getArgs(v...)
	s.log.Debug(msg, args...)
}

func (s *logAdapter) Debugf(format string, v ...interface{}) {
	s.log.Debug(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Info(v ...interface{}) {
	msg, args := getArgs(v...)
	s.log.Info(msg, args...)
}

func (s *logAdapter) Infof(format string, v ...interface{}) {
	s.log.Info(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Warn(v ...interface{}) {
	msg, args := getArgs(v...)
	s.log.Warn(msg, args...)
}

func (s *logAdapter) Warnf(format string, v ...interface{}) {
	s.log.Warn(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Error(v ...interface{}) {
	msg, args := getArgs(v...)
	s.log.Error(msg, args...)
}

func (s *logAdapter) Errorf(format string, v ...interface{}) {
	s.log.Error(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Fatal(v ...interface{}) {
	msg, args := getArgs(v...)
	s.log.Error(msg, args...)
	os.Exit(1)
}

func (s *logAdapter) Fatalf(format string, v ...interface{}) {
	s.log.Error(fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (s *logAdapter) Print(v ...interface{}) {
	msg, args := getArgs(v...)
	s.log.Info(msg, args...)
}

func (s *logAdapter) Println(v ...interface{}) {
	msg, args := getArgs(v...)
	s.log.Info(msg, args...)
}

func (s *logAdapter) Printf(format string, v ...interface{}) {
	s.log.Info(fmt.Sprintf(format, v...))
}

// getArgs returns the message and arguments from the variadic arguments.
func getArgs(v ...interface{}) (string, []interface{}) {
	if len(v) == 0 {
		return "", nil
	}
	if len(v) == 1 {
		return fmt.Sprint(v[0]), nil
	}
	// validate that the first argument is a string
	msg, ok := v[0].(string)
	if !ok {
		return fmt.Sprint(v...), nil
	}
	// validate that the rest of the arguments are Attr
	for i := 1; i < len(v); i++ {
		if _, ok := v[i].(slog.Attr); !ok {
			return fmt.Sprint(v...), nil
		}
	}
	return msg, v[1:]
}
