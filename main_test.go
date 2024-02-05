package main_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"task-manager/handler"
	"task-manager/model"
	"task-manager/repository"
	"task-manager/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// Create a new Gin router
	r := gin.Default()

	// Create a new repository, usecase, and handler
	re := repository.NewRepository()
	u := usecase.NewUsecase(re)
	h := handler.NewHandler(u)

	// Set up the routes
	r.GET("/tasks", func(c *gin.Context) {
		result, err := h.GetAllTask()
		c.JSON(200, gin.H{
			"tasks": result,
			"error": err,
		})
	})

	r.POST("/create-task", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		defer c.Request.Body.Close()

		result := h.CreateTask(body)

		if result.Error != nil {
			c.JSON(400, gin.H{
				"error": result.Error,
			})
		} else {
			c.JSON(200, result)
		}
	})

	// Create a test request for GET /tasks
	getAllTaskReq, _ := http.NewRequest("GET", "/tasks", nil)
	getAllTaskResp := httptest.NewRecorder()
	r.ServeHTTP(getAllTaskResp, getAllTaskReq)

	// Assert the response status code and body
	assert.Equal(t, http.StatusOK, getAllTaskResp.Code)

	var getAllTaskRespBody struct {
		Tasks []model.TaskData `json:"tasks"`
		Error error            `json:"error"`
	}
	json.Unmarshal(getAllTaskResp.Body.Bytes(), &getAllTaskRespBody)

	// Assert the response body
	expectedGetAllTaskRespBody := struct {
		Tasks []model.TaskData `json:"tasks"`
		Error error            `json:"error"`
	}{
		Tasks: []model.TaskData{{Id: "1", Name: "Mock Task", Time: "00:00:00"}},
		Error: nil,
	}
	assert.Equal(t, expectedGetAllTaskRespBody, getAllTaskRespBody)

	// Create a test request for POST /create-task
	createTaskReqBody := []byte(`{"name": "New Task", "time": "01:23:45"}`)
	createTaskReq, _ := http.NewRequest("POST", "/create-task", bytes.NewBuffer(createTaskReqBody))
	createTaskReq.Header.Set("Content-Type", "application/json")
	createTaskResp := httptest.NewRecorder()
	r.ServeHTTP(createTaskResp, createTaskReq)

	// Assert the response status code and body
	assert.Equal(t, http.StatusOK, createTaskResp.Code)

	var createTaskRespBody struct {
		Error error `json:"error"`
	}
	json.Unmarshal(createTaskResp.Body.Bytes(), &createTaskRespBody)

	// Assert the response body
	expectedCreateTaskRespBody := struct {
		Error error `json:"error"`
	}{
		Error: nil,
	}
	assert.Equal(t, expectedCreateTaskRespBody, createTaskRespBody)
}
