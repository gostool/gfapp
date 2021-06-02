package model

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

// 注册输入参数
type ClaimServiceReq struct {
	Id   string
	Salt string
}

func (c *ClaimServiceReq) SecSecret() []byte {
	// todo md5(uid.xxjkx.salt)
	mySigningKey := []byte("myBaseJWT")
	return mySigningKey
}
