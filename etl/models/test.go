package models

import (
	"fmt"
	"zdy.local/cecetl/message"
)

func Test(input message.Data){
	fmt.Println(input.Dbname)
}