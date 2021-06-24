package file

import (
	"gfapp/library/file/oss"
	"mime/multipart"
)

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

func NewOss(ossType string) OSS {
	switch ossType {
	case "aliYun":
		return &oss.AliYunOSS{}
	default:
		return nil
	}
}
