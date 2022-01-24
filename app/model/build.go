package model

// 查询id
type BuildApiReq struct {
	Id string `v:"required" json:"id"`
}

type BuildListApiReq struct {
	ProjectId string `v:"required"`
}

type EntityBuildList struct {
	BuildLogs
	Project *Project `json:"project"`
}
