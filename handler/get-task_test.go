package handler

import (
	"errors"
	"reflect"
	"task-manager/usecase"
	"testing"
)

type MockTaskUsecase struct{}

func (m *MockTaskUsecase) GetAllTask() ([]usecase.TaskData, error) {
	// Return your desired mock data for testing
	return []usecase.TaskData{{Id: "1", Name: "Mock Task", Time: "00:00:00"}}, nil
}

type MockTaskUsecaseWithError struct{}

func (m *MockTaskUsecaseWithError) GetAllTask() ([]usecase.TaskData, error) {
	// Return a mock error for testing error cases
	return nil, errors.New("Mock error for testing")
}

func TestGetAllTask(t *testing.T) {
	// Create a new taskHandlerImpl with the mock implementation
	handler := &taskHandlerImpl{
		usecase: &MockTaskUsecase{},
	}

	// Call the getAllTask method and check the result
	result, err := handler.getAllTask()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	// Assert the result based on the mock data
	expected := []usecase.TaskData{{Id: "1", Name: "Mock Task", Time: "00:00:00"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestGetAllTask_ErrorCase(t *testing.T) {
	// Create a new taskHandlerImpl with a mock implementation that returns an error
	handler := &taskHandlerImpl{
		usecase: &MockTaskUsecaseWithError{},
	}

	// Call the getAllTask method and check for the expected error
	_, err := handler.getAllTask()
	if err == nil {
		t.Error("Expected an error, but got nil")
		return
	}

	// Optionally, you can check the specific error message or type
	expectedError := "Mock error for testing"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
	}
}
