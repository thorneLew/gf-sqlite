package api

import (
	"project-build/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Check = checkApi{}

type checkApi struct{}

func (*checkApi) IsLogoutHandler(r *ghttp.Request) {
	response.JsonExit(r, 200, "登录中", g.Map{"isLogin": true})
}
