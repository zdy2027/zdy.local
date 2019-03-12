package mysqldriver

import (
	"time"
)

type His_lis_pacs struct {
	Lis_id string `orm:"pk"`
	Sfzh string
	Hisid string
	Type string
	Orgname string
	Duns string
	Name string
	Rawname string
	Typename string
	Measuredate time.Time
	Instid string
	Status int
	Created_date string `orm:"auto_now_add;type(datetime)"`
	Created_by string
	Updated_by string
	Updated_date time.Time `orm:"auto_now;type(datetime)"`
	Comx string
	Created_way string
	Urlx string
	Is_indexed int
	Filesize int64
	Minijpg string
	Jpg string
	Sick_id string
}

type Patient struct {
	Patientid string	`orm:"pk"`
	Patientuid string
	Sickid string
	Sickname string
	Age int
	Sex string
	Sicknamec string
	Birthday time.Time
	Duns string
	Source string
}

type Series struct {
	Studyuid string
	Seriesuid string	`orm:"pk"`
	Duns string
	Source string
	Seriesnumber int
	Imagecount int
	Studydatetime time.Time
	Modality string
	Exambodypart string
	Sickid string
	Patientuid string
}

type Image struct {
	Id int 				`orm:"auto"`
	Seriesuid string
	Imagename string	
	Duns string
	Source string
	Urlx string
	Size int64
	Typex string
	Dpath string
	Imagenum int
	Sickid string
	Patientuid string
	Studyuid string
	Prepath string
}

type Study struct {
	Duns string
	Source string
	Studyuid string
	Patientuid string
	Bodypart string
	Department string
	Clinical string
	Accessionnumber string
	Reportid string
	Sickid string
	Studytime time.Time
	//Patientsource string
	Devicename string
	Id int			`orm:"auto"`
}

type Report struct{
	Studyid string
	Patientuid string
	Reportid string 	`orm:"pk"`
	Duns string
	Source string
	Report string
	Reportlocate string
	Reporturl string
	Result string
	Resultlocate string
	Resulturl string
	Opinion string
	Description1 string
	Conclusion1 string
	Description2 string
	Conclusion2 string
	Info1 string
	Info2 string
	Part string
	Sickid string
}

type EtlJndi struct {
	Jndi string		`orm:"pk"`
	Type string
	Status int
	Val string
	Comx string
	Duns string
	Dunsname string
}

type Prothread struct{
	Duns string
	Orgname string
	Dbname string		`orm:"pk"`
	Patientsql string
	Studysql string
	Seriessql string
	Imagesql string
	Reportsql string
	Filesql string
	//Typex string
	//Source string
	Updateddate time.Time `orm:"auto_now;type(datetime)"`
	Uploadnumber int
	Filenumber int
	Pathlocate string
	Totalnum int
}