package kafka

import "zdy.local/cecetl/message"

type CliImpl interface {
	Init()bool
	GetStudyUID(dbname,patientUID string,res *message.Private)
	LoadData(dbname string,result message.Private)
	Close()
}