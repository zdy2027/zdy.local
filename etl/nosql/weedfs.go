package nosql

import (
	"strings"
	"os"
	"zdy.local/utils/nlogs"
	"zdy.local/utils/zipFile"
	"zdy.local/utils/weedOP"
	"github.com/astaxie/beego"
)

var (
	weedip = beego.AppConfig.String("weedfs::host")+":"+beego.AppConfig.String("weedfs::port")
)

func UploadWeed(filename string) (string ,int64){
	nlogs.FileLogs.Info("UPLOAD WEEDFS START ",filename)
	if strings.HasSuffix(filename,".ZIP") || strings.HasSuffix(filename,".PK") {
		err := zipFile.DeCompress(filename,filename[:strings.LastIndex(filename,"/")+1])
		if err != nil {
			nlogs.FileLogs.Error("fileOP.UploadWeed.DeCompress method error",err)
			return "",0
		}
		return upload(filename[:strings.LastIndex(filename,".")],true)
	}
	return upload(filename,false)
}

func upload(filename string,sign bool)(string ,int64){
	/*file, err := os.Open(filename)
	if err != nil {
		nlogs.FileLogs.Error("fileOP.upload.open file error",err)
		return "null",0,"null"
	}*/
	//defer file.Close()

	nfid,size,err := weedOP.UploadFile(weedip,filename)
	//nfid,size,err := weedOP.Upload(fileInfo.Name(),"text/plain",file,weedip)
	nlogs.ConsoleLogs.Debug("UPLOAD SUCCEED ",nfid)
	if err != nil {
		nlogs.FileLogs.Info("UPLOAD FAILED ",err.Error(),filename)
		return "",size
	}
	if sign {
		err = os.Remove(filename)
		if err != nil {
			nlogs.FileLogs.Error("fileOP.upload.os.remove method error",err)
		}
	}
	nlogs.FileLogs.Info("UPLOAD SUCCEED ",nfid)
	return nfid,size
}