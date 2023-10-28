package logadapter

import (
	"errors"
	"log/slog"
	"testing"
)

func TestLogadapter(t *testing.T) {
	log := LogAdapter(slog.Default())

	log.Info("hi")
	log.With("err", errors.New("uh oh"), "ps", "test", "num", 123).Warn("warnnn..")
	log.Info("yes")
}
