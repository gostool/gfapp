package api

import (
	"gfapp/app/model"
	"gfapp/app/service"
	"gfapp/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 用户API管理对象
var User = new(userApi)
var UserAuth = new(userAuth)

type userApi struct{}

type userAuth struct{}

// @summary 用户注册接口
// @tags    用户服务
// @produce json
// @param   entity  body model.UserApiSignUpReq true "注册请求"
// @router  /base/user/register [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) Register(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiSignUpReq
		serviceReq *model.UserServiceSignUpReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
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

// @summary 用户登录接口
// @tags    用户服务
// @produce json
// @param   passport formData string true "用户账号"
// @param   password formData string true "用户密码"
// @router  /base/user/sign-in [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) SignIn(r *ghttp.Request) {
	var (
		data *model.UserApiSignInReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.User.SignIn(r.Context(), data.Passport, data.Password); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// @summary 判断用户是否已经登录
// @tags    用户服务
// @produce json
// @router  /base/user/is-signed-in [GET]
// @success 200 {object} response.JsonResponse "执行结果:`true/false`"
func (a *userApi) IsSignedIn(r *ghttp.Request) {
	response.JsonExit(r, 0, "", service.User.IsSignedIn(r.Context()))
}

// @summary 用户注销/退出接口
// @tags    用户服务
// @produce json
// @router  /base/user/sign-out [GET]
// @success 200 {object} response.JsonResponse "执行结果, 1: 未登录"
func (a *userApi) SignOut(r *ghttp.Request) {
	if err := service.User.SignOut(r.Context()); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok")
}

// @summary 检测用户账号接口(唯一性校验)
// @tags    用户服务
// @produce json
// @param   passport query string true "用户账号"
// @router  /base/user/check-passport [GET]
// @success 200 {object} response.JsonResponse "执行结果:`true/false`"
func (a *userApi) CheckPassport(r *ghttp.Request) {
	var (
		data *model.UserApiCheckPassportReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if data.Passport != "" && !service.User.CheckPassport(data.Passport) {
		response.JsonExit(r, 1, "账号已经存在", false)
	}
	response.JsonExit(r, 0, "", true)
}

// @summary 检测用户昵称接口(唯一性校验)
// @tags    用户服务
// @produce json
// @param   nickname query string true "用户昵称"
// @router  /base/user/check-nick-name [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) CheckNickName(r *ghttp.Request) {
	var (
		data *model.UserApiCheckNickNameReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if data.Nickname != "" && !service.User.CheckNickName(data.Nickname) {
		response.JsonExit(r, 1, "昵称已经存在", false)
	}
	response.JsonExit(r, 0, "ok", true)
}

// @summary 获取用户详情信息
// @tags    用户服务
// @produce json
// @router  /api/user/get-profile [GET]
// @success 200 {object} model.User "用户信息"
func (a *userAuth) GetProfile(r *ghttp.Request) {
	response.JsonExit(r, 0, "", service.User.GetProfile(r.Context()))
}