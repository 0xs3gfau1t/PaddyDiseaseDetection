package storage

import (
	"bytes"
	"fmt"
	"log"
	"time"

	sb_storage "github.com/supabase-community/storage-go"
)

type SupaBaseStorage struct {
	client *sb_storage.Client
	bucket string
}

func NewSupaBaseStorage(apiUrl string, token string, bucket string) *SupaBaseStorage {
	return &SupaBaseStorage{
		client: sb_storage.NewClient(apiUrl, token, nil),
		bucket: bucket,
	}
}

func (fs *SupaBaseStorage) AddFile(filename string, buffer []byte) error {
	reader := bytes.NewReader(buffer)
	res, err := fs.client.UploadFile(fs.bucket, filename, reader)

	if err != nil {
		return fmt.Errorf("Failed to upload file:%w", err)
	}

	log.Println(res)

	return nil
}

func (fs *SupaBaseStorage) RemoveFile(filename string) error {
	res, err := fs.client.RemoveFile(fs.bucket, []string{filename})

	if err != nil {
		return fmt.Errorf("Failed to remove file:%w", err)
	}

	log.Println(res)

	return nil
}

func (fs *SupaBaseStorage) GetFileByte(filename string) ([]byte, error) {
	return fs.client.DownloadFile(fs.bucket, filename)
}

func (fs *SupaBaseStorage) GetFilePath(filename string) (string, error) {
	res, err := fs.client.CreateSignedUrl(fs.bucket, filename, int(time.Hour*3/time.Second))
	if err != nil {
		return "", fmt.Errorf("Failed to create signed url:%w", err)
	}
	return res.SignedURL, nil
}
