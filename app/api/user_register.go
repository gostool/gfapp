package api

import (
	"gfapp/app/model"
	"gfapp/app/service"
	"gfapp/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 用户API管理对象
var UserRegister = new(userRegisterApi)

type userRegisterApi struct{}

// @summary 用户注册接口
// @tags    用户服务
// @produce json
// @param   entity  body model.UserRegisterApiSignUpReq true "注册请求"
// @router  /api/user/register/account [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userRegisterApi) Account(r *ghttp.Request) {
	var (
		apiReq     *model.UserRegisterApiSignUpReq
		serviceReq *model.UserRegisterServiceSignUpReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if !service.Base.Verify(apiReq.CaptchaId, apiReq.Captcha, true) {
		response.JsonExit(r, 1, service.StoreError.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.User.Register(serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// @summary 用户Email注册接口
// @tags    用户服务
// @produce json
// @param   entity  body model.UserRegisterApiMailSignUpReq true "注册请求"
// @router  /api/user/register/email [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userRegisterApi) Email(r *ghttp.Request) {
	var (
		apiReq     *model.UserRegisterApiMailSignUpReq
		serviceReq *model.UserRegisterServiceSignUpReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if !service.Base.Verify(apiReq.CaptchaId, apiReq.Captcha, true) {
		response.JsonExit(r, 1, service.StoreError.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.User.Register(serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// @summary 用户Phone注册接口
// @tags    用户服务
// @produce json
// @param   entity  body model.UserRegisterApiPhoneSignUpReq true "注册请求"
// @router  /api/user/register/phone [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userRegisterApi) Phone(r *ghttp.Request) {
	var (
		apiReq     *model.UserRegisterApiPhoneSignUpReq
		serviceReq *model.UserRegisterServiceSignUpReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if !service.Base.Verify(apiReq.CaptchaId, apiReq.Captcha, true) {
		response.JsonExit(r, 1, service.StoreError.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.User.Register(serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}
