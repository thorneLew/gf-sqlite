package service

import (
	"fmt"
	"project-build/app/dao"
	"project-build/app/model"
	"project-build/app/utils"
)

var Build = buildService{}

type buildService struct{}

func (s *buildService) Start(req *model.BuildApiReq) error {
	var proj *model.Project
	data, err := dao.Project.One(req)
	if err != nil {
		return err
	}

	data.Struct(&proj)
	fmt.Println(proj.Path)
	inWrokPath := "cd " + proj.Path + "&& pwd" + "&& yarn build"

	utils.Command(inWrokPath)
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

	return nil
}
