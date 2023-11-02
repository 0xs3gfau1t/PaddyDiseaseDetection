package storage

import (
	"fmt"
	"log"
	"os"
	"path"
)

type LocalStorage struct {
	rootFolder string
}

func NewLocalStorage(rootFilePath string) (*LocalStorage, error) {
	folderPath := path.Dir(rootFilePath)

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Folder doesnot exists: %w", err)
	}

	return &LocalStorage{
		rootFolder: folderPath,
	}, nil
}

func (fs *LocalStorage) AddFile(filename string, buffer []byte) error {
	filepath := path.Join(fs.rootFolder, filename)

	file, err := os.Create(filepath)
	if err != nil {
		log.Fatalln("Error creating file:", err)
		return err
	}

	_, err = file.Write(buffer)
	if err != nil {
		log.Fatalln("Error writing to file:", err)
		return err
	}

	return nil

}

func (fs *LocalStorage) RemoveFile(filename string) error {
	filepath := path.Join(fs.rootFolder, filename)

	return os.Remove(filepath)
}
