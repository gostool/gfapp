package service

import (
	"gfapp/app/model"
	"gfapp/library/response"
	"net/http"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/net/ghttp"
)

// 中间件管理服务
var Middleware = middlewareService{}

type middlewareService struct{}

// 自定义上下文对象
func (s *middlewareService) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context.Init(r, customCtx)
	if user := Session.GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			Id:       user.Id,
			Passport: user.Passport,
			Name:     user.Name,
		}
	}
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// 允许接口跨域请求
func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// 鉴权中间件，只有登录成功之后才能通过
func (s *middlewareService) Auth(r *ghttp.Request) {
	if User.IsSignedIn(r.Context()) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

func (s *middlewareService) JwtAuth(r *ghttp.Request) {
	token := r.GetHeader("token")
	if token == "" {
		response.JsonExit(r, response.CODE_TOKEN, "token is not allow to be empty")
	}
	uid, err := Token.CheckToken(token)
	if err != nil {
		response.JsonExit(r, response.CODE_TOKEN, err.Error())
	}
	user, err := User.Find(gconv.Int64(uid))
	if err != nil {
		response.JsonExit(r, response.CODE_TOKEN, err.Error())
	} else {
		r.SetParam("uid", user.Id)
		r.SetParam("data", user.ToData())
		r.Middleware.Next()
	}
}

// 鉴权中间件，只有登录成功之后才能通过
func (s *middlewareService) AuthWeb(r *ghttp.Request) {
	if User.IsSignedInWeb(r.Header.Get("Token"), r.Header.Get("Id"), r.Header.Get("Salt")) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}
