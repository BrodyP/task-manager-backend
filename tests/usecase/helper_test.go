package usecase_test

import "task-manager/model"

type MockRepository struct{}

func (m *MockRepository) GetAllTask() (map[string]interface{}, error) {
	data := map[string]interface{}{
		"1": map[string]interface{}{
			"name":     "Task 1",
			"time":     "2024-02-03T12:00:00",
			"complete": false,
		},
		"2": map[string]interface{}{
			"name":     "Task 2",
			"time":     "2024-02-03T13:00:00",
			"complete": true,
		},
	}

	return data, nil
}

func (m *MockRepository) CreateTask(task model.CreateTaskData) (string, error) {
	return "xyz1234", nil
}
