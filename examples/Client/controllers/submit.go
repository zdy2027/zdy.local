package controllers

import (
	"encoding/json"
	"zdy.local/CECClient/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/toolbox"
)

var iniconf, _ = config.NewConfig("ini", "conf/app.conf")

// Submit api
type SubmitController struct {
	beego.Controller
}

type SubmitObject struct {
	Path string	`json:"upload"`
	PrePath string	`json:"preupload"`
	ServerIP string	`json:"server"`
	ServerPort string	`json:"lastmodify"`
	Lastmodify string	`json:"port"`
	Os string		`json:"os"`
	UpdateCycle string	`json:"update_cycle"`
	UploadCycle string	`json:"upload_cycle"`
}

// @Title submit
// @Description submit object
// @Param	body		body 	SubmitObject	true		"The object content"
// @Success 200 {string} SubmitObject.Path
// @Failure 403 body is empty
// @router / [post]
func (this *SubmitController) Post() {
	var ob SubmitObject
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	iniconf.Set("dev::lastmodify",ob.Lastmodify)
	iniconf.Set("dev::path",ob.Path)
	iniconf.Set("dev::prepath",ob.PrePath)
	iniconf.Set("dev::server",ob.ServerIP)
	iniconf.Set("dev::port",ob.ServerPort)
	iniconf.Set("dev::os",ob.Os)
	iniconf.Set("dev::update",ob.UpdateCycle)
	iniconf.Set("dev::upload",ob.UploadCycle)
	iniconf.SaveConfigFile("conf/app.conf")
	this.Data["json"] = map[string]string{"status": "succeed"}
	this.ServeJSON()
	toolbox.StopTask()
	tk1 := toolbox.NewTask("tk1", "0/3 * * * * *", func() error{
		models.UploadCycle()
		return nil
	})
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()
	this.Render()
}