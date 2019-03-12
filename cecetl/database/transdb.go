package database

import (
	"github.com/astaxie/beego/orm"
	"zdy.local/sql/mysql/mysqldriver"
	"zdy.local/utils/nlogs"
	"zdy.local/cecetl/message"
	"fmt"
	"zdy.local/cecetl/nosql"
	"encoding/json"
	"regexp"
)

type Transdb struct {
	Driver mysqldriver.Fzmysql
	esclient nosql.ES
	sd SaveData
}

func (ts *Transdb)Init(db string){
	ts.sd.Init()
	ts.esclient.Init()
	dbUrl := ts.Driver.InitJndi(db)
	ts.Driver.OrmSQL,_ = ts.Driver.SetOrmDriver(db,dbUrl)
}

func (ts *Transdb)TransFormdb(table,row1,row2 string){
	var maps []orm.Params
	ts.Driver.OrmSQL.Raw("select "+row1+","+row2+" from "+table).Values(&maps)
	reg,err :=  regexp.Compile("^(\\d{6})(\\d{8})(.*)")
	if err != nil {
		return
	}
	for _,term := range maps{
		if reg.MatchString(fmt.Sprint(term[row2])) != true{
			nlogs.ConsoleLogs.Debug("not find id card",term[row2])
			continue
		}
		nlogs.ConsoleLogs.Info("id_card is",term[row2])
		var result message.PrivateHits
		var patients []mysqldriver.Patient
		num,err:=ts.sd.OrmSQL.QueryTable("patient").Filter("SICKID",fmt.Sprint(term[row1])).All(&patients)
		if err!=nil || num ==0{
			continue
		}
		for _,p := range patients{
			p.Sickid = fmt.Sprint(term[row2])
			ts.sd.OrmSQL.Update(&p)
		}
		var studys []mysqldriver.Study
		ts.sd.OrmSQL.QueryTable("study").Filter("SICKID",fmt.Sprint(term[row1])).All(&studys)
		for _,p := range studys{
			p.Sickid = fmt.Sprint(term[row2])
			ts.sd.OrmSQL.Update(&p)
		}
		var series []mysqldriver.Series
		ts.sd.OrmSQL.QueryTable("series").Filter("SICKID",fmt.Sprint(term[row1])).All(&series)
		for _,p := range series{
			p.Sickid = fmt.Sprint(term[row2])
			ts.sd.OrmSQL.Update(&p)
		}
		var image []mysqldriver.Image
		ts.sd.OrmSQL.QueryTable("image").Filter("SICKID",fmt.Sprint(term[row1])).All(&image)
		for _,p := range image{
			p.Sickid = fmt.Sprint(term[row2])
			ts.sd.OrmSQL.Update(&p)
		}
		var report []mysqldriver.Report
		ts.sd.OrmSQL.QueryTable("report").Filter("SICKID",fmt.Sprint(term[row1])).All(&report)
		for _,p := range report{
			p.Sickid = fmt.Sprint(term[row2])
			ts.sd.OrmSQL.Update(&p)
		}
		nlogs.FileLogs.Info("UPDATE sick",fmt.Sprint(term[row1]))
		result = ts.esclient.GetPatients(fmt.Sprint(term[row1]))
		for _,e := range result.Info {
			e.SickInfo.SickID = fmt.Sprint(term[row2])
			text,_ := json.Marshal(e)
			ts.esclient.IndexData(e.Duns+e.SickInfo.PatientID+e.Source,"pacs",string(text))
		}
	}
	nlogs.ConsoleLogs.Alert("finished")
}