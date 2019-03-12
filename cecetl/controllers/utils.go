package controllers

import (
	"github.com/astaxie/beego"
	"zdy.local/cecetl/models"
	"zdy.local/fileOP"
)

type EsController struct {
	beego.Controller
}

// @Title serach in es
// @Description 根据身份证号查找es数据库
// @Param	idcard		query 	string	true		"返回病人影像数据"
// @Success 200 {object} message.PrivateHits
// @Failure 403 idcard is empty
// @router /essearch [get]
func (es *EsController)Get(){
	idcard := es.GetString("idcard")
	es.Data["json"] = models.GetPatients(idcard)
	//es.Data["json"],_ = json.Marshal(result)
	es.ServeJSON()
}

// @Title load json file
// @Description 加载json文件
// @Param	jsonpath		query 	string	true		"加载es备份的json文件"
// @Param	copydir			query 	string	true		"拷贝文件存储地址"
// @Success 200 {string} succeed
// @Failure 403 jsonpath is empty
// @router /read/json [get]
func (es *EsController)ReadJson(){
	path := es.GetString("jsonpath")
	copydir := es.GetString("copydir")
	go models.ReadJson(path,copydir)
	//es.Data["json"],_ = json.Marshal(result)
	//es.ServeJSON()
}

// @Title load json file and save in es
// @Description 加载json文件
// @Param	jsonpath		query 	string	true		"加载es备份的json文件"
// @Param	copydir			query 	string	true		"拷贝文件存储地址"
// @Success 200 {string} succeed
// @Failure 403 jsonpath is empty
// @router /save/json [get]
func (es *EsController)SaveJson(){
	path := es.GetString("jsonpath")
	copydir := es.GetString("copydir")
	go models.SaveJson(path,copydir)
	//es.Data["json"],_ = json.Marshal(result)
	//es.ServeJSON()
}

// @Title upload files in weedfs
// @Description 上传文件到weedfs
// @Param	filepath		query 	string	true		"上传文件到weedfs"
// @Success 200 {string} start upload
// @Failure 403 filepath is empty
// @router /upload [get]
func (es *EsController)UploadFile(){
	path := es.GetString("filepath")
	var obj fileOP.FileOP
	obj = new(models.Upload)
	obj.Init("")
	fileOP.Run(path,"",obj)
	//es.Data["json"],_ = json.Marshal(result)
	//es.ServeJSON()
}

// @Title cache images to mysql
// @Description 将mysql中image数据缓存至redis
// @Param	duns		query 	string	true		"将mysql中的image数据缓存至redis"
// @Success 200 {string} start upload
// @Failure 403 filepath is empty
// @router /Mysql2Redis [get]
func (es *EsController)Mysql2Redis(){
	path := es.GetString("duns")
	go models.CacheImg(path)
}

// @Title cache images to mysql
// @Description 删除weedfs文件
// @Param	duns		query 	string	true		"删除his_lis_pac中weedfs存储文件"
// @Success 200 {string} start upload
// @Failure 403 filepath is empty
// @router /DelWeed [get]
func (es *EsController)DelWeed(){
	path := es.GetString("duns")
	go models.CacheImg(path)
}

// @Title load json file and save in es
// @Description 新添加病人信息
// @Param	jsonpath		query 	string	true		"加载es备份的json文件"
// @Param	copydir			query 	string	true		"拷贝文件存储地址"
// @Success 200 {string} succeed
// @Failure 403 jsonpath is empty
// @router /add/data [get]
func (es *EsController)AddData(){
	path := es.GetString("jsonpath")
	copydir := es.GetString("copydir")
	go models.AddData(path,copydir)
	//es.Data["json"],_ = json.Marshal(result)
	//es.ServeJSON()
}