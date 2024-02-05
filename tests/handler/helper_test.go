package handler_test

import (
	"errors"
	"task-manager/model"
)

//Happy Case

type MockUsecase struct{}

var mockTaskData = model.TaskData{
	Id:       "1",
	Name:     "Mock Task",
	Time:     "00:00:00",
	Complete: false,
}

var errorMSG = "Mock error for testing"

func (m MockUsecase) GetAllTask() ([]model.TaskData, error) {
	return []model.TaskData{mockTaskData}, nil
}

func (m MockUsecase) CreateTask(task model.CreateTaskData) (*model.TaskData, error) {
	return &mockTaskData, nil
}

// Fail Case
type MockUsecaseWithError struct{}

func (m MockUsecaseWithError) GetAllTask() ([]model.TaskData, error) {
	// Return a mock error for testing error cases
	return nil, errors.New("Mock error for testing")
}

func (m MockUsecaseWithError) CreateTask(task model.CreateTaskData) (*model.TaskData, error) {
	return nil, errors.New("Mock error for testing")
}
