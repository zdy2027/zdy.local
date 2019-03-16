package controllers

import (
	"zdy.local/cecetl/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

// Operations about logs
type ProducterController struct {
	beego.Controller
}

// @Title Upload pacs
// @Description 查询病人信息，上传数据至kafka待分析存储
// @Param	dbname		query 	string	true		"上传医院数据库名"
// @Success 200 {string} 开始执行pacs上传程序
// @Failure 403 body is empty
// @router / [get]
func (t *ProducterController) Get(){
	dbname := t.GetString("dbname")
	go models.UploadPACS(dbname)
	t.Data["json"] = map[string]string{"result": "start upload pacs"}
	t.ServeJSON()
}

// @Title update sick id card
// @Description 查询病人身份证信息
// @Param	body		body 	models.IDTrans	true		"输入存储病人身份证信息和sickid的表信息"
// @Success 200 {string} 开始身份证更新程序
// @Failure 403 body is empty
// @router /dbtransform [post]
func (t *ProducterController) DbTrans(){
	var ob models.IDTrans
	json.Unmarshal(t.Ctx.Input.RequestBody, &ob)
	go models.DbTrans(ob.DbName,ob.TableName,ob.SickID,ob.ID_Card)
	t.Data["json"] = map[string]string{"result": "start update sick id card"}
	t.ServeJSON()
}