package controllers

import (
	"github.com/astaxie/beego"
)

type ConfController struct {
	beego.Controller
}

func (this *ConfController) Get() {
	this.Layout = "layout.html"
	this.TplName = "conf.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Head"] = "templates/confHead.tpl"
	this.LayoutSections["Scripts"] = "templates/comScripts.tpl"
}