package model

type ProjectApiReq struct {
	Name string `v:"required#名字不能为空" json:"name"`
	Path string `v:"required#路径不能为空" json:"path"`
}

type ProjectUpdateApiReq struct {
	Id   string
	Name string
	Path string
}
