package models

import (
	"zdy.local/fileOP"
	"zdy.local/cecetl/files"
)

type Data struct {
	Path string
	Dbname string
}

func Start(d Data)  {
	var obj fileOP.FileOP
	obj = new(files.Upload)
	fileOP.Run(d.Path,d.Dbname,obj)
}