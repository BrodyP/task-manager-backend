package handler_test

import (
	"reflect"
	"task-manager/handler"
	"task-manager/model"
	"testing"
)

func TestGetAllTask(t *testing.T) {
	var mockUseCase MockUsecase

	handler := handler.NewHandler(&mockUseCase)

	result, err := handler.GetAllTask()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	expected := []model.TaskData{{Id: "1", Name: "Mock Task", Time: "00:00:00"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestGetAllTask_ErrorCase(t *testing.T) {
	var mockUseCase MockUsecaseWithError

	handler := handler.NewHandler(&mockUseCase)

	_, err := handler.GetAllTask()
	if err == nil {
		t.Error("Expected an error, but got nil")
		return
	}

	expectedError := "Mock error for testing"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
	}
}
