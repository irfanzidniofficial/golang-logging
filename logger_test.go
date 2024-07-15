package golang_logging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true, // Enable full timestamp
	})
	logger.SetLevel(logrus.InfoLevel) // Set the log level to info

	logger.Println("hello logger")

	fmt.Println("Hello logger")
}
