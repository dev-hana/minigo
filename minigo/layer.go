package minigo

import "github.com/minio/minio-go/v7"

type MinioClient struct {
	*minio.Client
}

type Minigo interface {
	DownloadFile(bucket string, filePath string, fileName string, downloadPath string) (err error)
	UploadFile(bucket string, uploadedPath string, filePath string, fileName string) (err error)
}
