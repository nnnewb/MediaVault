package api

import (
	"errors"
	"io/fs"
	"net/http"

	"gitee.com/uniqptr/media-vault.git/internal/logging"
	"gitee.com/uniqptr/media-vault.git/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PathControllerV1 struct {
	s *service.PathService
}

func NewPathControllerV1(s *service.PathService) *PathControllerV1 {
	return &PathControllerV1{s: s}
}

func (controller *PathControllerV1) RegisterRoutes(router gin.IRouter) {
	g := router.Group("/v1")
	g.POST("/path/list", controller.PathList)
}

func (controller *PathControllerV1) PathList(c *gin.Context) {
	var req struct {
		Path string `json:"path"`
	}
	if err := c.BindJSON(&req); err != nil {
		logging.GetLogger().Error("failed to bind json", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, BadRequest(err))
		return
	}

	entries, err := controller.s.ReadDir(req.Path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if errors.Is(err, fs.ErrPermission) {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		logging.GetLogger().Error("failed to read dir", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, ServerError(err))
		return
	}

	c.JSON(http.StatusOK, OK(entries))
}
