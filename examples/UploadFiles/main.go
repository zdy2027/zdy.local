package main

import (
	"zdy.local/fileOP"
	"github.com/astaxie/beego/config"
	"./upload"
	"zdy.local/utils/nlogs"
)

func main() {
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	var obj fileOP.FileOP
	obj = new(upload.Upload)
	obj.Init("")
	input := cfg.String("input")
	nlogs.ConsoleLogs.Debug(input)
	fileOP.Run(input,"",obj)
}
