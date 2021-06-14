package api

import (
	"gfapp/app/model"
	"gfapp/app/service"
	"gfapp/library/response"

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
// @Router /api/tools/captcha [post]
func (t *toolsApi) Captcha(r *ghttp.Request) {
	result, err := service.Base.Captcha()
	//g.Log().Debugf("result:%v", result)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", result)
	}
}

// @Tags tools
// @Summary 生成token
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string response.JsonResponse
// @Router /api/tools/new-jwt [post]
func (t *toolsApi) NewJwt(r *ghttp.Request) {
	serviceReq := model.ClaimServiceReq{
		Id: "1",
	}
	token, err := service.Base.NewJwt(&serviceReq)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", token)
	}
}

func (t *toolsApi) ParseJwt(r *ghttp.Request) {
	serviceReq := model.ClaimServiceReq{
		Id:   "1",
		Salt: "",
	}
	tokenString := r.Header.Get("Token")
	claim, err := service.Base.ParseJwt(&serviceReq, tokenString)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", claim.Id, claim.StandardClaims.ExpiresAt)
}
