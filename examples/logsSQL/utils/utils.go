package utils

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"encoding/json"
)

var ConsoleLogs *logs.BeeLogger
var FileLogs *logs.BeeLogger

type configs struct {
	Filename string
	Separate []string
}

func init(){
	//beego.LoadAppConfig("ini","../conf/logs.conf")
	cfg := beego.AppConfig

	ConsoleLogs = logs.NewLogger(1000)
	ConsoleLogs.SetLogger("console")
	level,_ := cfg.Int("consul")
	ConsoleLogs.SetLevel(level)
	ConsoleLogs.EnableFuncCallDepth(true)

	conf := &configs{Filename:cfg.String("logdir"),Separate:cfg.Strings("separate")}
	result,_ := json.Marshal(conf)
	FileLogs = logs.NewLogger(1000)
	FileLogs.SetLogger(logs.AdapterMultiFile, string(result))
	level,_ = cfg.Int("file")
	FileLogs.SetLevel(level)
	FileLogs.EnableFuncCallDepth(true)
	FileLogs.Async()
}

func CloseLogs()  {
	ConsoleLogs.Close()
	FileLogs.Close()
}