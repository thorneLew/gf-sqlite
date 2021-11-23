package service

import (
	"project-build/app/dao"
	"project-build/app/model"
	"project-build/app/utils"
	"strconv"

	"github.com/gogf/gf/frame/g"
)

var Build = buildService{}

type buildService struct{}

func (s *buildService) Start(req *model.BuildApiReq) (string, error) {
	var proj *model.Project
	data, err := dao.Project.One(req)
	if err != nil {
		return "", err
	}
	data.Struct(&proj)

	currRow, _ := dao.BuildLogs.Insert(g.Map{
		"ProjectId": proj.Id,
		"status":    "0",
	})
	lastInsertId, _ := currRow.LastInsertId()

	strId := strconv.Itoa(int(lastInsertId))
	inWrokPath := "cd " + proj.Path + "&& pwd" + "&& " + proj.Cmd
	go utils.Command(inWrokPath, "build", strId+".log")
	return strId, nil
}

func (s *buildService) List(projectId string) ([]model.EntityBuildList, error) {
	var buildLogs []model.EntityBuildList

	if len(projectId) != 0 {
		dao.BuildLogs.OrderDesc("id").Where("project_id", projectId).ScanList(&buildLogs, "BuildLogs")
	} else {
		dao.BuildLogs.OrderDesc("id").ScanList(&buildLogs, "BuildLogs")
	}

	dao.Project.ScanList(&buildLogs, "Project", "BuildLogs", "id:ProjectId")
	return buildLogs, nil
}
