package main

import (
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
		result, err := h.GetAllTask(c)
		c.JSON(200, gin.H{
			"tasks": result,
			"error": err,
		})
	})

	r.Run()
}
