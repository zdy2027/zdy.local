package main

import (
	"zdy.local/utils/nlogs"
)

func test_nlog() {
	nlogs.ConsoleLogs.Info("info test")
	nlogs.FileLogs.Error("warn test")
	nlogs.CloseLogs()
}

func main() {
	test_nlog()
}