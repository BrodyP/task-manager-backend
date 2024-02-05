package repository

import "task-manager/model"

type Repository interface {
	GetAllTask() (map[string]interface{}, error)
	CreateTask(model.CreateTaskData) (string, error)
}

type repositoryImpl struct{}

func (r *repositoryImpl) GetAllTask() (map[string]interface{}, error) {
	return r.getAllTask()
}

func (r *repositoryImpl) CreateTask(task model.CreateTaskData) (string, error) {
	return r.createTask(task)
}

func NewRepository() Repository {
	return &repositoryImpl{}
}
