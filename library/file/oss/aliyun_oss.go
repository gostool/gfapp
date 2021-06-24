package oss

import (
	"errors"
	"mime/multipart"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliYunOSS struct {
	BucketUrl       string
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
}

func NewAliOSS(endpoint, accessKeyID, accessKeySecret, bucketName, bucketUrl string) (oss *AliYunOSS) {
	return &AliYunOSS{
		BucketUrl:       bucketUrl,
		Endpoint:        endpoint,
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		BucketName:      bucketName,
	}
}

func (t *AliYunOSS) NewBucket() (*oss.Bucket, error) {
	// 创建OSSClient实例。
	client, err := oss.New(t.Endpoint, t.AccessKeyID, t.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	return client.Bucket(t.BucketName)
}

func (t *AliYunOSS) GetBucket() (bucket *oss.Bucket, err error) {
	bucket, err = t.NewBucket()
	if err != nil {
		return nil, errors.New("function AliYunOSS.NewBucket() failed, err:" + err.Error())
	}
	return
}

func (t *AliYunOSS) UploadByFilePath(filename string, path string) (string, string, error) {
	// 获取存储空间。
	bucket, err := t.GetBucket()
	if err != nil {
		return "", "", errors.New("function AliYunOSS.NewBucket() failed, err:" + err.Error())
	}

	// 上传本地文件。
	filePath := NewFilePath(filename)
	err = bucket.PutObjectFromFile(filePath, path)
	if err != nil {
		return "", "", errors.New("function bucket.PutObjectFromFile failed, err:" + err.Error())
	}
	return t.BucketUrl + "/" + filePath, filePath, nil
}

func (t *AliYunOSS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	bucket, err := t.GetBucket()
	if err != nil {
		return "", "", errors.New("function AliYunOSS.NewBucket() failed, err:" + err.Error())
	}
	// 读取本地文件。
	fd, err := file.Open()
	if err != nil {
		return "", "", errors.New("function file.Open() failed, err:" + err.Error())
	}
	defer fd.Close()

	// 上传文件流。
	filePath := NewFilePath(file.Filename)
	err = bucket.PutObject(filePath, fd)
	if err != nil {
		return "", "", errors.New("function bucket.PutObject failed, err:" + err.Error())
	}
	return t.BucketUrl + "/" + filePath, filePath, nil
}

func (t *AliYunOSS) DeleteFile(key string) error {
	bucket, err := t.GetBucket()
	if err != nil {
		return err
	}
	err = bucket.DeleteObject(key)
	if err != nil {
		return errors.New("function bucket.DeleteObject() failed, err:" + err.Error())
	}
	return nil
}
