package fileOP

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"zdy.local/logsSQL/utils"
	"zdy.local/logsSQL/jkdaSQL"
	"zdy.local/logsSQL/weedOP"
	"zdy.local/logsSQL/zipFile"

	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

var (
	kDict = "dict.xlsx"
	dictFile, _ = xlsx.OpenFile(kDict)
	dict map[string]string
	o = jkdaSQL.OrmSQL						//获取数据库omr变量进行数据库操作
	weedip = beego.AppConfig.String("weedfs")+":"+beego.AppConfig.String("weedport")
	comx = "ETL_LOCAL_UPLOAD_ID="
)

func init(){
	dict = make(map[string]string)
	sheet := dictFile.Sheets[0]
	for i:=1;i<len(sheet.Rows);i++ {
		dict[sheet.Rows[i].Cells[1].Value] = sheet.Rows[i].Cells[2].Value
	}
}

func Init(inputs []string,comx string,obj FileOP){
	//输入文件可以是路径列表，此处遍历所有输入路径列表
	for _,input := range inputs {
		if isDir(input) {				//判断该路径是否是文件夹
			utils.ConsoleLogs.Debug("is dir")
			walkDir(input,obj)				//循环遍历文件夹下所有子文件
		}else {
			utils.ConsoleLogs.Debug(input)
			//readFile(input)				//开始读取文件内容
			obj.ReadFile(input)
		}
	}

}

func GetDuns(name string)(string){
	for tmp,value := range dict {
		if strings.Contains(tmp,name){
			return value[2:len(value)-1]
		}
	}
	return "none"
}

//循环遍历子文件夹
func walkDir(dir string,obj FileOP){
	filepath.Walk(dir,func(path string, fi os.FileInfo, err error)error{
		if err != nil {
			return err
		}
		utils.ConsoleLogs.Debug(fi.Name())
		if fi.IsDir() {
			return nil				//如果是子文件夹继续遍历
		}else {
			//readFile(path)				//如果是文件开始读取文件
			obj.ReadFile(path)
		}
		return nil
	})
}

func SaveDB(fid,filename,orgname,duns,modifytime string, size int64) error{
	ntime,_ := time.Parse(modifytime,modifytime)
	pacs := jkdaSQL.His_lis_pac{Lis_id:fid,Hisid:filename,Urlx:fid,Type:"2",Orgname:orgname,
		Duns:duns,Rawname:filename,Status:2,Created_by:"admin",Created_date:ntime,
		Comx:comx,Filesize:size}
	_, err := o.Insert(&pacs)		//重新插入数据库数据
	if err == nil {
		utils.ConsoleLogs.Debug("upload succeed ",fid)
		utils.FileLogs.Info("SAVE DB SUCCEED ",fid)
	}else {
		utils.FileLogs.Error("SAVE DB ERROR ",fid)
	}
	return  err
}

func upload(filename string,sign bool)(string ,int64, string){
	file, err := os.Open(filename)
	if err != nil {
		utils.FileLogs.Error("fileOP.upload.open file error",err)
		return "null",0,"null"
	}
	fileInfo, err := os.Stat(filename)
	if err != nil {
		utils.FileLogs.Error("fileOP.upload.os.Stat filename error",err)
		return "null",0,"null"
	}
	nfid,size,err := weedOP.Upload(fileInfo.Name(),"text/plain",file,weedip)
	if err != nil {
		utils.FileLogs.Info("UPLOAD FAILED ",fileInfo.Name())
		return "null",size,fileInfo.Name()
	}
	if sign {
		err = os.Remove(filename)
		if err != nil {
			utils.FileLogs.Error("fileOP.upload.os.remove method error",err)
		}
	}
	utils.FileLogs.Info("UPLOAD SUCCEED ",nfid)
	return nfid,size,fileInfo.Name()
}

func UploadWeed(filename string) (string ,int64 ,string){
	utils.FileLogs.Info("UPLOAD WEEDFS START ",filename)
	if strings.HasSuffix(filename,".ZIP") || strings.HasSuffix(filename,".PK") {
		err := zipFile.DeCompress(filename,filename[:strings.LastIndex(filename,"/")+1])
		if err != nil {
			utils.FileLogs.Error("fileOP.UploadWeed.DeCompress method error",err)
			return "null",0,"null"
		}
		return upload(filename[:strings.LastIndex(filename,".")],true)
	}
	return upload(filename,false)
}

//获取路径是否是文件夹判断
func isDir(filename string) bool {
	fileInfo, _ := os.Stat(filename)
	utils.ConsoleLogs.Debug("Is Directory: ", fileInfo.IsDir())
	return fileInfo.IsDir()
}
