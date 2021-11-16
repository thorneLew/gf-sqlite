package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Hello = helloApi{}

type helloApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Index(r *ghttp.Request) {
	db := g.DB()
	// db.Model("user").Delete(1)
	db.Model("user").Data(g.Map{"name": "小丽1", "age": "18"}).Insert()
	r.Response.Writeln("Hello World!")
}
