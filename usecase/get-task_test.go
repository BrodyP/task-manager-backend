package usecase

import (
	"errors"
	"reflect"
	"testing"
)

type repositoryMock struct{}

func (m *repositoryMock) GetAllTask() (map[string]interface{}, error) {
	mockResponse := map[string]interface{}{
		"task1": map[string]interface{}{
			"id":   "task1",
			"name": "Task One",
			"time": "10:00 AM",
		},
		"task2": map[string]interface{}{
			"id":   "task2",
			"name": "Task Two",
			"time": "2:00 PM",
		},
	}
	return mockResponse, nil
}

type repositoryMockError struct{}

func (m *repositoryMockError) GetAllTask() (map[string]interface{}, error) {
	return nil, errors.New("repository error")
}

func TestGetAllTask(t *testing.T) {
	usecase := &usecaseImpl{repository: &repositoryMock{}}

	// Call the getAllTask method
	result, err := usecase.getAllTask()

	// Check if there is an error
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the result is not nil
	if result == nil {
		t.Fatal("Expected non-nil result, got nil")
	}

	// Check if the result matches the expected output
	expectedResult := []TaskData{
		{Id: "task1", Name: "Task One", Time: "10:00 AM"},
		{Id: "task2", Name: "Task Two", Time: "2:00 PM"},
	}

	for _, task := range expectedResult {
		//find task in result
		var found bool
		for _, taskResult := range result {
			if task.Id == taskResult.Id && task.Name == taskResult.Name && task.Time == taskResult.Time {
				found = true
			}
		}
		if !found {
			t.Fatalf("Expected result %v, got %v", expectedResult, result)
		}
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Expected result %v, got %v", expectedResult, result)
	}
}

func TestGetAllTask_ErrorInRepository(t *testing.T) {
	// Create an instance of usecaseImpl with the mock repository
	usecase := &usecaseImpl{repository: &repositoryMockError{}}

	// Call the getAllTask method
	result, err := usecase.getAllTask()

	// Check if the error is as expected
	if err == nil || err.Error() != "repository error" {
		t.Fatalf("Expected 'repository error', got %v", err)
	}

	// Check if the result is nil
	if len(result) != 0 {
		t.Fatal("Expected nil result, got non-nil result")
	}
}
