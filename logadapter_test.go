package logadapter

import (
	"errors"
	"log/slog"
	"os"
	"testing"
)

func TestLogadapter(t *testing.T) {
	log := LogAdapter(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

	log.Info("hi")
	log.With("err", errors.New("uh oh"), "ps", "test", "num", 123).Warn("warnnn..")
	log.Info("yes")

	log.Warnf("failed due to %v", errors.New("uh oh"))
	log.Info()
	log.Info("done", "key", "value", "num", 123)
	log.Info("done", slog.String("key", "value"), slog.Int("num", 123))
}
