package model

// 查询id
type BuildApiReq struct {
	Id string `v:"required"`
}

type BuildListApiReq struct {
	ProjectId string `v:"required"`
}

type EntityBuildList struct {
	BuildLogs
	Project *Project `json:"project"`
}
