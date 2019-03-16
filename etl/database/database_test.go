package database

import (
	"testing"
	//"zdy.local/cecetl/message"
	"zdy.local/sql/mysql/jkda"
	"strconv"
)

func TestMysql_MultiSave(t *testing.T) {
	/*
	var m Mysql
	m.Init()
	p := Patient{Patientid:"12345x",Age:10,Sickname:"中文"}
	var patients []Patient
	patients = append(patients,p)
	m.MultiSave(patients)

	study := Study{Orgname:"test",Duns:"12345",Studyuid:"1.2.3.3",Patientid:"12345x",Bodypart:"chest"}
	var studys []Study
	studys = append(studys,study)
	m.MultiSave(studys)

	serie := Series{Seriesuid:"1.2.3.3.5",Studyuid:"1.2.3.3",Count:4}
	var series []Series
	series = append(series,serie)
	m.MultiSave(series)

	img := Image{Imagename:"test.dcm",Seriesuid:"1.2.3.3.5"}
	var imgs []Image
	imgs = append(imgs,img)
	m.MultiSave(imgs)*/
}

func TestMssql_WriteImg(t *testing.T) {
	/*
	res := fileOP.GetPath("/Users/zhangdongyao/go")
	fmt.Println(res)
	test := "/Users/zhangdongyao/go/bin/zdy/test/a.dcm"
	for _,p := range res {
		if strings.Contains(test,p){
			fmt.Println("/new/path"+test[strings.Index(test,p)-1:])
		}
	}
	*/
}

func TestMssql_GetSeriesUID(t *testing.T) {
	//var ms = Mssql{}
	/*
	ms := new(Mssql)
	input := message.Data{Server:"localhost",Msdb:"dbo"}
	ms.Init(input)
	//sql := fmt.Sprintf("SELECT exambodypart FROM Series WHERE StudyInstanceUID = '%s';","12345")
	sql := "SELECT * FROM Series;"
	rows ,err := ms.Db.Exec(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
	/*
	cols,_ := rows.Columns()
	var colsdata = make ([]interface{},len(cols))
	for i:=0;i<len(cols);i++ {
		colsdata[i] = new(interface{})
	}
	for rows.Next() {
		rows.Scan(colsdata...)
		fmt.Println("%s",*(colsdata[0].(*interface{})))
	}*/
}

func Test_Upsert(t *testing.T){
	Jkdasql.OrmSQL.Begin()
	var patient []jkda.Patient
	for i:=0;i<3 ;i++  {
		p := jkda.Patient{Patientid:"1001"+strconv.Itoa(i),Patientuid:"123"+strconv.Itoa(i),Age:i+30,Sickname:"ceshi"}
		//Jkdasql.UpSert(&p)
		patient = append(patient,p)
	}
	Jkdasql.MultiUpSert(patient)
	Jkdasql.OrmSQL.Commit()
}