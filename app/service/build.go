package service

import (
	"errors"
	"project-build/app/dao"
	"project-build/app/model"
	"project-build/app/utils"
	"strconv"
	"time"

	"github.com/gogf/gf/frame/g"
)

var Build = buildService{}

type buildService struct{}

func (s *buildService) Start(req *model.BuildApiReq) (string, error) {
	var proj *model.Project
	var buildLog *model.BuildLogs
	data, err := dao.Project.One(req)
	if err != nil {
		return "", err
	}
	buildLogRes, buildLogErr := dao.BuildLogs.OrderDesc("created_at").Where(g.Map{"project_id": req.Id}).One()

	if buildLogErr != nil {
		return "", buildLogErr
	}

	buildLogRes.Struct(&buildLog)

	formatTime := "2006-01-02 15:04:05"
	m, _ := time.ParseDuration("2m")
	nowTime, _ := time.Parse(formatTime, time.Now().Format(formatTime))
	createAt, _ := time.Parse(formatTime, buildLog.CreatedAt)
	createAt2Minute := createAt.Add(m)
	isOverTime := nowTime.Before(createAt2Minute)

	if isOverTime && buildLog.Status == "0" {
		return "", errors.New("此项目正在打包中！")
	}

	data.Struct(&proj)

	currRow, _ := dao.BuildLogs.Insert(g.Map{
		"ProjectId": proj.Id,
		"status":    0,
	})
	lastInsertId, _ := currRow.LastInsertId()

	strId := strconv.Itoa(int(lastInsertId))
	inWrokPath := "cd " + proj.Path + "&& pwd" + "&& " + proj.Cmd
	go utils.Command(inWrokPath, "build", strId, func() {
		dao.BuildLogs.Update(g.Map{"Status": 1}, "id", int(lastInsertId))
	})
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
