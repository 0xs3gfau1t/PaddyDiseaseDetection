package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/h2non/filetype"
)

func ToMB(size int64) int64 {
	return size / (1024 * 1024 * 1024)
}
func GetFileExtension(name string) string {
	splitted := strings.Split(name, ".")
	return splitted[len(splitted)-1]
}
func IsFileAllowed(fileName string, file *bytes.Buffer) (string, error) {
	ext := GetFileExtension(fileName)
	if !filetype.IsImage(file.Bytes()) {
		return fileName, errors.New("Unsupported file type")
	}
	fileName = fmt.Sprintf("%v.%v", uuid.New(), ext)
	return fileName, nil
}
