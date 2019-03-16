package database

import (
	"zdy.local/utils/nlogs"
	"zdy.local/cecetl/nosql"
	"zdy.local/sql/mysql/mysqldriver"
)

type GetMemCache struct {
	nosql.RedisConn
	//mysqldriver.Fzmysql
	sd SaveData
}

func (gm *GetMemCache)InitCache(){
	gm.InitRedis()
	gm.sd.Init()
}

func (gm *GetMemCache)GetImages(duns string){
	var image []mysqldriver.Image
	sql := "SELECT id,dpath,seriesuid from image where duns='"+duns+"';"
	nlogs.ConsoleLogs.Debug(sql)
	_,err := gm.sd.OrmSQL.Raw(sql).QueryRows(&image)
	if err!=nil{
		nlogs.ConsoleLogs.Error(err.Error())
	}else {
		for _,img := range image{
			nlogs.ConsoleLogs.Debug(nlogs.Fmt2String(img.Id),img.Dpath,img.Seriesuid)
			gm.SetMessage(img.Dpath,img)
		}
		nlogs.ConsoleLogs.Info("finish")
	}
	gm.Close()
}