package handler

import (
	"task-manager/usecase"
)

func (h *taskHandlerImpl) getAllTask() ([]usecase.TaskData, error) {
	resp, err := h.usecase.GetAllTask()
	return resp, err
}
