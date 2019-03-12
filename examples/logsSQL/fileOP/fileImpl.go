package fileOP

type FileOP interface {
	ReadFile(string) (bool)
	Init()
}
