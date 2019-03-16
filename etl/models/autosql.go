package models

import (
	"strings"
)

type StudySql struct {
	Source string			//查询表名
	Studyuid string			//study唯一标识
	Patientuid string		//
	Bodypart string			//检查部位
	Department string		//检查科室
	Clinical string			//临床诊断
	Accessionnumber string		//
	Reportid string			//检查报告id
	Devicename string		//设备名
	Studytime string		//研究日期
}

//SELECT StudyInstanceUID,StudyInformation,department,linchuangzhenduan,AccessionNumber,ReportID FROM Study WHERE PatientUID = '%s';

func AutoStudy(ob StudySql)string{
	if ob.Studyuid == "string" || ob.Studyuid =="" {
		return "Studyuid can not empty"
	}
	sql := "SELECT " + ob.Accessionnumber + " AS accessionnumber," +
	ob.Bodypart + " AS bodypart," + ob.Department + " AS department," +
	ob.Reportid + " AS reportid," +
	ob.Clinical + " AS clinical," + ob.Studyuid + " AS StudyInstanceUID " +
	"FROM " + ob.Source + " WHERE " + ob.Patientuid + "='%s'"
	tmp := strings.Split(sql,",")
	sql = ""
	for _,t := range tmp {
		if strings.Contains(t,"string AS "){
			continue
		}
		sql += t + ","
	}
	sql = strings.TrimSuffix(sql,",") + ";"
	return sql
}

type SeriesSql struct {
	Studyuid string		//所属study唯一标识
	Seriesuid string	//series唯一标识
	Source string		//
	Seriesnumber string	//series序号
	Imagecount string	//包含影像数量
	Studydatetime string	//研究日期
	Modality string		//诊断类型CT、US（超声）、ECG（心电图）、MR（脑部）
	Exambodypart string	//检查部位
	Comx string		//备注
}

func AutoSeries(ob SeriesSql)string{
	if ob.Seriesuid == "string" || ob.Seriesuid =="" {
		return "Seriesuid can not empty"
	}
	sql := "SELECT " + ob.Studydatetime + " AS datetimex," +
		ob.Modality + " AS modality," + ob.Comx + " AS comx," +
		ob.Imagecount + " AS count," + ob.Seriesnumber + " AS seriesnum," +
		ob.Exambodypart + " AS bodypart," + ob.Seriesuid + " AS seriesID " +
		"FROM " + ob.Source + " WHERE " + ob.Studyuid + "='%s'"
	tmp := strings.Split(sql,",")
	sql = ""
	for _,t := range tmp {
		if strings.Contains(t,"string AS "){
			continue
		}
		sql += t + ","
	}
	sql = strings.TrimSuffix(sql,",") + ";"
	return sql
}

type PatientSql struct {
	Patientid string	//患者唯一标识
	Patientuid string	//对应关联study
	Sickid string		//身份证号
	Sickname string		//患者姓名
	Age string		//年龄
	Sex string		//性别
	Sicknamec string	//姓名 英文
	Birthday string		//出生日期
	Source string		//
}
//SELECT PatientUID,PatientAge,PatientNameC,PatientSex,PATIENT_MEDICARENO,PatientName,PatientBirthday,PatientID FROM Patient;
func AutoPatient(ob PatientSql)string{
	if ob.Patientuid == "string" || ob.Patientuid =="" {
		return "Patientuid can not empty"
	}
	sql := "SELECT " + ob.Patientid + " AS patientid," +
		ob.Birthday + " AS birthday," + ob.Sickid + " AS IDcard," +
		ob.Age + " AS age," + ob.Sex + " AS sex," +
		ob.Sickname + " AS name," + ob.Sicknamec + " AS namec," +
		ob.Patientuid + " AS patientuid " +
		"FROM " + ob.Source
	tmp := strings.Split(sql,",")
	sql = ""
	for _,t := range tmp {
		if strings.Contains(t,"string AS "){
			continue
		}
		sql += t + ","
	}
	sql = strings.TrimSuffix(sql,",") + ";"
	return sql
}

type ImageSql struct {
	Seriesuid string	//所属series id
	Imagename string	//影像文件名
	Source string		//
	Typex string		//影像类型
	Dpath string		//存储路径
	Imagenum string		//影像编号
}
//SELECT Imagefilename,ImageLocate,ImageType,ImageNumber FROM Image WHERE SeriesInstanceUID = '%s';
func AutoImages(ob ImageSql)string{
	if ob.Dpath == "string" || ob.Dpath =="" {
		return "dpath can not empty"
	}
	sql := "SELECT " + ob.Imagename + " AS imgname," +
		ob.Imagenum + " AS imgnum," + ob.Typex + " AS type," +
		ob.Dpath + " AS path " +
		"FROM " + ob.Source + " WHERE " + ob.Seriesuid + "='%s'"
	tmp := strings.Split(sql,",")
	sql = ""
	for _,t := range tmp {
		if strings.Contains(t,"string AS "){
			continue
		}
		sql += t + ","
	}
	sql = strings.TrimSuffix(sql,",") + ";"
	return sql
}

type ReportSql struct{
	Studyid string		//报告所属study id
	Source string		//
	Report string		//检查报告内容
	Reportlocate string	//检查报告存储地址
	Reportname string	//检查报告文件名
	Result string		//检查结果内容
	Resultlocate string	//检查结果存储路径
	Resultname string	//检查结果文件名
	Opinion string		//医生建议
	Description1 string	//检查描述
	Conclusion1 string	//检查结论
	Description2 string	//检查描述（BLOB格式）
	Conclusion2 string	//检查结论（BLOB格式）
	Info1 string		//检查信息
	Info2 string		//检查信息（BLOB格式）
	Part string		//报告存储路径
}

func AutoReport(ob ReportSql)string{
	if ob.Studyid == "string" || ob.Studyid =="" {
		return "Studyid can not empty"
	}
	sql := "SELECT " + ob.Report + " AS report," +
		ob.Reportlocate + " AS reportpath," + ob.Reportname + " AS reportname," +
		ob.Result + " AS result," + ob.Resultlocate + " AS resultpath," +
		ob.Resultname + " AS resultname," + ob.Opinion + " AS opinion," +
		ob.Conclusion1 + " AS conclusion1," + ob.Description1 + " AS description1," +
		ob.Conclusion2 + " AS conclusion2," + ob.Description2 + " AS description2," +
		ob.Info1 + " AS info1," + ob.Info2 + " AS info2," + ob.Part + " AS bodypart " +
		"FROM " + ob.Source + " WHERE " + ob.Studyid + "='%s'"
	tmp := strings.Split(sql,",")
	sql = ""
	for _,t := range tmp {
		if strings.Contains(t,"string AS "){
			continue
		}
		sql += t + ","
	}
	sql = strings.TrimSuffix(sql,",") + ";"
	return sql
}