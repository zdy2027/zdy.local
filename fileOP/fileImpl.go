package fileOP

import(
	"sync"
)

type FileOP interface {
	ReadFile(string,*sync.WaitGroup,chan int) (bool)
	Init(string)
	Close()
}