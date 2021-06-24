package oss

import (
	"crypto/md5"
	"encoding/hex"
)

//@author: pilex
//@function: MD5V
//@description:md5加密
//@parm:str []byte
//@return: string

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
