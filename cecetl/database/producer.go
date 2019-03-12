package database

import (
	"github.com/astaxie/beego/orm"
	"zdy.local/utils/nlogs"
	"zdy.local/cecetl/message"
	"zdy.local/cecetl/kafka"
	"zdy.local/sql/mysql/mysqldriver"
	"fmt"
	"strconv"
	"time"
	"encoding/json"
	"strings"
)

type ProducePatient struct {
	Path string
	KafkaCli kafka.KafkaClient
	mysqldriver.Fzmysql
	sd SaveData
}

func (pro *ProducePatient)Init(dbname string){
	pro.sd.Init()
	pro.Path = pro.sd.Dblist[dbname].Pathlocate
	nlogs.ConsoleLogs.Debug(pro.Path)
	pro.KafkaCli.InitPorduce()
	dbUrl := pro.InitJndi(dbname)
	pro.OrmSQL,_ = pro.sd.SetOrmDriver(dbname,dbUrl)
}

/*
type Patient struct {
	SickID string		`json:"IDcard"`		4
	SickName string		`json:"name"`		5
	Age int			`json:"age"`		1
	Sex string		`json:"sex"`		3
	SickNamec string	`json:"namec"`		2
	PatientID string	`json:"patientid"`	7
	PatientUID string	`json:"patientuid"`	0
	BirthDay time.Time	`json:"birthday"`	6
}
 */

func (pro *ProducePatient)GetPatient(dbname string){
	var maps []orm.Params
	//pro.database.ChangeDB("jndi")
	nlogs.ConsoleLogs.Debug(dbname)
	sqls := strings.Split(pro.sd.Dblist[dbname].Patientsql,";")
	//num:=0
	for _,s := range sqls {
		_, err := pro.OrmSQL.Raw(s).Values(&maps)
		nlogs.ConsoleLogs.Debug(pro.sd.Dblist[dbname].Patientsql)
		if err != nil {
			nlogs.FileLogs.Error(err.Error())
		}
		local, _ := time.LoadLocation("UTC")
		for _, term := range maps {
			var result message.Private
			result.Duns = pro.sd.Dblist[dbname].Duns
			if fmt.Sprint(term["patientuid"]) == "" {
				continue
			}
			if fmt.Sprint(term["patientid"]) == "" {
				continue
			}
			result.SickInfo.SickID = fmt.Sprint(term["IDcard"])
			result.SickInfo.Age, _ = strconv.Atoi(strings.TrimSuffix(fmt.Sprint(term["age"]),"岁"))
			switch fmt.Sprint(term["sex"]) {
			case "M", "男":
				result.SickInfo.Sex = "男"
			case "F", "女":
				result.SickInfo.Sex = "女"
			default:
				result.SickInfo.Sex = "未知"
			}
			result.SickInfo.SickName = fmt.Sprint(term["name"])
			result.SickInfo.SickNamec = fmt.Sprint(term["namec"])
			var err error
			result.SickInfo.BirthDay, err = time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprint(term["birthday"]), local)
			if err != nil {
				nlogs.FileLogs.Warn("time parse error", err.Error())
				result.SickInfo.BirthDay, _ = time.Parse("2006-01-02 15:04:05 +0000 UTC", "0000-00-00 00:00:00 +0000 UTC")
			}
			result.SickInfo.PatientUID = fmt.Sprint(term["patientuid"])
			result.SickInfo.PatientID = fmt.Sprint(term["patientid"])
			result.Orgname = pro.sd.Dblist[dbname].Orgname
			result.Source = pro.sd.Dblist[dbname].Dbname
			text, _ := json.Marshal(result)
			pro.KafkaCli.Producter("upload", dbname, string(text))
			//num++
		}
	}
	//pro.KafkaCli.Producter("upload", "finish", "")
	defer pro.KafkaCli.CloseProducter()
	//nlogs.ConsoleLogs.Debug("producter finished ",nlogs.Fmt2String(num))
}