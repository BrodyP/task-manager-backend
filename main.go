package main

import (
	"io"
	"task-manager/handler"
	"task-manager/repository"
	"task-manager/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	re := repository.NewRepository()
	u := usecase.NewUsecase(re)
	h := handler.NewHandler(u)

	r.GET("/tasks", func(c *gin.Context) {
		result, err := h.GetAllTask()
		c.JSON(200, gin.H{
			"tasks": result,
			"error": err,
		})
	})

	r.POST("/create-task", func(c *gin.Context) {
		body := c.Request.Body
		defer body.Close()

		json, err := io.ReadAll(body)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		}
		result := h.CreateTask(json)

		if result.Error != nil {
			c.JSON(400, gin.H{
				"error": result.Error,
			})
		} else {
			c.JSON(200, result)
		}

	})

	r.Run()
}
