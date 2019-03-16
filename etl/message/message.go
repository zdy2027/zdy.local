package message

import (
	"time"
)

type Element struct {
	Group       uint16		`json:"group"`
	Element     uint16		`json:"element"`
	Name        string		`json:"name"`
	Vr          string		`json:"vr"`
	Vl	    uint32		`json:"vl"`
	Value       []string		`json:"value,omitempty"`
}

type Patient struct {
	SickID string		`json:"IDcard"`
	SickName string		`json:"name"`
	Age int			`json:"age"`
	Sex string		`json:"sex"`
	SickNamec string	`json:"namec"`
	PatientID string	`json:"patientid"`
	PatientUID string	`json:"patientuid"`
	BirthDay time.Time	`json:"birthday"`
}

type Series struct {
	SeriesID string		`json:"seriesID"`
	Imgs []Image		`json:"imgs,omitempty"`
	Count int		`json:"count"`
	SeriesNum int		`json:"seriesnum"`
	DateTimex time.Time	`json:"datetimex,omitempty"`
	BodyPart string		`json:"bodypart"`
	Modality string		`json:"modality"`
}

type Image struct {
	Imgname string		`json:"imgname"`
	Urlx string		`json:"url,omitempty"`
	Size int64		`json:"size,omitempty"`
	Typex string		`json:"type,omitempty"`
	Statusx bool 		`json:"status,omitempty"`
	Dpath string		`json:"dpath"`
	Imagenum int		`json:"imagenum"`
}

type Study struct {
	StudyID string		`json:"studyInstanceUID"`
	SeriesInfo []Series	`json:"seriesinfo"`
	BodyPart string		`json:"bodypart"`
	Reportx Report		`json:"report"`
	Department string	`json:"department"`
	Clinical string		`json:"clinical"`
	Accessionnumber string	`json:"accessionnumber"`
	Studytime time.Time	`json:"studytime"`
	//Patientsource string
	Devicename string	`json:"devicename"`
}

type PrivateHits struct {
	Info []Private
}

type Private struct {
	SickInfo Patient	`json:"sickinfo"`
	Duns string		`json:"duns"`
	Orgname string		`json:"orgname"`
	Studys []Study		`json:"studys"`
	Source string		`json:"source"`
}

type Report struct {
	ReportID string 	`json:"reportid"`
	Report string		`json:"report"`
	ReportLocate string	`json:"reportlocate"`
	ReportUrl string	`json:"reporturl"`
	Result string		`json:"result"`
	ResultLocate string	`json:"resultlocate"`
	ResultUrl string	`json:"resulturl"`
	Opinion string		`json:"opinion"`
	Description1 string	`json:"description1"`
	Conclusion1 string	`json:"conclusion1"`
	Description2 string	`json:"description2"`
	Conclusion2 string	`json:"conclusion2"`
	Info1 string		`json:"info1"`
	Info2 string		`json:"info2"`
	BodyPart string		`json:"bodypart"`
}

type Data struct {
	Path	 string
	Dbname	string
}

type DCM struct {
	Pacs []Element		`json:"pacs"`
	SickID string		`json:"IDcard"`
	StudyID string		`json:"studyInstanceUID,omitempty"`
	SeriesID string		`json:"seriesID,omitempty"`
	Datetimex time.Time	`json:"datetimex,omitempty"`
}