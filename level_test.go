package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true, // Enable full timestamp
	})

	logger.Trace("This is trace") // tidak keluar di console
	logger.Debug("This is debug") // tidak keluar di console
	logger.Info("This is info")
	logger.Warn("This is warn")
	logger.Error("This is error")
}

func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true, // Enable full timestamp
	})
	logger.SetLevel(logrus.WarnLevel)

	logger.Trace("This is trace") // tidak keluar di console
	logger.Debug("This is debug") // tidak keluar di console
	logger.Info("This is info")
	logger.Warn("This is warn")
	logger.Error("This is error")
}
