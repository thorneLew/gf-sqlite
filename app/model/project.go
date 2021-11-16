package model

type ProjectApiReq struct {
	Name string `v:"required#名字不能为空"`
	Path string `v:"required#路径不能为空"`
}
