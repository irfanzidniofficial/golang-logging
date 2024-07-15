package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {

	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Wold!")
	logger.WithField("username", "irfan").Info("Hello Wold!")

	logger.WithField("username", "irfan").
		WithField("name", "Irfan Zidni").
		Info("Hello Wold!")

}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "irfan",
		"name":     "Irfan Zidni",
	}).Info("Hello world")

}
