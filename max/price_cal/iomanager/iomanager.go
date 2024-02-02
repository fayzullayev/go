package iomanager

type IOManager interface {
	ReadLine() ([]string, error)
	WriteResult(data interface{}) error
}
