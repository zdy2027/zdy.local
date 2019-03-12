package pacs

import (
	"io/ioutil"
	"sync"
	"zdy.local/utils/nlogs"
	"github.com/gillesdemey/go-dicom"
)

type DCMExtract struct {
	Sfzh string		`json:"sfzh"`
	Type string		`json:"type"`
	Orgname string		`json:"orgname"`
	Rawname string		`json:"rawname"`
	Duns string		`json:"duns"`
	Created_date string	`json:"created_date"`
	Created_by string	`json:"created_by"`
	Urlx string		`json:"urlx"`
	Pacs_info []Element	`json:"info"`
}

type Element struct {
	Group       uint16		`json:"group"`
	Element     uint16		`json:"element"`
	Name        string		`json:"name"`
	Vr          string		`json:"vr"`
	Value       interface{}		`json:"value,omitempty"`
}

func (d *DCMExtract) Init() {

}

func (d *DCMExtract) ReadFile(filename string) bool {
	bytes, err := ioutil.ReadFile(filename)
	if err!=nil {
		nlogs.ConsoleLogs.Error(err.Error())
		nlogs.FileLogs.Error(err.Error())
		return false
	}

	parser, _ := dicom.NewParser()
	datax, c := parser.Parse(bytes)

	dcm := &dicom.DicomFile{}
	//ppln := dcm.Parse(bytes)
	gw := new(sync.WaitGroup)
	//ppln = dcm.WriteImagesToFolder(ppln, gw, "floder/")
	dcm.Discard(c,gw)
	gw.Wait()
	d.formatElem(datax.Elements)

	return true
}

func (d *DCMExtract) Close()  {
	
}

func BytesToInt(buf []uint8) (result []int) {
	for _,i := range buf {
		result = append(result,int(i))
	}
	return result
}

func (d *DCMExtract) formatElem(elems []dicom.DicomElement)  {
	for _,elem := range elems {
		var tmp = Element{}
		tmp.Element = elem.Element
		tmp.Group = elem.Group
		tmp.Name = elem.Name
		tmp.Vr = elem.Vr
		if len(elem.Value) != 0 {
			switch elem.Value[0].(type) {
			case []uint8:
				tmp.Value = BytesToInt(elem.Value[0].([]uint8))
			default:
				tmp.Value = elem.Value
			}
		}
		d.Pacs_info = append(d.Pacs_info,tmp)
	}
}