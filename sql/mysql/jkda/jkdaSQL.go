package jkda

import (
	"fmt"
	"time"
	"reflect"
	"zdy.local/utils/nlogs"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/config"
)

type His_lis_pacs struct {
	Lis_id string `orm:"pk"`
	Sfzh string
	Hisid string
	Type string
	Orgname string
	Duns string
	Name string
	Rawname string
	Typename string
	Measuredate time.Time
	Instid string
	Status int
	Created_date string `orm:"auto_now_add;type(datetime)"`
	Created_by string
	Updated_by string
	Updated_date time.Time `orm:"auto_now;type(datetime)"`
	Comx string
	Created_way string
	Urlx string
	Is_indexed int
	Filesize int64
	Minijpg string
	Jpg string
	Sick_id string
}

type Patient struct {
	Patientid string	`orm:"pk"`
	Patientuid string   
	Sickid string
	Sickname string
	Age int
	Sex string
	Sicknamec string
	Birthday time.Time
	Duns string
	Source string
}

type Series struct {
	Studyuid string
	Seriesuid string	`orm:"pk"`
	Duns string
	Source string
	Seriesnumber int
	Imagecount int
	Studydatetime time.Time
	Modality string
	Exambodypart string
}

type Image struct {
	Seriesuid string
	Imagename string	`orm:"pk"`
	Duns string
	Source string
	Urlx string
	Size int64
	Typex string
	Dpath string
	Imagenum int
}

type Study struct {
	Duns string
	Source string
	Studyuid string		`orm:"pk"`
	Patientuid string   
	Bodypart string
	Department string
	Clinical string
	Accessionnumber string
	Reportid string
}

type Report struct{
	Studyid string		
	Patientuid string   
	Reportid string 	`orm:"pk"`
	Duns string
	Source string
	Report string
	Reportlocate string
	Reporturl string
	Result string
	Resultlocate string
	Resulturl string
	Opinion string
	Description1 string
	Conclusion1 string
	Description2 string
	Conclusion2 string
	Info1 string
	Info2 string
	Part string
}

type JKDA struct{
	OrmSQL orm.Ormer
}

func (j *JKDA)Init(){
	//beego.LoadAppConfig("ini","../conf/database.conf")
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	dbUser := cfg.String("jkda::user")
	dbPwd := cfg.String("jkda::password")
	dbHost := cfg.String("jkda::host")
	dbPort := cfg.String("jkda::port")
	dbName := cfg.String("jkda::dbname")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPwd, dbHost, dbPort, dbName)
	err = orm.RegisterDataBase("default", "mysql", dbUrl, 30)
	if err != nil {
		panic(err)
	}
}

func (j *JKDA)RegisterModelx(modelx interface{}){
	orm.RegisterModel(modelx)
}

func (j *JKDA)SetOrm(){
	j.OrmSQL = orm.NewOrm()
}

func (j *JKDA)UpSert(data interface{})(int64,error){
	if num, err := j.OrmSQL.InsertOrUpdate(data); err != nil {
			//panic(err)
			return num,err
		} else {
			return num,err
		}
}

func (j *JKDA)MultiUpSert(pacs interface{}) {
	sind := reflect.Indirect(reflect.ValueOf(pacs))
	num := sind.Len()
	//j.OrmSQL.Begin()
	for i:=0;i<num;i++{
		_,err := j.OrmSQL.InsertOrUpdate(sind.Index(i).Addr().Interface())
		if err != nil {
			//j.OrmSQL.Rollback()
			nlogs.ConsoleLogs.Debug("UPSERT DATA ERROR ",err)
			nlogs.ConsoleLogs.Debug("UPSERT DATA ",fmt.Sprint(sind.Index(i)))
			nlogs.FileLogs.Error("UPSERT DATA ERROR ",err)
			nlogs.FileLogs.Error("UPSERT DATA ",fmt.Sprint(sind.Index(i)))
			//panic(err)
		}
	}
	//j.OrmSQL.Commit()
}

func (j *JKDA)MultiSave(pacs interface{}) (int64,error){
	num := reflect.TypeOf(pacs).Elem().NumField()
	sind := reflect.Indirect(reflect.ValueOf(pacs))
	nlogs.ConsoleLogs.Debug(nlogs.Fmt2String(sind.Len()))
	if (num * sind.Len()) > 65530 {
		num = 65530/num
		var totalnum int64
		totalnum = 0
		for i:=0;i<sind.Len();i+=num{
			if i+num > sind.Len() {
				nlogs.ConsoleLogs.Debug(nlogs.Fmt2String(i),nlogs.Fmt2String(sind.Len()))
				n,err := j.OrmSQL.InsertMulti(sind.Slice(i,sind.Len()).Len(), sind.Slice(i,sind.Len()).Interface())
				if err != nil {
					//panic(err)
					return totalnum,err
				}else {
					totalnum += n
				}
			}else {
				nlogs.ConsoleLogs.Debug(nlogs.Fmt2String(i),nlogs.Fmt2String(i+num))
				n,err := j.OrmSQL.InsertMulti(sind.Slice(i,i+num).Len(), sind.Slice(i,i+num).Interface())
				if err != nil {
					//panic(err)
					return totalnum,err
				}else {
					totalnum += n
				}
			}
		}
		return totalnum,nil
	}else {
		if num, err := j.OrmSQL.InsertMulti(sind.Len(), pacs); err != nil {
			//panic(err)
			return num,err
		} else {
			return num,err
		}
	}
}	