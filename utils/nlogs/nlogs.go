package nlogs

import (
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/config"
)

var ConsoleLogs *logs.BeeLogger
var FileLogs *logs.BeeLogger

type configs struct {
	Filename string 	`json:"filename"`
	Separate []string 	`json:"separate,omitempty"`
}

func init(){
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		ConsoleLogs = logs.NewLogger(1000)
		ConsoleLogs.SetLogger("console")
		ConsoleLogs.SetLevel(7)
		ConsoleLogs.EnableFuncCallDepth(true)
	}else{
		ConsoleLogs = logs.NewLogger(1000)
		ConsoleLogs.SetLogger("console")
		
		level,err := cfg.Int("logs::consul")
		if err!= nil{
			ConsoleLogs.SetLevel(7)
		}else{
			ConsoleLogs.SetLevel(level)
		}
		ConsoleLogs.EnableFuncCallDepth(true)

		if cfg.String("logs::logdir") != ""{
			FileLogs = logs.NewLogger(1000)
			if len(cfg.Strings("logs::separate")) != 0{
				conf := &configs{Filename:cfg.String("logs::logdir"),Separate:cfg.Strings("logs::separate")}
				result,_ := json.Marshal(conf)
				FileLogs.SetLogger(logs.AdapterMultiFile, string(result))
			}else{
				conf := &configs{Filename:cfg.String("logs::logdir")}
				result,_ := json.Marshal(conf)
				FileLogs.SetLogger(logs.AdapterMultiFile, string(result))
			}
			level,err = cfg.Int("logs::file")
			if err != nil {
				FileLogs.SetLevel(7)
			}else{
				FileLogs.SetLevel(level)
			}
			FileLogs.EnableFuncCallDepth(true)
			//FileLogs.Async()
		}
	}
}

func Fmt2String(e interface{}) string {
	return fmt.Sprint(e)
}

func CloseLogs()  {
	ConsoleLogs.Close()
	FileLogs.Close()
}