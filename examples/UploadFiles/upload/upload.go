package upload

import (
	"github.com/astaxie/beego/config"
	"zdy.local/utils/nlogs"
	"zdy.local/utils/weedOP"
	"os"
)

type Upload struct {
	WeedServer string
}

func (l *Upload) Init(comx string){
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	weedHost := cfg.String("weedfs::host")
	weedPort := cfg.String("weedfs::port")
	l.WeedServer = weedHost + ":" + weedPort
}

func (l *Upload) ReadFile(filename string) (bool){
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		nlogs.FileLogs.Error("fileOP.upload.open file error",err)
		return false
	}
	fileInfo, err := os.Stat(filename)
	if err != nil {
		nlogs.FileLogs.Error("fileOP.upload.os.Stat filename error",err)
		return false
	}
	nfid,_,err := weedOP.Upload(fileInfo.Name(),"text/plain",file,l.WeedServer)
	nlogs.ConsoleLogs.Debug(filename,nfid)
	nlogs.FileLogs.Info(filename,nfid)
	if err != nil {
		nlogs.FileLogs.Error("UPLOAD FAILED ",fileInfo.Name())
		return false
	}
	nlogs.FileLogs.Debug("UPLOAD SUCCEED ",nfid)

	return true
}

func (l *Upload)Close()  {
	
}