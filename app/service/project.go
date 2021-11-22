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

func (s *projectService) Update(r *model.ProjectUpdateApiReq) error {
	_, err := dao.Project.Update(r, "id", r.Id)
	return err
}

func (s *projectService) Delete(r *model.ProjectUpdateApiReq) error {
	_, err := dao.Project.Delete("id", r.Id)
	return err
}
