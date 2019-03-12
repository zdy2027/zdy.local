package main

import (
	"zdy.local/logsSQL/fileOP"
)

func main() {
	var obj fileOP.FileOP
	obj = new(fileOP.LogsOP)
	obj.Init()
	input := []string{"../input"}
	fileOP.Init(input,"10009",obj)
}