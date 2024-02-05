package repository_test

import (
	"encoding/json"
	"task-manager/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTask(t *testing.T) {
	mockingResponse := map[string]interface{}{
		"1": map[string]interface{}{
			"name":     "Task 1",
			"time":     "2024-02-03T12:00:00",
			"complete": false,
		},
		"2": map[string]interface{}{
			"name":     "Task 1",
			"time":     "2024-02-03T12:00:00",
			"complete": false,
		},
	}

	// convert mockingResponse to JSON bytes
	jsonBytes, _ := json.Marshal(mockingResponse)
	myFunction := func() {
		repo := repository.NewRepository()
		result, err := repo.GetAllTask()
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result, mockingResponse)
	}

	mockingServer(jsonBytes, myFunction)

}
