package database

import (
	"encoding/json"
	"zdy.local/cecetl/message"
	"zdy.local/sql/mysql/mysqldriver"
	"zdy.local/utils/nlogs"
	"zdy.local/cecetl/nosql"
	"fmt"
	"sync"
)

type SaveData struct {
	mysqldriver.Fzmysql
	Es nosql.ES
	sign bool
	uploadNum map[string]int
	m *sync.Mutex
}

func (sd *SaveData)Init() {
	sd.uploadNum = make(map[string]int)
	sd.m = new(sync.Mutex)

	sd.InitJkda()
	for db,_ := range sd.Dblist{
		sd.uploadNum[db] = 0
	}
	sd.Es.Init()
}

func (sd *SaveData)SaveData(result message.Private,dbname string)  {
	//Database.ChangeDB("default")
	//m.Lock()
	//Database.Sql.OrmSQL.Begin()
	sick := mysqldriver.Patient{
		Patientid:result.SickInfo.PatientID,
		Patientuid:result.SickInfo.PatientUID,
		Sickid:result.SickInfo.SickID,
		Sickname:result.SickInfo.SickName,
		Sicknamec:result.SickInfo.SickNamec,
		Sex:result.SickInfo.Sex,
		Age:result.SickInfo.Age,
		Birthday:result.SickInfo.BirthDay,
		Duns:result.Duns,
		Source:result.Source,
	}
	nlogs.ConsoleLogs.Debug(fmt.Sprint(sick.Birthday))
	_,err := sd.UpSert(&sick)
	if err != nil {
		nlogs.FileLogs.Error("SAVE PATIENT ERROR ",result.SickInfo.PatientID)
		nlogs.FileLogs.Error("SAVE PATIENT ERROR MESSAGE",err.Error())
		return
	}
	var studys []mysqldriver.Study
	var series []mysqldriver.Series
	var images []mysqldriver.Image
	var reports []mysqldriver.Report
	for _,study_data := range result.Studys{
		//study := jkda.Study{}
		study := mysqldriver.Study{
			Duns:result.Duns,
			Patientuid:result.SickInfo.PatientUID,
			Studyuid:study_data.StudyID,
			Bodypart:study_data.BodyPart,
			Department:study_data.Department,
			Clinical:study_data.Clinical,
			Reportid:study_data.Reportx.ReportID,
			Accessionnumber:study_data.Accessionnumber,
			Source:result.Source,
			Sickid:result.SickInfo.SickID,
		}
		studys = append(studys,study)
		for _,series_data := range study_data.SeriesInfo{
			serie := mysqldriver.Series{
				Studyuid:study_data.StudyID,
				Seriesuid:series_data.SeriesID,
				Studydatetime:series_data.DateTimex,
				Exambodypart:series_data.BodyPart,
				Imagecount:series_data.Count,
				Seriesnumber:series_data.SeriesNum,
				Modality:series_data.Modality,
				Duns:result.Duns,
				Source:result.Source,
				Sickid:result.SickInfo.SickID,
				Patientuid:result.SickInfo.PatientUID,
			}
			for _,img_data := range series_data.Imgs {
				if img_data.Urlx != "" {
					sd.m.Lock()
					sd.uploadNum[dbname]++
					nlogs.ConsoleLogs.Alert("save number is",sd.uploadNum[dbname])
					sd.m.Unlock()
				}
				image := mysqldriver.Image{
					Seriesuid:series_data.SeriesID,
					Imagename:img_data.Imgname,
					Size:img_data.Size,
					Urlx:img_data.Urlx,
					Typex:img_data.Typex,
					Dpath:img_data.Dpath,
					Duns:result.Duns,
					Source:result.Source,
					Imagenum:img_data.Imagenum,
					Patientuid:result.SickInfo.PatientUID,
					Sickid:result.SickInfo.SickID,
					Studyuid:study_data.StudyID,
				}
				images = append(images,image)
				//}
			}
			series = append(series,serie)
		}
		report := mysqldriver.Report{
			Studyid:study_data.StudyID,
			Patientuid:result.SickInfo.PatientUID,
			Reportid:study_data.Reportx.ReportID,
			Report:study_data.Reportx.Report,
			Reportlocate:study_data.Reportx.ReportLocate,
			Reporturl:study_data.Reportx.ReportUrl,
			Result:study_data.Reportx.Result,
			Resultlocate:study_data.Reportx.ResultLocate,
			Resulturl:study_data.Reportx.ResultUrl,
			Opinion:study_data.Reportx.Opinion,
			Description1:study_data.Reportx.Description1,
			Conclusion1:study_data.Reportx.Conclusion1,
			Description2:study_data.Reportx.Description2,
			Conclusion2:study_data.Reportx.Conclusion2,
			Info1:study_data.Reportx.Info1,
			Info2:study_data.Reportx.Info2,
			Part:study_data.Reportx.BodyPart,
			Duns:result.Duns,
			Source:result.Source,
			Sickid:result.SickInfo.SickID,
		}
		reports = append(reports,report)
	}
	//_,err = jkda.MultiSave(studys)
	sd.MultiUpSert(studys,"patientuid","studyuid","duns")//MultiUpSert(studys)
	sd.MultiUpSert(series)//MultiSave(series)
	sd.UpdateImg(images)
	sd.MultiUpSert(reports)//MultiUpSert(reports)
	sd.m.Lock()
	sd.Dblist[dbname].Uploadnumber = sd.uploadNum[dbname]
	sd.m.Unlock()
	sd.UpSert(sd.Dblist[dbname])
	//Database.Sql.OrmSQL.Commit()
	//Jkdasql.OrmSQL.Commit()
	nlogs.ConsoleLogs.Debug("START SAVE ES")
	nlogs.FileLogs.Info("SAVE SICK",result.SickInfo.SickID)
	text,_ := json.Marshal(result)
	//Es.IndexData(string(text),result.SickInfo.SickID,"pacs")
	nlogs.ConsoleLogs.Debug(string(text))
	sd.Es.IndexData(result.Duns+result.SickInfo.PatientID+dbname,"pacs",string(text))
	//m.Unlock()
}

func (sd *SaveData)SaveFile(imgname string) (string,int64){
	fid,size := nosql.UploadWeed(imgname)
	/*
	var dcm message.DCM
	dcm.Pacs = dicom.DCMExtract(imgname)
	dcm.SickID = idcard
	//dcm.StudyID = studyid
	dcm.SeriesID = seriesid
	dcm.Datetimex = studytime
	text,_ := json.Marshal(dcm)
	Es.IndexData(string(text),fid,"dicom")
	*/
	return fid,size
}

func (sd *SaveData)UpdateImg(imgs []mysqldriver.Image){
	for _, img := range imgs{
		var image mysqldriver.Image
		sql := "select id,seriesuid,studyuid,prepath from image where duns = '"+img.Duns+"' and dpath = '" +img.Dpath +"'"
		err := sd.OrmSQL.Raw(sql).QueryRow(&image)
		//image,err := sd.GetMessage(strings.Split(img.Dpath,"/")[len(strings.Split(img.Dpath,"/"))-1] + "/" + img.Imagename)
		if err!=nil{
			nlogs.FileLogs.Error("sql error",err.Error(),img.Dpath)
			_,err = sd.OrmSQL.Insert(&img)
			if err != nil{
				nlogs.FileLogs.Error("insert error",err.Error())
			}
		}else {
			img.Prepath = image.Prepath
			//nlogs.ConsoleLogs.Alert("prepath is",image.Prepath,image.Studyuid)
			if image.Seriesuid == "unknown" {
				img.Id = image.Id
				_,err := sd.OrmSQL.Update(&img)
				if err!= nil{
					nlogs.ConsoleLogs.Info("repeat")
					nlogs.FileLogs.Debug("repeat sql",err,img.Prepath)
				}
			}else {
				if img.Seriesuid == image.Seriesuid && img.Studyuid == image.Studyuid {
					continue
				}else {
					_,err = sd.OrmSQL.Insert(&img)
					if err!= nil{
						nlogs.ConsoleLogs.Info("repeat",img.Dpath)
						nlogs.FileLogs.Debug("repeat sql",err,img.Prepath,img.Seriesuid,image.Id)
					}
				}
			}
		}
	}
}