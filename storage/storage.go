package storage

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage struct {
	minioClient *minio.Client
}

func NewStorage(address string, accessKeyID string, secretAccessKey string, useSSL bool) Storage {
	minioClient, _ := minio.New(address, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	storage := Storage{minioClient}
	return storage
}

func (storage Storage) Store(ctx context.Context, bucketName string, objectName string, filePath string, contentType string) {
	storage.minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
}

func (storage Storage) Get(ctx context.Context, bucketName string, objectName string, filePath string) {
	storage.minioClient.FGetObject(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
}
