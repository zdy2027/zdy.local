package cecdcm

import (
	"io/ioutil"
	"sync"
	"zdy.local/cecetl/message"
	"zdy.local/utils/nlogs"
	"github.com/gillesdemey/go-dicom"
	"fmt"
	"strings"
	"encoding/hex"
	"bytes"
	"encoding/binary"
)

func DCMExtract(filename string) []message.Element {
	bytes, err := ioutil.ReadFile(filename)
	if err!=nil {
		nlogs.ConsoleLogs.Debug("open file error")
		panic(err)
	}
	parser, _ := dicom.NewParser()
	datax, c := parser.Parse(bytes)
	dcm := &dicom.DicomFile{}
	gw := new(sync.WaitGroup)
	dcm.Discard(c,gw)
	gw.Wait()
	return formatElem(datax.Elements)
}

func BytesToInt(buf []uint8) (result []int) {
	for _,i := range buf {
		result = append(result,int(i))
	}
	return result
}

func U2S(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return
}

func formatElem(elems []dicom.DicomElement) []message.Element {
	var Pacs_info []message.Element
	for _,elem := range elems {
		var tmp = message.Element{}
		tmp.Element = elem.Element
		tmp.Group = elem.Group
		tmp.Name = elem.Name
		tmp.Vr = elem.Vr
		tmp.Vl = elem.Vl
		//tmp.Value = strings.Replace(fmt.Sprint(elem.Value),`\`,`\\`,-1)
		if len(elem.Value) != 0 {
			for _,v := range elem.Value{
				tmp.Value = append(tmp.Value,strings.Replace(fmt.Sprint(v),`\`,`\\`,-1))
			}
			/*
			switch elem.Value[0].(type) {
			case []uint8:
				tmp.Value = BytesToInt(elem.Value[0].([]uint8))
			case string:
				if elem.Name == "Private Data"{
					//tmp.Value = strings.Replace(fmt.Sprint(elem.Value[0]),`\`,`\\`,-1)
					continue
					//elem.Value[0] = Ucode2UTF(elem.Value[0].(string))
					//nlogs.ConsoleLogs.Debug(elem.Value[0].(string))
				}
				tmp.Value = elem.Value
			default:
				tmp.Value = elem.Value
			}*/
		}
		//nlogs.ConsoleLogs.Debug("value is ",tmp.Value)
		Pacs_info = append(Pacs_info,tmp)
	}
	return Pacs_info
}