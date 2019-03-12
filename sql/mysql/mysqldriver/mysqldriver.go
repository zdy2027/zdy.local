package mysqldriver

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"zdy.local/utils/nlogs"
	"database/sql"
	"reflect"
	"fmt"
)

type MysqlDriver struct {
	OrmSQL orm.Ormer
}

func (ms *MysqlDriver)RegisterDatabase(alias,dburl string)  {
	err := orm.RegisterDataBase(alias, "mysql", dburl, 30)
	if err != nil {
		panic(err)
	}
}

func (ms *MysqlDriver)RegisterModelx(modelx interface{}){
	orm.RegisterModel(modelx)
}

func (ms *MysqlDriver)SetOrm(){
	ms.OrmSQL = orm.NewOrm()
}

func (ms *MysqlDriver)SetOrmDriver(aliasName,dburl string) (orm.Ormer,error){
	dbdriver, err := sql.Open("mysql", dburl)
	if err != nil {
		nlogs.ConsoleLogs.Error(err.Error())
		panic(err)
	}
	dbdriver.SetMaxOpenConns(2000)
	dbdriver.SetMaxIdleConns(1000)
	//local, _ := time.LoadLocation("UTC")
	//orm.AddAliasWthDB(aliasName,"mysql",dbdriver)
	//orm.SetDataBaseTZ(aliasName,local)
	return orm.NewOrmWithDB("mysql",aliasName,dbdriver)
}

func (ms *MysqlDriver)GetOrm(dbName string)(*sql.DB,error){
	return orm.GetDB(dbName)
}

func (ms *MysqlDriver)UpSert(data interface{})(int64,error){
	//nlogs.ConsoleLogs.Debug(fmt.Sprint(data))
	if num, err := ms.OrmSQL.InsertOrUpdate(data); err != nil {
		//panic(err)
		return num,err
	} else {
		return num,err
	}
}

func (ms *MysqlDriver)MultiUpSert(pacs interface{},confli ...string) {
	sind := reflect.Indirect(reflect.ValueOf(pacs))
	num := sind.Len()
	//ms.OrmSQL.Begin()
	for i:=0;i<num;i++{
		_,err := ms.OrmSQL.InsertOrUpdate(sind.Index(i).Addr().Interface(),confli...)
		if err != nil {
			//ms.OrmSQL.Rollback()
			nlogs.ConsoleLogs.Debug("UPSERT DATA ERROR ",err)
			nlogs.ConsoleLogs.Debug("UPSERT DATA ",fmt.Sprint(sind.Index(i)))
			nlogs.FileLogs.Error("UPSERT DATA ERROR ",err)
			nlogs.FileLogs.Error("UPSERT DATA ",fmt.Sprint(sind.Index(i)))
			//panic(err)
		}
	}
	//ms.OrmSQL.Commit()
}

func (ms *MysqlDriver)MultiSave(pacs interface{}) (int64,error){
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
				bulk := sind.Slice(i,sind.Len()).Len()
				n,err := ms.OrmSQL.InsertMulti(bulk, sind.Slice(i,sind.Len()).Interface())
				if err != nil {
					//panic(err)
					return totalnum,err
				}else {
					if n<int64(bulk){
						nlogs.FileLogs.Notice("there are some sql error",n,bulk)
					}
					totalnum += int64(bulk)
				}
			}else {
				nlogs.ConsoleLogs.Debug(nlogs.Fmt2String(i),nlogs.Fmt2String(i+num))
				bulk := sind.Slice(i,i+num).Len()
				n,err := ms.OrmSQL.InsertMulti(bulk, sind.Slice(i,i+num).Interface())
				if err != nil {
					//panic(err)
					return totalnum,err
				}else {
					if n<int64(bulk){
						nlogs.FileLogs.Notice("there are some sql error",n,bulk)
					}
					totalnum += int64(bulk)
				}
			}
		}
		return totalnum,nil
	}else {
		if num, err := ms.OrmSQL.InsertMulti(sind.Len(), pacs); err != nil {
			//panic(err)
			return num,err
		} else {
			return num,err
		}
	}
}