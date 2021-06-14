package router

import (
	"gfapp/app/api"
	"gfapp/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.CORS,
		)
		group.Group("/api", func(group *ghttp.RouterGroup) {
			group.ALLMap(g.Map{
				"/user/register": api.UserRegister,
				"/chat":          api.Chat,
				"/tools":         api.Tools,
				"/user":          api.User,
			})
			group.Middleware(service.Middleware.Auth)
			group.ALL("/user", api.UserAuth)
		})
		group.Group("/api", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.AuthWeb)
			group.ALL("/userWeb", api.UserAuthWeb)
		})
	})
}

//func init() {
//	s := g.Server()
//	// 分组路由注册方式
//	s.Group("/", func(group *ghttp.RouterGroup) {
//		group.Middleware(
//			service.Middleware.Ctx,
//			service.Middleware.CORS,
//		)
//		group.ALL("/chat", api.Chat)
//		group.ALL("/user", api.User)
//		group.Group("/", func(group *ghttp.RouterGroup) {
//			group.Middleware(service.Middleware.Auth)
//			group.ALL("/user/profile", api.User.Profile)
//		})
//	})
//}
