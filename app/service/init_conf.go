package service

import (
	"errors"
	"github.com/mojocn/base64Captcha"
)

// captcha service conf
var Store base64Captcha.Store
var StoreError error

func init() {
	Store = base64Captcha.DefaultMemStore
	StoreError = errors.New("verify error")
}
