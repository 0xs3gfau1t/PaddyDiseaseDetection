package storage

type Storage interface {
	AddFile(filename string, buffer []byte) error
	RemoveFile(filename string) error
}
