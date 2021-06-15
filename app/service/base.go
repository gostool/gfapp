package service

import (
	"gfapp/app/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 中间件管理服务
var Base = baseService{}

type baseService struct{}

func (b *baseService) NewJwt(r *model.ClaimServiceReq) (data string, err error) {
	claims := model.Claim{
		Id: r.Id,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "tt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	data, err = token.SignedString(r.SecSecret())
	if err != nil {
		return "", err
	}
	return data, nil
}

func (b *baseService) ParseJwt(r *model.ClaimServiceReq, tokenString string) (claim *model.Claim, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return r.SecSecret(), nil
	})

	if claim, ok := token.Claims.(*model.Claim); ok && token.Valid {
		return claim, nil
	}
	return nil, err
}
