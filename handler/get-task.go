package handler

import "task-manager/model"

func (h taskHandlerImpl) getAllTask() ([]model.TaskData, error) {
	resp, err := h.usecase.GetAllTask()
	return resp, err
}
