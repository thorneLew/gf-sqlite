package api

import (
	"project-build/app/model"
	"project-build/app/service"
	"project-build/library/response"

	"github.com/gogf/gf/net/ghttp"
)

var Build = buildApi{}

type buildApi struct{}

// @summary 打包一个项目
// @tags   	打包模块
// @produce json
// @param   entity  body BuildApiReq true "创建请求"
// @router  /api/build [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (*buildApi) Index(r *ghttp.Request) {
	var (
		data *model.BuildApiReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.Build.Start(data); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "打包成功")
	}
}
