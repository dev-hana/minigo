package minigo

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/minio/minio-go/v7"
)

func DownloadFile(client *minio.Client, bucket string, filePath string, fileName string, downloadPath string) (err error) {
	fmt.Println("client", client.EndpointURL())
	object, err := client.GetObject(context.Background(), bucket, filePath+"/"+fileName, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	downloaded, err := os.Create(downloadPath + "/" + fileName)
	if err != nil {
		return err
	}

	defer downloaded.Close()

	stat, err := object.Stat()
	if err != nil {
		return err
	}

	if _, err := io.CopyN(downloaded, object, stat.Size); err != nil {
		return err
	}
	return nil
}
