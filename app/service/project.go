package service

import (
	"project-build/app/dao"
	"project-build/app/model"
)

var Project = projectService{}

type projectService struct{}

func (s *projectService) Create(r *model.ProjectApiReq) error {
	if _, err := dao.Project.Insert(r); err != nil {
		return err
	}
	return nil
}
