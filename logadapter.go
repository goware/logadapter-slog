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
	s.log.Debug(fmt.Sprint(v...))
}

func (s *logAdapter) Debugf(format string, v ...interface{}) {
	s.log.Debug(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Info(v ...interface{}) {
	s.log.Info(fmt.Sprint(v...))
}

func (s *logAdapter) Infof(format string, v ...interface{}) {
	s.log.Info(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Warn(v ...interface{}) {
	s.log.Warn(fmt.Sprint(v...))
}

func (s *logAdapter) Warnf(format string, v ...interface{}) {
	s.log.Warn(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Error(v ...interface{}) {
	s.log.Error(fmt.Sprint(v...))
}

func (s *logAdapter) Errorf(format string, v ...interface{}) {
	s.log.Error(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Fatal(v ...interface{}) {
	s.log.Error(fmt.Sprint(v...))
	os.Exit(1)
}

func (s *logAdapter) Fatalf(format string, v ...interface{}) {
	s.log.Error(fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (s *logAdapter) Print(v ...interface{}) {
	s.log.Info(fmt.Sprint(v...))
}

func (s *logAdapter) Println(v ...interface{}) {
	s.log.Info(fmt.Sprint(v...))
}

func (s *logAdapter) Printf(format string, v ...interface{}) {
	s.log.Info(fmt.Sprintf(format, v...))
}
