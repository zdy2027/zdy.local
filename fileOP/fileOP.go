package fileOP

import (
	"path/filepath"
	"os"
	"io"
	"io/ioutil"
	"sync"
	"zdy.local/utils/nlogs"
	"github.com/astaxie/beego"
)

var(
	wg sync.WaitGroup
	goroutine chan int
)

func init(){
	goroutine_num,_ := beego.AppConfig.Int("golang::goroutine")
	goroutine = make(chan int,goroutine_num)
}

func Run(input string,comx string,obj FileOP){
	//输入文件可以是路径列表，此处遍历所有输入路径列表
	obj.Init(comx)
	defer obj.Close()
	fileInfo, err := os.Stat(input)
	if err != nil {
		nlogs.FileLogs.Error("INPUT PATH ERROR",err.Error())
		return
	}
	if fileInfo.IsDir() {				//判断该路径是否是文件夹
		walkDir(input,obj)				//循环遍历文件夹下所有子文件
	}else {
		wg.Add(1)
		goroutine <- 1
		obj.ReadFile(input,&wg,goroutine)
	}
	wg.Wait()
}

//循环遍历子文件夹
func walkDir(dir string,obj FileOP){
	filepath.Walk(dir,func(path string, info os.FileInfo, err error)error{
		if err != nil {
			return err
		}
		if filepath.Ext(info.Name())==".DS_Store"{
			return nil
		}
		if info.IsDir() {
			return nil				//如果是子文件夹继续遍历
		}else {
			//readFile(path)				//如果是文件开始读取文件
			wg.Add(1)
			goroutine <- 1
			go obj.ReadFile(path,&wg,goroutine)
		}
		return nil
	})
}

func GetPath(dirPath string)([]string){
	var result []string
	filepath.Walk(dirPath,func(path string, info os.FileInfo, err error)error{
		if err != nil {
			return err
		}
		if info.IsDir() {
			result = append(result,path)
			files, err := ioutil.ReadDir(path)//读取目录下文件
		    if err != nil{
		        return nil
		    }
		    if len(files)>30{
		    	return filepath.SkipDir
		    }
			return nil				//如果是子文件夹继续遍历
		}
		return nil
	})
	return result
}

func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		nlogs.ConsoleLogs.Error("src file error",err.Error())
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		nlogs.ConsoleLogs.Error("dest file error",err.Error())
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func PathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}