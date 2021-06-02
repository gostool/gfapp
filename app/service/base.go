package service

import (
	"gfapp/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/mojocn/base64Captcha"
)

// 中间件管理服务
var Base = baseService{}

type baseService struct{}

var Store base64Captcha.Store

func init() {
	Store = base64Captcha.DefaultMemStore
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 生成二维码的信息
func (b *baseService) Captcha() (result *response.Captcha, err error) {
	var info response.Captcha
	v := g.Config().GetVar("captcha")
	captchaConf := v.MapStrVar()
	l := captchaConf["key-long"].Int()
	w := captchaConf["image-width"].Int()
	h := captchaConf["image-height"].Int()
	var driver = base64Captcha.NewDriverDigit(h, w, l, 0.7, 80) // 字符,公式,验证码配置, 生成默认数字的driver
	var captcha = base64Captcha.NewCaptcha(driver, Store)
	info.Id, info.Path, err = captcha.Generate()
	return &info, err
}