package handler

import (
	"encoding/json"
	"task-manager/model"
)

func (h taskHandlerImpl) createTask(body []byte) model.CreateTaskResponse {
	var createTaskData model.CreateTaskData

	err := json.Unmarshal(body, &createTaskData)
	if err != nil {
		msg := err.Error()
		return model.CreateTaskResponse{
			IsSuccess: false,
			Error:     &msg,
		}
	}

	task, err := h.usecase.CreateTask(createTaskData)
	if err != nil {
		msg := err.Error()
		return model.CreateTaskResponse{
			IsSuccess: false,
			Error:     &msg,
		}
	}

	return model.CreateTaskResponse{
		IsSuccess: true,
		TaskData:  task,
	}

}
