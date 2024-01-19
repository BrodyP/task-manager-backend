package handler

import (
	"task-manager/usecase"

	"github.com/gin-gonic/gin"
)

type taskHandlerImpl struct {
	usecase usecase.Usecase
}

func (h *taskHandlerImpl) GetAllTask(c *gin.Context) ([]usecase.TaskData, error) {
	return h.getAllTask()
}

func NewHandler(u usecase.Usecase) *taskHandlerImpl {
	return &taskHandlerImpl{
		usecase: u,
	}
}
