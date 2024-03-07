package log

import (
	"testing"
)

func TestInit(t *testing.T) {
	Init()
	SugarLogger.Info("log test successfully")
}
