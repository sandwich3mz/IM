package utils

import (
	"IM/configs"
	"fmt"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
)

var MinioClientGlobal *MinioClient

type MinioClient struct {
	Client *minio.Client
}

func init() {
	minioConfig := configs.Conf.MinioConfig
	// 初始化 Minio 客户端
	minioClient, err := NewMinioClient(minioConfig.Endpoint, minioConfig.Access, minioConfig.Secret)
	if err != nil {
		logrus.Println("new minio client fail: ", err)
		return
	}
	MinioClientGlobal = minioClient
	logrus.Println("new minio client success")
}

// NewMinioClient 初始化minio
func NewMinioClient(endpoint, accessKey, secretKey string) (*MinioClient, error) {
	// 初始化 Minio 客户端
	minioClient, err := minio.New(endpoint, accessKey, secretKey, false)
	if err != nil {
		logrus.Println("new minio client fail: ", err)
		return nil, err
	}
	client := &MinioClient{
		Client: minioClient,
	}
	MinioClientGlobal = client
	return client, nil
}

// UploadFile 上传文件
func (m *MinioClient) UploadFile(bucketName string, objectName string, reader io.Reader, objectSize int64) error {
	_, err := m.Client.PutObject(bucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		logrus.Errorf("putObject fail: %v", err)
		return err
	}
	fmt.Println("Successfully uploaded", objectName)

	return nil
}

// DownloadFile 下载文件
func (m *MinioClient) DownloadFile(bucketName, objectName, filePath string) error {
	// 创建本地文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 下载存储桶中的文件到本地
	err = m.Client.FGetObject(bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	fmt.Println("Successfully downloaded", objectName)
	return nil
}

// DeleteFile 删除文件
func (m *MinioClient) DeleteFile(bucketName, objectName string) (bool, error) {
	// 删除存储桶中的文件
	err := m.Client.RemoveObject(bucketName, objectName)
	if err != nil {
		log.Println("remove object fail: ", err)
		return false, err
	}

	fmt.Println("Successfully deleted", objectName)
	return true, err
}

// ListObjects 列出文件
func (m *MinioClient) ListObjects(bucketName, prefix string) ([]string, error) {
	var objectNames []string

	for object := range m.Client.ListObjects(bucketName, prefix, true, nil) {
		if object.Err != nil {
			return nil, object.Err
		}

		objectNames = append(objectNames, object.Key)
	}

	return objectNames, nil
}

// GetUrl 返回对象的url地址，有效期时间为expires
func (m *MinioClient) GetUrl(bucketName string, objectName string, expires time.Duration) (string, error) {
	object, err := m.Client.PresignedGetObject(bucketName, objectName, expires, nil)
	if err != nil {
		log.Println("get object fail: ", err)
		return "", err
	}

	return object.String(), nil
}
