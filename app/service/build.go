package service

import (
	"fmt"
	"project-build/app/dao"
	"project-build/app/model"
	"project-build/app/utils"
	"strconv"

	"github.com/gogf/gf/database/gdb"
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

	fmt.Print(proj.Id)
	currRow, _ := dao.BuildLogs.Insert(g.Map{
		"ProjectId": proj.Id,
		"status":    "0",
	})
	lastInsertId, _ := currRow.LastInsertId()

	strId := strconv.Itoa(int(lastInsertId))
	inWrokPath := "cd " + proj.Path + "&& pwd" + "&& yarn build"
	go utils.Command(inWrokPath, "build", strId+".log")
	return strId, nil
	// cmd := exec.Command("/bin/bash", "-c", inWrokPath)
	// stdin, _ := cmd.StdinPipe()
	// stdout, _ := cmd.StdoutPipe()
	// if err := cmd.Start(); err != nil {
	// 	fmt.Println("Execute failed when Start:" + err.Error())
	// }

	// stdin.Write([]byte("go text for grep\n"))
	// stdin.Write([]byte("go test text for grep\n"))
	// stdin.Close()

	// out_bytes, _ := ioutil.ReadAll(stdout)
	// stdout.Close()

	// if err := cmd.Wait(); err != nil {
	// 	fmt.Println("Execute failed when Wait:" + err.Error())
	// 	return nil
	// }

	// fmt.Println("Execute finished:" + string(out_bytes))
}

func (s *buildService) List(projectId string) (gdb.Result, error) {
	data, err := dao.BuildLogs.Where("project_id", projectId).All()
	return data, err
}
