package router

import (
	"project-build/app/api"

	"github.com/gogf/gf-jwt/example/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func middlewareAuth(r *ghttp.Request) {
	api.Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}

func init() {
	s := g.Server()
	// s.Group("/", func(group *ghttp.RouterGroup) {
	// 	group.ALL("/hello", api.Hello)
	// })
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.ALL("/login", api.Auth.LoginHandler)
		group.ALL("/refresh_token", api.Auth.RefreshHandler)
		group.ALL("/logout", api.Auth.LogoutHandler)
		group.Middleware(service.Middleware.CORS, middlewareAuth)
		group.ALL("/check", api.Check)
		group.ALL("/project", api.Project)
		group.ALL("/build", api.Build)
	})
}
