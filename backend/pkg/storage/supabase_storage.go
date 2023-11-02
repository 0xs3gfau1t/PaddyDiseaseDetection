package storage

import (
	"bytes"
	"log"

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
		log.Fatalln("Failed to upload file:", err)
		return err
	}

	log.Println(res)

	return nil
}

func (fs *SupaBaseStorage) RemoveFile(filename string) error {
	res, err := fs.client.RemoveFile(fs.bucket, []string{filename})

	if err != nil {
		log.Fatalln("Failed to remove file:", err)
		return err
	}

	log.Println(res)

	return nil
}
