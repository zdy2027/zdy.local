package controllers

import (
	"zdy.local/cecetl/models"
	"github.com/astaxie/beego"
)

// Operations about logs
type DownLoadController struct {
	beego.Controller
}

// @Title Extract pacs
// @Description 启动kafka消费者，根据病人信息获取医学影像数据存储
// @Success 200 {string} 开始执行pacs上传程序
// @Failure 403 body is empty
// @router / [get]
func (t *DownLoadController) Get(){
	go models.DownloadPACS()
	t.Data["json"] = map[string]string{"result": "start upload pacs"}
	t.ServeJSON()
}