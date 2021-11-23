package model

type ProjectApiReq struct {
	Name string `v:"required#名字不能为空" json:"name"`
	Path string `v:"required#路径不能为空" json:"path"`
	Cmd  string `v:"required#命令不能为空" json:"cmd"`
}

type ProjectUpdateApiReq struct {
	Id   string
	Name string
	Path string
	Cmd  string
}
