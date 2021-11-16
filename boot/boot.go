package boot

import (
	_ "project-build/packed"

	_ "github.com/logoove/sqlite"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
}
