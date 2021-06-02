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
		group.ALL("/hello", api.Hello)
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.CORS,
		)
		group.Group("/base", func(group *ghttp.RouterGroup) {
			group.ALLMap(g.Map{
				"/user":  api.User, // 用户
				"/chat":  api.Chat,
				"/tools": api.Tools,
			})
		})
		group.Group("/api", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.Auth)
			group.ALL("/user", api.UserAuth)
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
