package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nnnewb/media-vault/internal/logging"
	"github.com/nnnewb/media-vault/internal/models"
	"github.com/nnnewb/media-vault/internal/service"
	"go.uber.org/zap"
)

type TaskControllerV1 struct {
	taskService *service.TaskService
}

func NewTaskControllerV1(taskService *service.TaskService) *TaskControllerV1 {
	return &TaskControllerV1{taskService: taskService}
}

func (controller *TaskControllerV1) RegisterRoutes(router gin.IRouter) {
	g := router.Group("/v1")
	g.POST("/task/list", controller.TaskListV1)
}

func (controller *TaskControllerV1) TaskListV1(c *gin.Context) {
	var req struct {
		Pagination
	}
	if err := c.BindJSON(&req); err != nil {
		logging.GetLogger().Error("failed to bind json", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, BadRequest(err))
		return
	}

	options := []service.QueryOption{
		req.Pagination.WithDefault().QueryOption(),
	}
	tasks, err := controller.taskService.List(options...)
	if err != nil {
		logging.GetLogger().Error("failed to list tasks", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, ServerError(err))
		return
	}

	data := make([]models.TaskDTO, len(tasks))
	for i := range tasks {
		tasks[i].AsDTO(&data[i])
	}
	c.JSON(http.StatusOK, OK(data))
}
