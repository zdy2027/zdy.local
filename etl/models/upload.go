package models

import (
	"zdy.local/cecetl/database"
	"zdy.local/utils/nlogs"
	"zdy.local/cecetl/kafka"
)

func UploadPACS(dbname string)(bool){
	var ms database.ProducePatient
	ms.Init(dbname)
	nlogs.ConsoleLogs.Info("START GET PATIENT")
	ms.GetPatient(dbname)
	return true
}

func DownloadPACS()(bool){
	var ms kafka.CliImpl
	ms = new(database.ConsumerCli)
	sign := ms.Init()
	defer ms.Close()
	if sign{
		kc := new(kafka.KafkaClient)
		kc.InitConsumer()
		kc.Consume(ms)
	}
	nlogs.ConsoleLogs.Alert("consum finished")
	return true
}

type IDTrans struct {
	DbName string
	TableName string
	SickID string
	ID_Card string
}

func DbTrans(db,table,row1,row2 string){
	var dbs database.Transdb
	dbs.Init(db)
	go dbs.TransFormdb(table,row1,row2)
}