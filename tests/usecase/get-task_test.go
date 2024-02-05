package usecase_test

import (
	"task-manager/model"
	"task-manager/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTask(t *testing.T) {

	repo := &MockRepository{}
	usecase := usecase.NewUsecase(repo)

	result, err := usecase.GetAllTask()

	assert.NoError(t, err)

	// // Assert the expected result based on the mock data
	expectedResult := []model.TaskData{
		{Id: "1", Name: "Task 1", Time: "2024-02-03T12:00:00", Complete: false},
		{Id: "2", Name: "Task 2", Time: "2024-02-03T13:00:00", Complete: true},
	}

	assert.Equal(t, expectedResult, result)
}
