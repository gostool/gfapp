package api

import (
	"gfapp/app/service"
	"gfapp/library/response"
	"time"

	//"github.com/gogf/gf-jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Tools = new(toolsApi)

type toolsApi struct{}

// @Tags tools
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/tools/captcha [post]
func (t *toolsApi) Captcha(r *ghttp.Request) {
	result, err := service.Base.Captcha()
	g.Log().Debugf("result:%v", result)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", result)
	}
}

func (t *toolsApi) NewJwt(r *ghttp.Request) {
	mySigningKey := []byte("myBaseJWT")
	type MyCustomClaims struct {
		Id string `json:"1"`
		jwt.StandardClaims
	}
	// Create the Claims
	claims := MyCustomClaims{
		"1",
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "tt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	r.Response.Writef("%v %v", ss, err)
}

func (t *toolsApi) ParseJwt(r *ghttp.Request) {
	type MyCustomClaims struct {
		Id string `json:"1"`
		jwt.StandardClaims
	}
	tokenString := r.Header.Get("xToken")
	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("myBaseJWT"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		r.Response.Writef("%v %v", claims.Id, claims.StandardClaims.ExpiresAt)
	} else {
		r.Response.Writeln(err)
	}
}
