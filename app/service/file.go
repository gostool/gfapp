package service

import (
	"fmt"

	"github.com/gogf/gf/net/ghttp"

	"github.com/gogf/gf/frame/g"
)

var File = fileService{}

func init() {
	v := g.Config().GetVar("upload")
	uploadConf := v.MapStrVar()
	File.Dir = uploadConf["dir"].String()
	File.Link = uploadConf["link"].String()
}

type fileService struct {
	Dir  string
	Link string
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
	filename = f.GetLink(filename)
	return filename, nil
}

//todo
func (f *fileService) SaveOss(file *ghttp.UploadFile) (filename string, err error) {
	return file.Save("/tmp")
}
