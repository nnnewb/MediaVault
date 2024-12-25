package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"gitee.com/uniqptr/media-vault.git/internal/logging"
	"gitee.com/uniqptr/media-vault.git/internal/models"
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
	g.GET("/media/cover/:id", controller.MediaCoverDownloadV1)
}

// MediaListV1 列出媒体
func (controller *MediaControllerV1) MediaListV1(c *gin.Context) {
	var req struct {
		Pagination
		OrderBy
	}
	if err := c.BindJSON(&req); err != nil {
		logging.GetLogger().Error("failed to bind json", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, BadRequest(err))
		return
	}

	options := []service.QueryOption{
		req.OrderBy.WithDefault().QueryOption(),
		req.Pagination.WithDefault().QueryOption(),
	}

	medias, err := controller.s.List(options...)
	if err != nil {
		logging.GetLogger().Error("failed to list media", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, ServerError(err))
		return
	}

	data := make([]models.MediaDTO, len(medias))
	for i := 0; i < len(medias); i++ {
		medias[i].AsDTO(&data[i])
	}
	c.JSON(http.StatusOK, NewResponse(0, "OK", data))
}

func (controller *MediaControllerV1) MediaDetailV1(c *gin.Context) {
	logging.GetLogger().Panic("not implemented")
}

func (controller *MediaControllerV1) MediaAddV1(c *gin.Context) {
	var req struct {
		Paths []string `json:"paths"`
	}
	if err := c.BindJSON(&req); err != nil {
		logging.GetLogger().Error("failed to bind json", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, BadRequest(err))
		return
	}

	medias, err := controller.s.Add(req.Paths...)
	if err != nil {
		logging.GetLogger().Error("failed to add media", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, ServerError(err))
		return
	}
	c.JSON(http.StatusOK, OK(medias))
}

func (controller *MediaControllerV1) MediaUpdateV1(c *gin.Context) {
	logging.GetLogger().Panic("not implemented")
}

func (controller *MediaControllerV1) MediaDeleteV1(c *gin.Context) {
	logging.GetLogger().Panic("not implemented")
}

func (controller *MediaControllerV1) MediaCoverDownloadV1(c *gin.Context) {
	var req struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := c.BindUri(&req); err != nil {
		logging.GetLogger().Error("failed to bind uri", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, BadRequest(err))
		return
	}

	cover, err := controller.s.GetCover(req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(400)
			return
		}

		logging.GetLogger().Error("failed to get cover", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, ServerError(err))
		return
	}

	c.Header("Content-Type", "image/jpeg")
	c.File(cover.Path)
}
