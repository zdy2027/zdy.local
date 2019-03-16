package cecdcm

import (
	"testing"
	"encoding/json"
	"fmt"
)

func Test_DCMExtract(t *testing.T)  {
	Pacs_test := DCMExtract("test.DCM")
	testx,_ := json.Marshal(Pacs_test)
	fmt.Println(string(testx))
	//DCMExtract("test.DCM")
}