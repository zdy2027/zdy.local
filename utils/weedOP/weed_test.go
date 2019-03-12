package weedOP

import (
	"testing"
	"os"
)

var (
	filename = "../hello.txt"
)

func TestWeedUpload(t *testing.T)  {
	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	fid,size,err := Upload(filename,"text/plain",file,"192.168.2.104:9333")
	if err !=nil {
		t.Fatal(err)
	}else {
		t.Log(fid,size)
	}
}
