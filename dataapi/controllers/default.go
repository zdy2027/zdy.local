package controllers

import (
	"github.com/astaxie/beego"
	"github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "john" {
		// password is "hello"
		return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"
	}
	return ""
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Prepare() {
	a := auth.NewBasicAuthenticator("example.com", Secret)
	if username := a.CheckAuth(this.Ctx.Request); username == "" {
		a.RequireAuth(this.Ctx.ResponseWriter, this.Ctx.Request)
	}
}

func (c *MainController) Get() {
	c.Data["Username"] = "astaxie"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
