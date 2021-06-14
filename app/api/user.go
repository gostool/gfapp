package api

import (
	"gfapp/app/model"
	"gfapp/app/service"
	"gfapp/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 用户API管理对象
var User = new(userApi)
var UserAuth = new(userAuth)
var UserAuthWeb = new(userAuthWeb)

type userApi struct{}

type userAuth struct{}

type userAuthWeb struct{}

// @summary 用户登录接口
// @tags    用户服务
// @produce json
// @param   passport formData string true "用户账号"
// @param   password formData string true "用户密码"
// @router  /api/user/sign-in [POST]
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

// @summary 用户登录接口web
// @tags    用户服务
// @produce json
// @param   passport formData string true "用户账号"
// @param   password formData string true "用户密码"
// @router  /api/user/sign-in [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) SignInWeb(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiSignInWebReq
		serviceReq *model.UserServiceSignInWebReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	user, err := service.User.SignInWeb(r.Context(), serviceReq)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	token, err := service.Base.NewJwt(&model.ClaimServiceReq{
		Id:   gconv.String(user.Id),
		Salt: "",
	})
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", g.Map{
		"Id":    user.Id,
		"token": token,
	})
}

// @summary 用户注销/退出接口
// @tags    用户服务
// @produce json
// @router  /api/user/sign-out [GET]
// @success 200 {object} response.JsonResponse "执行结果, 1: 未登录"
func (a *userApi) SignOut(r *ghttp.Request) {
	if err := service.User.SignOut(r.Context()); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok")
}
