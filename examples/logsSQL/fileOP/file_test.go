package fileOP

import (
	"testing"
)

var (
	filename = "../hello.txt"
)

func TestFile(t *testing.T){
	err := readFile(filename)
	if err != true {
		t.Fatal(err)
	}
	t.Log("readfile success")
}
