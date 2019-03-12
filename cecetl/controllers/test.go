package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about test
type TestController struct {
	beego.Controller
}

// @Title Get
// @Description get data
// @Success 200 {string} test
// @Failure 403 body is empty
// @router / [get]
func (t *TestController) Get(){
	//var ob message.Data
	//json.Unmarshal(t.Ctx.Input.RequestBody, &ob)
	//models.Test(ob)
	var result = ""
	for i:=0;i<100000;i++{
		result += "a"
	}
	t.Data["json"] = map[string]string{"result": result}
	t.ServeJSON()
}
