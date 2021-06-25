package api

import (
	"gfapp/app/model"
	"gfapp/app/service"
	"gfapp/library/response"

	"github.com/gogf/gf/crypto/gmd5"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 用户API管理对象
var User = new(userApi)
var UserAuth = new(userAuthApi)

type userApi struct{}

type userAuthApi struct{}

// @summary 用户登录接口
// @tags    用户服务
// @produce json
// @param   entity body model.UserApiLoginReq
// @router  /api/user/login [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) Login(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiLoginReq
		serviceReq *model.UserServiceLoginReq
		token      string
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, response.CODE_BAD, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, response.CODE_ERR, err.Error())
	}
	if !service.Base.Verify(apiReq.CaptchaId, apiReq.Captcha, true) {
		response.JsonExit(r, response.CODE_BAD, service.StoreError.Error())
	}
	//md5 password
	password, err := gmd5.EncryptString(serviceReq.Password)
	if err != nil {
		response.JsonExit(r, response.CODE_ERR, err.Error())
	}
	serviceReq.Password = password
	user, err := service.User.Login(serviceReq)
	if err != nil {
		response.JsonExit(r, response.CODE_BAD, err.Error())
	}
	token, err = service.Token.GenToken(gconv.String(user.Id), 0)
	if err != nil {
		response.JsonExit(r, response.CODE_BAD, err.Error())
	}
	user.Token = token
	response.JsonExit(r, response.CODE_OK, "ok", user)
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
	user, err := service.User.LoginWeb(r.Context(), serviceReq)
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
// @router  /api/user/logout [GET]
// @success 200 {object} response.JsonResponse "执行结果, 1: 未登录"
func (a *userAuthApi) Logout(r *ghttp.Request) {
	if err := service.User.LogOut(r.Context()); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok")
}

// Profile
// @summary 获取用户详情
// @tags    用户服务
// @produce json
// @router  /api/user/profile [GET]
// @success 200 {object} response.JsonResponse "用户详情"
func (a *userAuthApi) GetProfile(r *ghttp.Request) {
	data := r.GetParam("data")
	response.JsonExit(r, response.CODE_OK, "ok", data)
}
