package minigo

import (
	"context"
	"os"

	"github.com/minio/minio-go/v7"
)

func (client *MinioClient) UploadFile(bucket string, uploadedPath string, filePath string, fileName string) (err error) {
	file, err := os.Open(filePath + "/" + fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	_, err = client.PutObject(context.Background(), bucket, uploadedPath+"/"+fileName, file, stat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return
	}

	return nil
}
