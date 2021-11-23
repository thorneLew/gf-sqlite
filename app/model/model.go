// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

import (
	"github.com/gogf/gf/os/gtime"
)

// Project is the golang structure for table project.
type Project struct {
	Id        int         `orm:"id"         json:"id"`        //
	Name      string      `orm:"name"       json:"name"`      //
	Path      string      `orm:"path"       json:"path"`      //
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` //
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` //
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deletedAt"` //
	Cmd       string      `orm:"cmd"        json:"cmd"`       //
}

// BuildLogs is the golang structure for table build_logs.
type BuildLogs struct {
	Id        int    `orm:"id"         json:"id"`        //
	Status    string `orm:"status"     json:"status"`    //
	CreatedAt string `orm:"created_at" json:"createdAt"` //
	UpdatedAt string `orm:"updated_at" json:"updatedAt"` //
	DeletedAt string `orm:"deleted_at" json:"deletedAt"` //
	ProjectId string `orm:"project_id" json:"projectId"` //
}
