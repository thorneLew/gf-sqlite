package api

import (
	"project-build/app/dao"
	"project-build/app/model"
	"project-build/app/service"
	"project-build/library/response"

	"github.com/gogf/gf/net/ghttp"
)

var Project = projectApi{}

type projectApi struct{}

// @summary 获取项目列表
// @tags    项目管理
// @produce json
// @router  /api/project/list [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (*projectApi) List(r *ghttp.Request) {
	data, _ := dao.Project.All()
	response.JsonExit(r, 0, "ok", data)
}

// @summary 获取单个项目
// @tags    项目管理
// @produce json
// @param  id  query string true "项目ID"
// @router  /api/project/info [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (*projectApi) Info(r *ghttp.Request) {
	id := r.GetInt("id")
	data, _ := dao.Project.One("id", id)
	response.JsonExit(r, 0, "ok", data)
}

// @summary 创建项目
// @tags    项目管理
// @produce json
// @param   entity  body model.ProjectApiReq true "创建请求"
// @router  /api/project [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (*projectApi) Index(r *ghttp.Request) {
	var (
		data *model.ProjectApiReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Project.Create(data); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// @summary 更新项目信息
// @tags    项目管理
// @produce json
// @param   entity  body model.ProjectUpdateApiReq true "更新请求"
// @router  /api/project/update [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (*projectApi) Update(r *ghttp.Request) {
	var (
		data *model.ProjectUpdateApiReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Project.Update(data); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// @summary 删除项目
// @tags    项目管理
// @produce json
// @param   entity  body model.ProjectUpdateApiReq true "删除项目请求"
// @router  /api/project/delete [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (*projectApi) Delete(r *ghttp.Request) {
	var (
		data *model.ProjectUpdateApiReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Project.Delete(data); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}
