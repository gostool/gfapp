package service

import (
	"github.com/gogf/gf/frame/g"

	"github.com/skip2/go-qrcode"
)

func (b *baseService) Qrcode(url string, size int) (result []byte, err error) {
	//最常见的二维码
	q, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return nil, err
	}
	v := g.Config().GetVar("qrcode")
	qrConf := v.MapStrVar()
	disableBorder := qrConf["disableBorder"].Bool()
	q.DisableBorder = disableBorder //去掉边框

	return q.PNG(size)
}
