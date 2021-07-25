package services

import (
	"log"
	"os"
	"testing"
)

type nilLogger struct {
}

func (l nilLogger) Write(p []byte) (n int, err error) {
	return 0, nil
}
func TestMain(m *testing.M) {
	log.SetOutput(&nilLogger{})
	os.Exit(m.Run())
}
