package api

import (
	"gfapp/app/service"
	"gfapp/library/response"
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
