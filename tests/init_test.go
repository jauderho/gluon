package tests

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	if level, err := logrus.ParseLevel(os.Getenv("GOMSRV_LOG_LEVEL")); err == nil {
		logrus.SetLevel(level)
	}
}
