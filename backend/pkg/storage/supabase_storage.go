package storage

import (
	"bytes"
	"fmt"
	"log"
	"time"

	sb_storage "github.com/supabase-community/storage-go"
)

type SupaBaseStorage struct {
	Client *sb_storage.Client
	Bucket string
}

func NewSupaBaseStorage(apiUrl string, token string, bucket string) *SupaBaseStorage {
	return &SupaBaseStorage{
		Client: sb_storage.NewClient(apiUrl, token, nil),
		Bucket: bucket,
	}
}

func (fs *SupaBaseStorage) AddFile(filename string, buffer []byte) error {
	reader := bytes.NewReader(buffer)
	res, err := fs.Client.UploadFile(fs.Bucket, filename, reader)

	if err != nil {
		return fmt.Errorf("Failed to upload file:%w", err)
	}

	log.Println(res)

	return nil
}

func (fs *SupaBaseStorage) RemoveFile(filename string) error {
	res, err := fs.Client.RemoveFile(fs.Bucket, []string{filename})

	if err != nil {
		return fmt.Errorf("Failed to remove file:%w", err)
	}

	log.Println(res)

	return nil
}

func (fs *SupaBaseStorage) GetFileByte(filename string) ([]byte, error) {
	return fs.Client.DownloadFile(fs.Bucket, filename)
}

func (fs *SupaBaseStorage) GetFilePath(filename string) (string, error) {
	res, err := fs.Client.CreateSignedUrl(fs.Bucket, filename, int(time.Hour*3/time.Second))
	if err != nil {
		return "", fmt.Errorf("Failed to create signed url:%w", err)
	}
	return res.SignedURL, nil
}
