package files

import (
	"zdy.local/utils/nlogs"
	"zdy.local/cecetl/nosql"
	"zdy.local/sql/mysql/mysqldriver"
	"sync"
	"strings"
)

type Upload struct {
	Duns string
	Num int
	DbName string
	sd mysqldriver.Fzmysql
	m *sync.Mutex
}

func (l *Upload) Init(comx string){
	l.DbName = comx
	l.Num = 0
	l.sd.InitJkda()
	nlogs.ConsoleLogs.Info(l.DbName)
	l.Duns = l.sd.Dblist[comx].Duns
	//l.sd.Sql.OrmSQL.Begin()
	l.m = new(sync.Mutex)
}

func (l *Upload) ReadFile(filename string,wg *sync.WaitGroup,goroutine chan int) (bool){
	defer wg.Done()
	defer func() {
		if err := recover(); err != nil {
			nlogs.ConsoleLogs.Info("panic error")
		}
	}()
	nlogs.ConsoleLogs.Debug("start upload pacs",filename)
	var image mysqldriver.Image
	//var dcm message.DCM
	if filename[len(filename)-3:] == "DCM" || filename[len(filename)-3:]=="dcm" {
		//dcm.Pacs = cecdcm.DCMExtract(filename)
	}
	p := strings.Split(filename,"/")
	n := len(p)
	image = mysqldriver.Image{
		Dpath:filename,
		Duns:l.Duns,
		Seriesuid:"unknown",
		Source:l.DbName,
		Imagename:p[n-1],
	}
	err := l.sd.OrmSQL.Read(&image,"Dpath","Duns","Seriesuid")
	if err != nil{
		//nlogs.ConsoleLogs.Info("read sql error",err.Error())
		nfid,size := nosql.UploadWeed(filename)
		nlogs.ConsoleLogs.Info(filename,nfid)
		image = mysqldriver.Image{
			Dpath:filename,
			Duns:l.Duns,
			Imagename:p[n-1],
			Seriesuid:"unknown",
			Size:size,
			Urlx:nfid,
			Source:l.DbName,
			Prepath:p[n-3]+"/"+p[n-2]+"/"+p[n-1],
		}
		_,err := l.sd.OrmSQL.Insert(&image)
		if err != nil {
			nlogs.FileLogs.Error("Save sql error ",filename,nfid,err.Error())
			nlogs.ConsoleLogs.Info("Save sql error",err.Error())
			//panic(err)
		}
		l.m.Lock()
		l.Num++
		l.m.Unlock()
	}else {
		//image.Imagename = p[n-1]
		image.Prepath = p[n-3]+"/"+p[n-2]+"/"+p[n-1]
		_,err := l.sd.OrmSQL.Update(&image)
		if err != nil {
			nlogs.ConsoleLogs.Error("update error",err.Error())
			nlogs.FileLogs.Error("update error",err.Error())
		}
		nlogs.ConsoleLogs.Info(image.Prepath,image.Urlx)
	}
	<- goroutine
	return true
}

func (l *Upload)Close()  {
	nlogs.ConsoleLogs.Info("stopped upload")
	l.sd.Dblist[l.DbName].Filenumber = l.Num
	l.sd.UpSert(l.sd.Dblist[l.DbName])
	//l.sd.Sql.OrmSQL.Commit()
}
