package fileOP

import (
	"strings"
	"os"
	"bufio"
	"io"

	"zdy.local/logsSQL/utils"

	"github.com/astaxie/beego"
)

type TestOP struct {
	Pos int
	Sign string
}

func (t *TestOP) Init(){
	t.Pos = 3
	t.Sign = "test"
	utils.ConsoleLogs.Info("init sign is ",t.Sign)
}

func (t *TestOP) ReadFile(filename string) (bool){
	utils.ConsoleLogs.Info("sign is ",t.Sign)
	utils.ConsoleLogs.Debug("test filename is ",filename)
	return true
}

type LogsOP struct {
	Pos string
	Sign string
}

func (l *LogsOP) Init(){
	l.Pos = beego.AppConfig.String("position")
	cfg := beego.AppConfig
	if cfg.String("os")=="win" {
		l.Sign = "\\"
	}else {
		l.Sign = "/"
	}
}

func (l *LogsOP) ReadFile(filename string) (bool){
	utils.FileLogs.Info("START FILE ",filename)
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)		//打开文件
	if err != nil {
		utils.FileLogs.Error("logsOP.ReadFile openfile error",err)
		return false
	}
	buf := bufio.NewReader(file)
	num := -1
	for {
		line, errfile := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		num = strings.Index(line, "UPLOAD SUCC")		//查找日志文件中存在"UPLOAD SUCC"子串的行
		if num != -1 {
			fid,filename,orgname,duns,modifytime := l.splitLins(line,num)
			if fid != "null" {				//文件上传成功数据库操作失败时
				err = SaveDB(fid,filename,orgname,duns,modifytime,0)
				if err != nil {
					utils.FileLogs.Info("SAVE FID ERROR:",fid)
				}else {
					utils.FileLogs.Info("SAVE FID SUCCEED:",fid)
				}
			}else{						//上传weedfs失败时处理
				utils.ConsoleLogs.Debug("upload weedfs")
				utils.FileLogs.Info("UPLOAD FILE ",line)		//记录上传失败文件
				nfid,size,filename := UploadWeed(strings.Split(line[num + 24:], ", ")[0])	//上传weedfs
				if nfid == "null" {
					utils.ConsoleLogs.Debug("upload error")
					continue
				}
				err = SaveDB(nfid,filename,orgname,duns,modifytime,size)
				if err != nil {
					utils.FileLogs.Info("SAVE FID ERROR:",nfid)
				}else {
					utils.FileLogs.Info("SAVE FID SUCCEED:",nfid)
				}
			}
		}
		if errfile != nil {
			if err == io.EOF {
				utils.ConsoleLogs.Error( "read file error.", err.Error())
				utils.FileLogs.Info("READ FINISHED")
				file.Close()
				return true
			}
			break
		}
	}
	file.Close()
	return true
}

func (l *LogsOP) splitLins(line string, num int) (string,string,string,string,string) {
	utils.ConsoleLogs.Debug("line is ",line)
	modifytime := line[8:27]			//获取上传成功时的时刻时间
	utils.ConsoleLogs.Debug("modifytime is ",modifytime)
	filename := strings.Split(line[num + 24:], ", ")//获取文件名的绝对路径
	name := strings.Split(filename[0], l.Sign)
	tmp := name[len(name) - 1]			//获取文件名
	orgname := l.Pos				//获取医院名
	duns := GetDuns(orgname)
	if duns == "none" {
		utils.ConsoleLogs.Error("未能找到医院duns")
		utils.FileLogs.Error("未能找到医院duns",orgname)
	}
	utils.ConsoleLogs.Debug("org name is ",orgname)
	utils.ConsoleLogs.Debug("upload file name is ",tmp)
	fid := strings.Split(filename[1], "=")[1]	//获取上传weedfs成功后返回的fid
	utils.ConsoleLogs.Debug("fid is ",fid)
	return fid,tmp,orgname,duns,modifytime
}