package mysqldriver

import (
	"github.com/astaxie/beego/config"
	"zdy.local/utils/nlogs"
	"fmt"
)

type Fzmysql struct {
	//ThreadSql Prothread
	MysqlDriver
	Dblist map[string]*Prothread
	Sign bool
}

func (fz *Fzmysql)InitJkda()  {
	fz.Sign = false
	fz.Dblist = make(map[string]*Prothread)
	fz.RegisterModelx(new(Patient))
	fz.RegisterModelx(new(Series))
	fz.RegisterModelx(new(Image))
	fz.RegisterModelx(new(Study))
	fz.RegisterModelx(new(Report))
	fz.RegisterModelx(new(Prothread))

	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		nlogs.ConsoleLogs.Error(err.Error())
		panic(err)
	}
	dbUser := cfg.String("jkda::user")
	dbPwd := cfg.String("jkda::password")
	dbHost := cfg.String("jkda::host")
	dbPort := cfg.String("jkda::port")
	dbName := cfg.String("jkda::dbname")
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPwd, dbHost, dbPort, dbName)
	//fz.Sql.RegisterDatabase("default",dbUrl)
	fz.OrmSQL,_ = fz.SetOrmDriver("default", dbUrl)
	fz.InitDBList()
	//fz.GetSQL(dbName)
}

func (fz *Fzmysql)InitJndi(db string) string{
	//fz.Sql.RegisterModelx(new(EtlJndi))
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		nlogs.ConsoleLogs.Error(err.Error())
		panic(err)
	}
	dbUser := cfg.String("jndi::user")
	dbPwd := cfg.String("jndi::password")
	dbHost := cfg.String("jndi::host")
	dbPort := cfg.String("jndi::port")
	//dbName := cfg.String("jndi::dbname")
	//dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPwd, dbHost, dbPort, dbName)
	//fz.Sql.RegisterDatabase("cecetl",dbUrl)

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPwd, dbHost, dbPort, db)
	//fz.Sql.RegisterDatabase("jndi",dbUrl)
	return dbUrl
	//fz.Sql.SetOrmDriver("jndi", dbUrl)
}
/*
func (fz *Fzmysql)GetSQL(dbName string){
	//fz.Sql.OrmSQL.Using("cecetl")
	//fz.Sql.OrmSQL, err := orm.GetDB()
	fz.ThreadSql = Prothread{Dbname:dbName}
	err := fz.Sql.OrmSQL.Read(&fz.ThreadSql)
	if err != nil{
		panic(err)
	}
}*/

func (fz *Fzmysql)PanicSql(){
	if err := recover(); err != nil {
		fmt.Println("panic error, start commit")
		fz.Sign = true
	}
}

func (fz *Fzmysql)CommitSql(){
	fz.OrmSQL.Commit()
}

/*func (fz *Fzmysql)UpdateImg(imgs []Image){
	for _, img := range imgs{
		var image Image
		sql := "select id,seriesuid from image where duns = '"+img.Duns+"' and dpath = '" +img.Dpath +"'"
		err := fz.OrmSQL.Raw(sql).QueryRow(&image)
		if err!=nil{
			nlogs.FileLogs.Error("sql error",img.Dpath)
			fz.OrmSQL.Insert(img)
		}else {
			if image.Seriesuid == "unknown" {
				img.Id = image.Id
				_,err := fz.OrmSQL.Update(img)
				if err!= nil{
					nlogs.ConsoleLogs.Info("repeat")
				}
			}else {
				fz.OrmSQL.Insert(img)
			}
		}
	}
}*/

func (fz *Fzmysql)ChangeDB(dbname string){
	fz.OrmSQL.Using(dbname)
}

func (fz *Fzmysql)InitDBList(){
	var dblist []Prothread
	_,err := fz.OrmSQL.QueryTable("prothread").All(&dblist)
	if err != nil{
		panic(err)
	}
	for db,_ := range dblist{
		fz.Dblist[dblist[db].Dbname] = &dblist[db]
		nlogs.ConsoleLogs.Debug(fz.Dblist[dblist[db].Dbname].Patientsql)
	}
}