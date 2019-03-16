package models

import (
	"fmt"
	"os/exec"
	"github.com/astaxie/beego/config"
)

var iniconf, _ = config.NewConfig("ini", "conf/app.conf")

func UpdateCycle() error {
	fmt.Println("tk1");
	
	return nil
}

func UploadCycle() error {
	if iniconf.String("os")=="win"{
		exec.Command("CECClient.exe")
	}else {
		exec.Command("./CECClient.exe")
	}
	return nil
}