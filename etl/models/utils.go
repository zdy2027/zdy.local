package models

import (
	"github.com/astaxie/beego/config"
	"zdy.local/cecetl/nosql"
	"zdy.local/cecetl/message"
	"zdy.local/utils/nlogs"
	"zdy.local/utils/weedOP"
	"zdy.local/fileOP"
	"io/ioutil"
	"encoding/json"
	"os"
	"strings"
	"strconv"
	"sync"
	"zdy.local/cecetl/database"
	"fmt"
	"time"
)

func GetPatients(idcard string)message.PrivateHits{
	var esclient nosql.ES
	esclient.Init()
	result := esclient.GetPatients(idcard)
	return result
}

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

func (l *Upload) ReadFile(filename string,wg *sync.WaitGroup,goroutine chan int) (bool){
	defer wg.Done()
	//nfid,_,err := weedOP.Upload(fileInfo.Name(),"text/plain",file,l.WeedServer)
	nfid,_,err := weedOP.UploadFile(l.WeedServer,filename)
	nlogs.ConsoleLogs.Debug(filename,nfid)
	nlogs.FileLogs.Info(filename,nfid)
	<- goroutine
	if err != nil {
		nlogs.FileLogs.Error("UPLOAD FAILED ",filename)
		return false
	}
	nlogs.FileLogs.Debug("UPLOAD SUCCEED ",nfid)

	return true
}

func (l *Upload)Close()  {

}

func ReadJson(filename,destdir string){
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		nlogs.ConsoleLogs.Error("ReadFile: ", err.Error())
	}
	var result message.PrivateHits
	if err := json.Unmarshal(bytes, &result); err != nil {
		nlogs.ConsoleLogs.Error("Unmarshal: ", err.Error())
	}
	for _,info := range result.Info{
		for _,study := range info.Studys{
			for _,series := range study.SeriesInfo{
				for _,img := range series.Imgs{
					nlogs.ConsoleLogs.Info(img.Dpath)
					fileinfo,err := os.Stat(img.Dpath)
					if err != nil{
						nlogs.ConsoleLogs.Error("info error",img.Dpath)
						continue
					}
					nlogs.ConsoleLogs.Info(fileinfo.Name())
					fileOP.CopyFile(img.Dpath,destdir+fileinfo.Name())
				}
			}
		}
	}
}

func SaveJson(filename,destdir string){
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		nlogs.ConsoleLogs.Error("ReadFile: ", err.Error())
	}
	var result message.PrivateHits
	if err := json.Unmarshal(bytes, &result); err != nil {
		nlogs.ConsoleLogs.Error("Unmarshal: ", err.Error())
	}
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	weedHost := cfg.String("weedfs::host")
	weedPort := cfg.String("weedfs::port")
	WeedServer := weedHost + ":" + weedPort
	var Es nosql.ES
	Es.Init()
	for i,info := range result.Info{
		for j,study := range info.Studys{
			for k,series := range study.SeriesInfo{
				for l,img := range series.Imgs{
					nlogs.ConsoleLogs.Info(img.Dpath)
					pathlist := strings.Split(img.Dpath,"/")
					path := destdir + pathlist[len(pathlist)-1]
					fileinfo,err := os.Stat(path)
					if err != nil{
						nlogs.ConsoleLogs.Error("info error",img.Dpath)
						continue
					}
					nlogs.ConsoleLogs.Info(fileinfo.Name())
					nfid,_,err := weedOP.UploadFile(WeedServer,path)
					nlogs.ConsoleLogs.Debug(filename,nfid)
					//nlogs.FileLogs.Info(filename,nfid)
					if err != nil {
						nlogs.FileLogs.Error("UPLOAD FAILED ",path)
						continue
					}
					result.Info[i].Studys[j].SeriesInfo[k].Imgs[l].Urlx = nfid
					//img.Urlx = nfid
					//nlogs.FileLogs.Debug("UPLOAD SUCCEED ",nfid)
				}
			}
		}
		text,_ := json.Marshal(info)
		//Es.IndexData(string(text),result.SickInfo.SickID,"pacs")
		nlogs.ConsoleLogs.Info(string(text))
		Es.IndexData(strconv.Itoa(i),"dcm",string(text))
	}
}

func AddData(filename,destdir string){
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		nlogs.ConsoleLogs.Error("ReadFile: ", err.Error())
	}
	var result message.PrivateHits
	if err := json.Unmarshal(bytes, &result); err != nil {
		nlogs.ConsoleLogs.Error("Unmarshal: ", err.Error())
	}
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	weedHost := cfg.String("weedfs::host")
	weedPort := cfg.String("weedfs::port")
	WeedServer := weedHost + ":" + weedPort
	var Es nosql.ES
	Es.Init()
	dir, err := ioutil.ReadDir(destdir)
	if err != nil {
		nlogs.ConsoleLogs.Error("open dir error",err)
	}
	for i,info := range result.Info{
		for _, fi := range dir{
			nlogs.ConsoleLogs.Debug("upload dir",fi.Name())
			var seriesdata message.Series
			seriesdata.SeriesID = fi.Name()
			seriesdata.BodyPart = "肺部"
			seriesdata.DateTimex = time.Now()
			seriesdata.Modality = "CT"
			files,_ := ioutil.ReadDir(destdir+fi.Name())
			for j,f := range files {
				//nlogs.ConsoleLogs.Debug("uploa")
				var ctimgs message.Image
				nfid,size,err := weedOP.UploadFile(WeedServer,destdir+fi.Name()+"/"+f.Name())
				nlogs.ConsoleLogs.Debug(f.Name(),nfid)
				//nlogs.FileLogs.Info(filename,nfid)
				if err != nil {
					nlogs.FileLogs.Error("UPLOAD FAILED ",f)
					continue
				}
				ctimgs.Dpath = fmt.Sprint(f)
				ctimgs.Imagenum = j
				ctimgs.Urlx = nfid
				ctimgs.Imgname = f.Name()
				ctimgs.Size = size
				ctimgs.Statusx = true
				seriesdata.Imgs = append(seriesdata.Imgs,ctimgs)
			}
			result.Info[0].Studys[0].SeriesInfo = append(result.Info[0].Studys[0].SeriesInfo,seriesdata)
		}
		text,_ := json.Marshal(info)
		//Es.IndexData(string(text),result.SickInfo.SickID,"pacs")
		nlogs.ConsoleLogs.Info(string(text))
		Es.IndexData(strconv.Itoa(i),"dcm",string(text))
	}
}

func CacheImg(duns string){
	var m2r database.GetMemCache
	m2r.InitCache()
	m2r.GetImages(duns)
}