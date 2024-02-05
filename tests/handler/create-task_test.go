package handler_test

import (
	"encoding/json"
	"fmt"
	"task-manager/handler"
	"task-manager/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	var mockUseCase MockUsecase
	h := handler.NewHandler(mockUseCase)

	b := model.CreateTaskData{Name: mockTaskData.Name, Time: mockTaskData.Time, Complete: mockTaskData.Complete}

	body, _ := json.Marshal(b)
	result := h.CreateTask(body)

	assert.True(t, result.IsSuccess)
	assert.Nil(t, result.Error)

	fmt.Println(*result.TaskData)
	assert.Equal(t, mockTaskData, *result.TaskData)

}
