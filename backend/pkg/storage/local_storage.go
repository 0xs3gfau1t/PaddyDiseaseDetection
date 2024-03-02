package storage

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

type LocalStorage struct {
	RootFolder string
	ServerPath string
}

func NewLocalStorage(rootFilePath string, serverUrl string) (*LocalStorage, error) {
	folderPath := path.Dir(rootFilePath)

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Folder doesnot exists: %w", err)
	}

	return &LocalStorage{
		RootFolder: folderPath,
		ServerPath: serverUrl,
	}, nil
}

func (fs *LocalStorage) AddFile(filename string, buffer []byte) error {
	filepath := path.Join(fs.RootFolder, filename)
	if err := os.MkdirAll(path.Dir(filepath), 0770); err != nil {
		return err
	}

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
	filepath := path.Join(fs.RootFolder, filename)

	return os.Remove(filepath)
}

func (fs *LocalStorage) GetFileByte(filename string) ([]byte, error) {
	filepath := path.Join(fs.RootFolder, filename)

	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file:%w", err)
	}
	defer f.Close()

	return io.ReadAll(f)
}

func (fs *LocalStorage) GetFilePath(filename string) (string, error) {
	return fmt.Sprintf("%v/%v", fs.ServerPath, filename), nil
}
