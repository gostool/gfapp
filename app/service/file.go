package service

import (
	"fmt"
	"gfapp/library/file/oss"

	"github.com/gogf/gf/net/ghttp"

	"github.com/gogf/gf/frame/g"
)

var File = fileService{}
var aliyun = oss.AliYunOSS{}

func init() {
	v := g.Config().GetVar("upload")
	uploadConf := v.MapStrVar()
	File.Dir = uploadConf["dir"].String()
	File.Link = uploadConf["link"].String()
	File.FileServerEnabled = uploadConf["fileServerEnabled"].Bool()

	v = g.Config().GetVar("aliyun")
	aliyunConf := v.MapStrVar()
	aliyun.BucketUrl = aliyunConf["bucketurl"].String()
	aliyun.Endpoint = aliyunConf["endpoint"].String()
	aliyun.AccessKeyID = aliyunConf["accessKeyID"].String()
	aliyun.AccessKeySecret = aliyunConf["accessKeySecret"].String()
	aliyun.BucketName = aliyunConf["bucketName"].String()
}

type fileService struct {
	Dir               string
	Link              string
	FileServerEnabled bool
}

func (f *fileService) GetLink(filename string) string {
	return fmt.Sprintf("%v/%v", File.Link, filename)
}

//service test api
func (f *fileService) Save(file *ghttp.UploadFile) (filename string, err error) {
	filename, err = file.Save(File.Dir)
	if err != nil {
		return "", err
	}
	if File.FileServerEnabled {
		filename = f.GetLink(filename)
	}
	return filename, nil
}

//todo
func (f *fileService) SaveOss(file *ghttp.UploadFile) (fileUrl string, filePath string, err error) {
	fileUrl, filePath, err = aliyun.UploadFile(file.FileHeader)
	if err != nil {
		return "", "", err
	}
	return fileUrl, filePath, nil
}
