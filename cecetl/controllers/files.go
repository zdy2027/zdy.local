package controllers

import (
	"github.com/astaxie/beego"
	"zdy.local/cecetl/models"
	"encoding/json"
)

type FileUploadController struct {
	beego.Controller
}

// @Title Auto sql study
// @Description 反向上传pacs文件
// @Param	body		body 	models.Data	true		"每个字段对应的字段名以及来源信息"
// @Success 200 {string} sql
// @Failure 403 body is empty
// @router / [post]
func (file *FileUploadController) Post() {
	var ob models.Data
	json.Unmarshal(file.Ctx.Input.RequestBody, &ob)
	go models.Start(ob)
	file.Data["json"] = map[string]string{"result": "start run"}
	file.ServeJSON()

}