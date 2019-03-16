package database

import (
	"zdy.local/utils/nlogs"
	"zdy.local/cecetl/message"
	"zdy.local/sql/mysql/mysqldriver"
	"fmt"
	"strconv"
	"time"
	"strings"
	"sync"
)

type ConsumerCli struct {
	//Paths map[string][]string
	database map[string]*mysqldriver.Fzmysql
	sd SaveData
	m *sync.Mutex
	pushChans chan bool
	GetMemCache
}

func (cl *ConsumerCli) Init()bool{
	cl.pushChans = make(chan bool, 1)
	cl.sd.Init()
	//cl.InitCache()
	//cl.Paths = make(map[string][]string)
	cl.database = make(map[string]*mysqldriver.Fzmysql)
	for _,db := range cl.sd.Dblist {
		if db.Pathlocate == ""{
			break
		}
		if !strings.HasSuffix(db.Pathlocate,"/"){
			db.Pathlocate = db.Pathlocate + "/"
		}
		cl.database[db.Dbname] = new(mysqldriver.Fzmysql)
		dbUrl := cl.database[db.Dbname].InitJndi(db.Dbname)
		nlogs.ConsoleLogs.Warn("start init",db.Dbname)
		//cl.Paths[db.Dbname] = fileOP.GetPath(db.Pathlocate)
		cl.database[db.Dbname].OrmSQL,_ = cl.sd.SetOrmDriver(db.Dbname,dbUrl)
	}
	nlogs.ConsoleLogs.Warn("init finished")
	//cl.sd.Database.Sql.OrmSQL.Begin()
	return true
}

func (cl *ConsumerCli)Close(){
	cl.sd.OrmSQL.Commit()
}
/*
type Study struct {
	StudyID string		`json:"studyInstanceUID"`	0
	Duns string		`json:"duns"`
	SickInfo Patient	`json:"sickinfo"`		func
	SeriesInfo []Series	`json:"seriesinfo"`		func
	BodyPart string		`json:"bodypart"`		1
	Reportx []Report	`json:"report"`			func
	Department string	`json:"department"`		2
	Clinical string		`json:"clinical"`		3
	Accessionnumber string	`json:"accessionnumber"`	4
}
 */

type study_sql_info struct {
	Studyinstanceuid string
	Bodypart string
	Department string
	Clinical string
	Accessionnumber string
	Reportid string
	Studytime string
	//Patientsource string
	Devicename string
}

func (cl *ConsumerCli) GetStudyUID(dbname,patientUID string,res *message.Private){
	var common message.Study
	//nlogs.ConsoleLogs.Alert("start get file",patientUID)
	cl.GetFile(dbname,patientUID,&common,res.SickInfo.SickID)
	var maps []study_sql_info
	//cl.m.Lock()
	sql := fmt.Sprintf(cl.sd.Dblist[dbname].Studysql,patientUID)
	//cl.m.Unlock()
	nlogs.ConsoleLogs.Debug(sql)
	_, err := cl.database[dbname].OrmSQL.Raw(sql).QueryRows(&maps)
	if err!=nil {
		nlogs.FileLogs.Error(err.Error())
		return
	}
	local, _ := time.LoadLocation("UTC")
	for _,term := range maps{
		var result message.Study
		if term.Studyinstanceuid == ""{
			continue
		}
		result.StudyID = term.Studyinstanceuid
		result.BodyPart = term.Bodypart
		result.Department = term.Department
		result.Clinical = term.Clinical
		result.Accessionnumber = term.Accessionnumber
		result.Reportx.ReportID = term.Reportid
		result.Studytime,err = time.ParseInLocation("2006-01-02 15:04:05",term.Studytime,local)
		if err != nil {
			nlogs.FileLogs.Warn("time parse error")
			result.Studytime,err = time.Parse("2006-01-02 15:04:05 +0000 UTC","0000-00-00 00:00:00 +0000 UTC")
		}
		result.Devicename = term.Devicename
		nlogs.ConsoleLogs.Debug(result.StudyID)
		cl.GetReport(dbname,result.StudyID,&result.Reportx)						//查询report数据信息存储至report变量中
		cl.GetSeriesUID(dbname,result.StudyID,&result,res.SickInfo.SickID)	//查询series数据信息存储至result变量中
		res.Studys = append(res.Studys,result)
	}
}

func (cl *ConsumerCli) LoadData(dbname string,result message.Private){
	cl.sd.SaveData(result,cl.sd.Dblist[dbname].Dbname)
}

/*
type Series struct {
	SeriesID string		`json:"seriesID"`		0
	Imgs []Image		`json:"imgs,omitempty"`		func
	Count int		`json:"count"`			1
	SeriesNum int		`json:"seriesnum"`		5
	DateTimex time.Time	`json:"datetimex,omitempty"`	2
	BodyPart string		`json:"bodypart"`		3
	Modality string		`json:"modality"`		4
}
 */

type series_sql_info struct {
	Seriesid string
	Count string
	Datetimex string
	Bodypart string
	Comx string
	Modality string
	Seriesnum string
}

func (cl *ConsumerCli)GetSeriesUID(dbname,StudyInstanceUID string,study *message.Study,idcard string){
	var maps []series_sql_info
	sql := fmt.Sprintf(cl.sd.Dblist[dbname].Seriessql,StudyInstanceUID)
	nlogs.ConsoleLogs.Debug(sql)
	_, err := cl.database[dbname].OrmSQL.Raw(sql).QueryRows(&maps)
	if err!=nil {
		nlogs.FileLogs.Error(err.Error())
		return
	}
	local, _ := time.LoadLocation("UTC")
	for _,term := range maps{
		var result message.Series
		var err error
		if term.Seriesid == ""{
			continue
		}
		result.SeriesID = term.Seriesid
		result.Count,err = strconv.Atoi(term.Count)
		if err != nil {
			result.Count = 0
		}
		result.DateTimex,err = time.ParseInLocation("2006-01-02 15:04:05",term.Datetimex,local)
		if err != nil {
			nlogs.FileLogs.Warn("time parse error")
			result.DateTimex,err = time.Parse("2006-01-02 15:04:05 +0000 UTC","0000-00-00 00:00:00 +0000 UTC")
		}
		if term.Bodypart!=""{
			result.BodyPart = term.Bodypart
		}else {
			result.BodyPart = term.Comx
		}
		result.Modality = term.Modality
		result.SeriesNum,err = strconv.Atoi(term.Seriesnum)
		if err != nil {
			result.SeriesNum = 0
		}
		cl.GetImg(dbname,result.SeriesID,&result,idcard)
		study.SeriesInfo = append(study.SeriesInfo,result)
	}
}

/*
type Image struct {
	Imgname string		`json:"imgname"`		0
	Urlx string		`json:"url,omitempty"`		func
	Size int64		`json:"size,omitempty"`		func
	Typex string		`json:"type,omitempty"`		2
	Statusx bool 		`json:"status,omitempty"`	func
	Dpath string		`json:"dpath"`			1(path) + 0 + 替换
	Imagenum int		`json:"imgnum"`
}
 */

type file_sql_info struct {
	Storepath string
	patientUID string
	Modality string
	Bodypart string
	Department string
}

func (cl *ConsumerCli)GetFile(dbname,patientUID string,result *message.Study,idcard string){
	//nlogs.FileLogs.Debug("upload instance",patientUID)
	var files []file_sql_info
	sql := fmt.Sprintf(cl.sd.Dblist[dbname].Filesql,patientUID)
	nlogs.ConsoleLogs.Debug(sql)
	_, err := cl.database[dbname].OrmSQL.Raw(sql).QueryRows(&files)
	if err!=nil {
		nlogs.FileLogs.Error(err.Error())
		return
	}
	for _,term := range files{
		if term.Storepath == ""{
			continue
		}
		result.StudyID = patientUID
		result.BodyPart = term.Bodypart
		result.Department = term.Department
		var tmp_series message.Series
		tmp_series.SeriesID = "series"
		var tmp message.Image
		//tmp.Spath = fmt.Sprintf("%s",*(data[1].(*interface{}))) + fmt.Sprintf("%s",*(data[0].(*interface{})))
		tmp.Typex = term.Modality
		tmp.Dpath = strings.Replace(term.Storepath,"\\","/",-1)
		//tmp.Imagenum,_ = strconv.Atoi(data.Imgnum)
		//tmp.Imgname = data.Imgname
		p := strings.Split(tmp.Dpath,"/")
		n := len(p)
		var images []mysqldriver.Image
		sql := "select dpath,size,urlx from image where source = '"+dbname+"' and prepath like '"+p[n-2]+"/"+p[n-1]+ "/%'"
		nlogs.ConsoleLogs.Debug(sql)
		_,err := cl.sd.OrmSQL.Raw(sql).QueryRows(&images)
		//image,err := cl.GetMessage(strings.Split(tmp.Dpath,"/")[len(strings.Split(tmp.Dpath,"/"))-1] + "/" + data.Imgname)
		if err!=nil{
			nlogs.FileLogs.Error("not find file",tmp.Dpath)
			tmp.Statusx = false
			tmp_series.Imgs = append(tmp_series.Imgs,tmp)
			//tmp.Dpath = data.Path + data.Imgname
		}else {
			for _,img := range images{
				var tmp message.Image
				tmp.Typex = term.Modality
				nlogs.ConsoleLogs.Debug(img.Dpath)
				tmp.Dpath = img.Dpath
				tmp.Size = img.Size
				tmp.Urlx = img.Urlx
				tmp.Statusx = true
				tmp.Imgname = img.Imagename
				tmp_series.Imgs = append(tmp_series.Imgs,tmp)
			}
		}
		result.SeriesInfo = append(result.SeriesInfo,tmp_series)
	}
}

type image_sql_info struct {
	Imgname string
 	Typex string
	Path string
	Imgnum string
}

func (cl *ConsumerCli)GetImg(dbname,SeriesInstanceUID string,result *message.Series,idcard string){
	var maps []image_sql_info
	sql := fmt.Sprintf(cl.sd.Dblist[dbname].Imagesql,SeriesInstanceUID)
	nlogs.ConsoleLogs.Debug(sql)
	_, err := cl.database[dbname].OrmSQL.Raw(sql).QueryRows(&maps)
	if err!=nil {
		nlogs.FileLogs.Error(err.Error())
		return
	}
	for _,term := range maps{
		if term.Imgname == ""{
			continue
		}
		cl.WriteImg(result,term,idcard,dbname)
	}
}

func (cl *ConsumerCli)WriteImg(result *message.Series, data image_sql_info,idcard,dbname string) string{
	var tmp message.Image
	//tmp.Spath = fmt.Sprintf("%s",*(data[1].(*interface{}))) + fmt.Sprintf("%s",*(data[0].(*interface{})))
	tmp.Typex = data.Typex
	tmp.Dpath = strings.Replace(data.Path,"\\","/",-1)
	tmp.Imagenum,_ = strconv.Atoi(data.Imgnum)
	tmp.Imgname = data.Imgname
	var image mysqldriver.Image
	p := strings.Split(tmp.Dpath,"/")
	n := len(p)
	sql := "select dpath,size,urlx from image where source = '"+dbname+"' and prepath = '"+p[n-2]+"/"+p[n-1]+"/"+data.Imgname+
		"' limit 1"
	nlogs.ConsoleLogs.Debug(sql)
	err := cl.sd.OrmSQL.Raw(sql).QueryRow(&image)
	//image,err := cl.GetMessage(strings.Split(tmp.Dpath,"/")[len(strings.Split(tmp.Dpath,"/"))-1] + "/" + data.Imgname)
	if err!=nil{
		nlogs.FileLogs.Error("not find image",p[n-2]+"/"+p[n-1]+"/"+data.Imgname)
		tmp.Statusx = false
		tmp.Dpath = data.Path + data.Imgname
	}else {
		nlogs.ConsoleLogs.Debug(image.Dpath)
		tmp.Dpath = image.Dpath
		tmp.Size = image.Size
		tmp.Urlx = image.Urlx
		tmp.Statusx = true
	}
	//paths := fileOP.GetPath(cl.Path)
	/*
	num := 0
	for _,p := range cl.Paths[dbname]{
		fileinfo,err:= os.Stat(p)
		if err != nil{
			continue
		}
		num_index := strings.Index(tmp.Dpath,fileinfo.Name())
		if num_index != -1 {
			tmp.Dpath = strings.TrimSuffix(p,fileinfo.Name())+tmp.Dpath[num_index:]
			if strings.HasSuffix(tmp.Dpath,"/"){
				tmp.Dpath = tmp.Dpath + data.Imgname
			}else {
				tmp.Dpath = tmp.Dpath + "/" + data.Imgname
			}
			tmp.Imagenum,_ = strconv.Atoi(data.Imgnum)
			//cl.pushChans <- true
			fileinfo,err:=os.Stat(tmp.Dpath)
			if err!=nil{
				continue
			}else {
				tmp.Imgname = fileinfo.Name()
				nlogs.ConsoleLogs.Info("SAVE FILE ",tmp.Dpath)
				fid,size := cl.sd.SaveFile(tmp.Dpath)
				tmp.Size = size
				tmp.Urlx = fid
				tmp.Statusx = true
				num++
				//<- cl.pushChans
				break
			}
			//<- cl.pushChans
		}
		//num++
	}
	if num == len(cl.Paths[dbname]){
		tmp.Statusx = false
		tmp.Imgname = data.Imgname
		tmp.Dpath = data.Path + data.Imgname
	}*/
	result.Imgs = append(result.Imgs,tmp)
	return tmp.Imgname
}

/*
type Report struct {
	ReportID string 	`json:"reportid"`
	Report string		`json:"report"`
	ReportLocate string	`json:"reportlocate"`
	ReportUrl string	`json:"reporturl"`
	Result string		`json:"result"`
	ResultLocate string	`json:"resultlocate"`
	ResultUrl string	`json:"resulturl"`
	Opinion string		`json:"opinion"`
	Description1 string	`json:"description1"`
	Conclusion1 string	`json:"conclusion1"`
	Description2 string	`json:"description2"`
	Conclusion2 string	`json:"conclusion2"`
	Info1 string		`json:"info1"`
	Info2 string		`json:"info2"`
	BodyPart string		`json:"bodypart"`
}

 */

type report_sql_info struct {
	Report string
	Reportpath string
	Reportname string
	Result string
	Resultpath string
	Resultname string
	Opinion string
	Description1 string
	Conclusion1 string
	Info1 string
	Description2 string
	Conclusion2 string
	Info2 string
	Bodypart string
}

func (cl *ConsumerCli)GetReport(dbname,StudyInstanceUID string,result *message.Report){
	var maps []report_sql_info
	sql := fmt.Sprintf(cl.sd.Dblist[dbname].Reportsql,StudyInstanceUID)
	_, err := cl.database[dbname].OrmSQL.Raw(sql).QueryRows(&maps)
	if err!=nil {
		nlogs.FileLogs.Error(err.Error())
		return
	}
	for _,term := range maps{
		result.Report = term.Report
		var image mysqldriver.Image
		Dpath := strings.Replace(term.Reportpath,"\\","/",-1)
		result.ReportLocate = Dpath + term.Reportname
		if term.Reportname != ""{
			p := strings.Split(Dpath,"/")
			n := len(p)
			sql := "select dpath,urlx from image where source = '"+dbname+"' and and prepath = '"+
				p[n-2]+"/"+p[n-1]+"/"+term.Reportname +"' limit 1"
			err := cl.sd.OrmSQL.Raw(sql).QueryRow(&image)
			//image,err := cl.GetMessage(strings.Split(term.Reportpath,"/")[len(strings.Split(term.Reportpath,"/"))-1] + "/" + term.Reportname)
			if err!=nil{
				nlogs.FileLogs.Error("sql error",term.Reportname)
			}else {
				result.ReportLocate = image.Dpath
				result.ReportUrl = image.Urlx
			}
		}
		/*for _,p := range cl.Paths[dbname]{
			num_index := strings.Index(Dpath,p)
			if num_index > 0 {
				Dpath = cl.sd.Database.Dblist[dbname].Pathlocate+Dpath[num_index-1:]
				break
			}
		}
		Dpath = Dpath + term.Reportname
		//cl.pushChans <- true
		_,err=os.Stat(Dpath)
		if err==nil{
			result.ReportUrl,_ = cl.sd.SaveFile(Dpath)
		}*/
		//<- cl.pushChans
		result.Result = term.Result
		Dpath = strings.Replace(term.Resultpath,"\\","/",-1)
		result.ResultLocate = Dpath + term.Resultname
		if term.Resultname != ""{
			p := strings.Split(Dpath,"/")
			n := len(p)
			sql := "select dpath,urlx from image where source = '"+dbname+"' and and prepath = '"+
				p[n-2]+"/"+p[n-1]+"/"+term.Resultname + "' limit 1"
			//image,err := cl.GetMessage(strings.Split(term.Resultpath,"/")[len(strings.Split(term.Resultpath,"/"))-1] + "/" + term.Resultname)
			err := cl.sd.OrmSQL.Raw(sql).QueryRow(&image)
			if err!=nil{
				nlogs.FileLogs.Error("sql error",term.Resultname)
			}else {
				result.ResultLocate = image.Dpath
				result.ResultUrl = image.Urlx
			}
		}
		/*Dpath = strings.Replace(term.Resultpath,"\\","/",-1)
		for _,p := range cl.Paths[dbname]{
			num_index := strings.Index(Dpath,p)
			if num_index > 0 {
				Dpath = cl.sd.Database.Dblist[dbname].Pathlocate+Dpath[num_index-1:]
				break
			}
		}
		Dpath = Dpath + term.Resultname
		//cl.pushChans <- true
		_,err=os.Stat(Dpath)
		if err==nil{
			result.ResultUrl,_ = cl.sd.SaveFile(Dpath)
		}*/
		//<- cl.pushChans
		//result.ResultUrl

		result.Opinion = term.Opinion

		result.Description1 = term.Description1
		result.Conclusion1 = term.Conclusion1

		result.Description2 = term.Description2
		result.Conclusion2 = term.Conclusion2

		result.Info1 = term.Info1
		result.Info2 = term.Info2

		result.BodyPart = term.Bodypart
	}
}