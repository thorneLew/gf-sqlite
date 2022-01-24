package api

import (
	"fmt"
	"project-build/app/model"
	"project-build/app/service"
	"project-build/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/hpcloud/tail"
)

var Build = buildApi{}

type buildApi struct{}

// @summary 打包一个项目
// @tags   	打包模块
// @produce json
// @param   entity  body model.BuildApiReq true "创建请求"
// @router  /api/build [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (*buildApi) Index(r *ghttp.Request) {
	// ws, err := r.WebSocket()
	// if err != nil {
	// 	glog.Error(err)
	// 	r.Exit()
	// }

	// for {
	// 	msgType, msg, err := ws.ReadMessage()
	// 	if err != nil {
	// 		return
	// 	}
	// 	if err = ws.WriteMessage(msgType, msg); err != nil {
	// 		return
	// 	}
	// }

	var (
		data *model.BuildApiReq
	)

	if err := r.ParseForm(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if id, err := service.Build.Start(data); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "打包成功", g.Map{"id": id})
	}
}

// @summary 打包列表
// @tags   	打包模块
// @produce json
// @param   project_id query string false "工程ID"
// @router  /api/build/list [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (*buildApi) List(r *ghttp.Request) {
	project_id := r.Request.URL.Query().Get("project_id")

	if data, err := service.Build.List(project_id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", data)
	}
}

// @summary 读取打包详情
// @tags   	打包模块
// @produce json
// @param   project_id query string true "工程id"
// @router  /api/build/info [GET]
func (*buildApi) Info(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		glog.Error(err)
		r.Exit()
	}

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		project_id := string(msg)
		t, _ := tail.TailFile("./logs/build/"+project_id+".log", tail.Config{Follow: true})
		for line := range t.Lines {
			if err = ws.WriteMessage(msgType, []byte(line.Text)); err != nil {
				return
			}
		}
	}
}
