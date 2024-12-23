package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"gitee.com/uniqptr/media-vault.git/internal/logging"
	"gitee.com/uniqptr/media-vault.git/internal/service"
)

type MediaControllerV1 struct {
	s *service.MediaService
}

func NewMediaControllerV1(s *service.MediaService) *MediaControllerV1 {
	return &MediaControllerV1{
		s: s,
	}
}

func (controller *MediaControllerV1) RegisterRoutes(router gin.IRouter) {
	g := router.Group("/v1")
	g.POST("/media/list", controller.MediaListV1)
	g.POST("/media/detail", controller.MediaDetailV1)
	g.POST("/media/add", controller.MediaAddV1)
	g.POST("/media/update", controller.MediaUpdateV1)
}

// MediaListV1 列出媒体
func (controller *MediaControllerV1) MediaListV1(c *gin.Context) {
	medias, err := controller.s.List(service.WithOrderBy("", true), service.WithPaginate(1, 10))
	if err != nil {
		logging.GetLogger().Error("failed to list media", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, ServerError(err))
		return
	}
	c.JSON(http.StatusOK, NewResponse(0, "OK", medias))
}

func (controller *MediaControllerV1) MediaDetailV1(c *gin.Context) {
	logging.GetLogger().Panic("not implemented")
}

func (controller *MediaControllerV1) MediaAddV1(c *gin.Context) {
	logging.GetLogger().Panic("not implemented")
}

func (controller *MediaControllerV1) MediaUpdateV1(c *gin.Context) {
	logging.GetLogger().Panic("not implemented")
}

func (controller *MediaControllerV1) MediaDeleteV1(c *gin.Context) {
	logging.GetLogger().Panic("not implemented")
}
