package service

import (
	"errors"

	"github.com/gogf/gf/os/glog"

	"github.com/gogf/gf/frame/g"
	"github.com/mojocn/base64Captcha"
)

var logger *glog.Logger

// captcha service conf
var Store base64Captcha.Store
var StoreError error

//jwt
var jwtSalt string
var jwtExp int64

func init() {
	logger = g.Log("debug")

	Store = base64Captcha.DefaultMemStore
	StoreError = errors.New("verify error")

	//jwt
	v := g.Config().GetVar("jwt")
	jwtConf := v.MapStrVar()
	jwtSalt = jwtConf["salt"].String()
	jwtExp = jwtConf["exp"].Int64()
}
