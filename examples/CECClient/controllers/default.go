package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Layout = "layout.html"
	this.TplName = "index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Head"] = "templates/indexHead.tpl"
	this.LayoutSections["Scripts"] = "templates/comScripts.tpl"
	this.Render()
}