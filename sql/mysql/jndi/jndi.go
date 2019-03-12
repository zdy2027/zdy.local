package jndi

import (
	"fmt"
	"zdy.local/utils/nlogs"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/config"
)

type EtlJndi struct {
	Jndi string		`orm:"pk"`
	Type string
	Status int
	Val string
	Comx string
	Duns string
	Dunsname string
}

type Prothread struct{
	Duns string
	Orgname string
	Dbname string		`orm:"pk"`
	Patientsql string
	Studysql string
	Seriessql string
	Imagesql string
	Reportsql string
	Typex string
}

type JNDI struct{
	ormSQL orm.Ormer
	ThreadSql Prothread
}

func (j *JNDI)Init(dbName string){
	//cfg, err := config.NewConfig("ini","database.conf")
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		nlogs.ConsoleLogs.Error(err.Error())
		panic(err)
	}
	dbUser := cfg.String("jndi::user")
	dbPwd := cfg.String("jndi::password")
	dbHost := cfg.String("jndi::host")
	dbPort := cfg.String("jndi::port")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPwd, dbHost, dbPort, "cecetl")
	nlogs.ConsoleLogs.Debug(dbUrl)
	err = orm.RegisterDataBase("cecetl", "mysql", dbUrl, 30)
	//orm.RegisterModel(new(Prothread))
	if err != nil {
		panic(err)
	}
	j.ormSQL = orm.NewOrm()
	j.GetSQL(dbName)
}

func (j *JNDI)GetSQL(dbName string){
	j.ormSQL.Using("cecetl")
	j.ThreadSql = Prothread{Dbname:dbName}
	err := j.ormSQL.Read(&j.ThreadSql)
	if err != nil{
		panic(err)
	}
}

func (j *JNDI)SetDB(dbname string) orm.Ormer {
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		nlogs.ConsoleLogs.Error(err.Error())
		panic(err)
	}
	dbUser := cfg.String("jndi::user")
	dbPwd := cfg.String("jndi::password")
	dbHost := cfg.String("jndi::host")
	dbPort := cfg.String("jndi::port")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPwd, dbHost, dbPort,dbname)
	nlogs.ConsoleLogs.Debug(dbUrl)
	err = orm.RegisterDataBase("jndi", "mysql", dbUrl, 30)
	if err != nil {
		panic(err)
	}
	return orm.NewOrm()
}