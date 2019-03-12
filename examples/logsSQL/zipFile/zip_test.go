package zipFile

import (
	"os"
	"testing"
)

func TestCompress(t *testing.T) {
	f1, err := os.Open("1010531.3.12.2.1107.5.1.4.85646.30000015102700122848400001169.dic")
	if err != nil {
		t.Fatal(err)
	}
	defer f1.Close()
	f2, err := os.Open("1010531.3.12.2.1107.5.1.4.85646.30000015102700122848400001142.dic")
	if err != nil {
		t.Fatal(err)
	}
	defer f2.Close()
	var files = []*os.File{f1, f2}
	dest := "./test.zip"
	err = Compress(files, dest)
	if err != nil {
		t.Fatal(err)
	}
}
func TestDeCompress(t *testing.T) {
	err := DeCompress("../df_CR/1325421.3.12.2.1107.5.3.57.21079.11.201510291325420421.dic.PK", "../df_CR/")
	if err != nil {
		t.Fatal(err)
	}
}