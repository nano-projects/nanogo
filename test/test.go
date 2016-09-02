package test

import (
	"github.com/Sirupsen/logrus"
	"testing"
)

func DebugMode(t *testing.T) {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		t.Error(err)
		return
	}

	logrus.SetLevel(level)
}
