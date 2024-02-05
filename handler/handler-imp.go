package handler

import (
	"task-manager/model"
	"task-manager/usecase"
)

type taskHandlerImpl struct {
	usecase usecase.Usecase
}

func (h taskHandlerImpl) GetAllTask() ([]model.TaskData, error) {
	return h.getAllTask()
}

func (h taskHandlerImpl) CreateTask(body []byte) model.CreateTaskResponse {
	return h.createTask(body)
}

func NewHandler(u usecase.Usecase) *taskHandlerImpl {
	return &taskHandlerImpl{
		usecase: u,
	}
}
