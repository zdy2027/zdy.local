package utils

import (
	"testing"
)

type fileLog struct {
	Filename string
}

func TestLogs(t *testing.T){
	FileLogs.Info("info test")
	FileLogs.Error("error test")
	FileLogs.Alert("alert test")
	FileLogs.Debug("debug test")
	FileLogs.Critical("critical test")
	FileLogs.Emergency("emergency test")
	//FileLogs.Trace("trace test")
	//FileLogs.Warn("warn test")
	FileLogs.Notice("notice test")
	//FileLogs.Informational("information test")
	FileLogs.Warning("warning test")
}
