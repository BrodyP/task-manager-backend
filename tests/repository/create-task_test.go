package repository_test

import (
	"encoding/json"
	"task-manager/model"
	"task-manager/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"task-manager/model"
// 	"testing"
// )

func TestCreateTask(t *testing.T) {
	repo := repository.NewRepository()
	mockId := "someId"

	mockData := map[string]interface{}{
		"name": mockId,
	}

	myFunction := func() {

		result, err := repo.CreateTask(model.CreateTaskData{
			Name:     "Test Task",
			Time:     "2020-01-01T00:00:00Z",
			Complete: false,
		})
		assert.NoError(t, err)
		assert.NotEqual(t, result, "")
		assert.Equal(t, result, mockId)
	}
	jsonBytes, _ := json.Marshal(mockData)
	mockingServer(jsonBytes, myFunction)

}
