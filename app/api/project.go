package api

import (
	"github.com/gogf/gf/net/ghttp"
)

var Project = projectApi{}

type projectApi struct{}

func (*projectApi) Index(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}
