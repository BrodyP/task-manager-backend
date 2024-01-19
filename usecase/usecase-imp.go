package usecase

import "task-manager/repository"

type Usecase interface {
	GetAllTask() ([]TaskData, error)
}

type usecaseImpl struct {
	repository repository.Repository
}

func (u *usecaseImpl) GetAllTask() ([]TaskData, error) {
	return u.getAllTask()
}

func NewUsecase(r repository.Repository) *usecaseImpl {
	return &usecaseImpl{
		repository: r,
	}
}
