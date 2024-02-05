package usecase

import (
	"task-manager/model"
	"task-manager/repository"
)

type Usecase interface {
	GetAllTask() ([]model.TaskData, error)
	CreateTask(task model.CreateTaskData) (*model.TaskData, error)
}

type usecaseImpl struct {
	repository repository.Repository
}

func (u *usecaseImpl) GetAllTask() ([]model.TaskData, error) {
	return u.getAllTask()
}

func (u *usecaseImpl) CreateTask(task model.CreateTaskData) (*model.TaskData, error) {
	return u.createTask(task)
}

func NewUsecase(r repository.Repository) *usecaseImpl {
	return &usecaseImpl{
		repository: r,
	}
}
