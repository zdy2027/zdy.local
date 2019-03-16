package controllers

import (
	"github.com/astaxie/beego"
	"zdy.local/cecetl/models"
	"encoding/json"
)

var (
	sql string
)

type AutoSqlController struct {
	beego.Controller
}

// @Title Auto sql study
// @Description 根据json字段自动生成对表sql语句
// @Param	body		body 	models.StudySql	true		"每个字段对应的字段名以及来源信息"
// @Success 200 {string} sql
// @Failure 403 body is empty
// @router /AutoStudy [post]
func (study *AutoSqlController) AutoStudy() {
	var ob models.StudySql
	json.Unmarshal(study.Ctx.Input.RequestBody, &ob)
	sql = models.AutoStudy(ob)
	study.Data["json"] = map[string]string{"sql": sql}
	study.ServeJSON()

}

// @Title Auto sql patient
// @Description 根据json字段自动生成对表sql语句
// @Param	body		body 	models.PatientSql	true		"病人信息属性表"
// @Success 200 {string} sql
// @Failure 403 body is empty
// @router /AutoPatient [post]
func (patient *AutoSqlController) AutoPatient() {
	var ob models.PatientSql
	json.Unmarshal(patient.Ctx.Input.RequestBody, &ob)
	sql = models.AutoPatient(ob)
	patient.Data["json"] = map[string]string{"sql": sql}
	patient.ServeJSON()
}

// @Title Auto sql series
// @Description 根据json字段自动生成对表sql语句
// @Param	body		body 	models.SeriesSql	true		"每个字段对应的字段名以及来源信息"
// @Success 200 {string} sql
// @Failure 403 body is empty
// @router /AutoSeries [post]
func (series *AutoSqlController) AutoSeries() {
	var ob models.SeriesSql
	json.Unmarshal(series.Ctx.Input.RequestBody, &ob)
	sql = models.AutoSeries(ob)
	series.Data["json"] = map[string]string{"sql": sql}
	series.ServeJSON()
}

// @Title Auto sql image
// @Description 根据json字段自动生成对表sql语句
// @Param	body		body 	models.ImageSql	true		"每个字段对应的字段名以及来源信息"
// @Success 200 {string} sql
// @Failure 403 body is empty
// @router /AutoImage [post]
func (image *AutoSqlController) AutoImage() {
	var ob models.ImageSql
	json.Unmarshal(image.Ctx.Input.RequestBody, &ob)
	sql = models.AutoImages(ob)
	image.Data["json"] = map[string]string{"sql": sql}
	image.ServeJSON()
}

// @Title Auto sql report
// @Description 根据json字段自动生成对表sql语句
// @Param	body		body 	models.ReportSql	true		"每个字段对应的字段名以及来源信息"
// @Success 200 {string} sql
// @Failure 403 body is empty
// @router /AutoReport [post]
func (report *AutoSqlController) AutoReport() {
	var ob models.ReportSql
	json.Unmarshal(report.Ctx.Input.RequestBody, &ob)
	sql = models.AutoReport(ob)
	report.Data["json"] = map[string]string{"sql": sql}
	report.ServeJSON()
}