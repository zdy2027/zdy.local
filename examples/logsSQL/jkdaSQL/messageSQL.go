package jkdaSQL

import (
	"fmt"
	"time"

	"zdy.local/logsSQL/utils"

	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"os"
)

type His_lis_pac struct {
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
	Created_date time.Time `orm:"auto_now_add;type(datetime)"`
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

var OrmSQL orm.Ormer

func init(){
	//beego.LoadAppConfig("ini","../conf/database.conf")
	cfg := beego.AppConfig
	dbUser := cfg.String("user")
	dbPwd := cfg.String("password")
	dbHost := cfg.String("host")
	dbPort := cfg.String("port")
	dbName := cfg.String("dbname")
	orm.RegisterModel(new(His_lis_pac))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPwd, dbHost, dbPort, dbName)
	utils.ConsoleLogs.Debug(dbUrl)
	err := orm.RegisterDataBase("default", "mysql", dbUrl, 30)
	if err != nil {
		utils.FileLogs.Error("SQL INIT ERROR",err)
		os.Exit(1)
	}
	OrmSQL = orm.NewOrm()
}