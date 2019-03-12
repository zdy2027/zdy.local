package fileOP

import "github.com/astaxie/beego"

type PacsOBJ struct {
	Pos int
	Sign string
}

func (l *PacsOBJ) Init(){
	l.Pos,_ = beego.AppConfig.Int("position")
	cfg := beego.AppConfig
	if cfg.String("os")=="win" {
		l.Sign = "\\"
	}else {
		l.Sign = "/"
	}
}
