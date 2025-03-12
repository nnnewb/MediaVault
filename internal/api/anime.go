package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/nnnewb/media-vault/internal/logging"
	"github.com/nnnewb/media-vault/internal/models"
	"github.com/nnnewb/media-vault/internal/service"
)

type AnimeControllerV1 struct {
	s *service.AnimeService
}

func NewAnimeControllerV1(s *service.AnimeService) *AnimeControllerV1 {
	return &AnimeControllerV1{s: s}
}

func (a *AnimeControllerV1) RegisterRoutes(router gin.IRouter) {
	g := router.Group("/v1")
	g.POST("/anime/list", a.ListV1)
	g.POST("/anime/search", a.SearchV1)
}

func (a *AnimeControllerV1) ListV1(c *gin.Context) {
	var req struct {
		Pagination
	}
	err := c.BindJSON(&req)
	if err != nil {
		logging.GetLogger().Error("failed to bind json", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, BadRequest(err))
		return
	}

	orderBy := service.OrderBy{}
	pagination := service.Pagination{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	anime, total, err := a.s.List(pagination, orderBy)
	if err != nil {
		logging.GetLogger().Error("failed to list anime", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, ServerError(err))
		return
	}

	data := make([]models.AnimeDTO, len(anime))
	for i := 0; i < len(anime); i++ {
		anime[i].AsDTO(&data[i])
	}
	c.JSON(http.StatusOK, OKList(data, total))
}

// SearchV1 按关键词搜索动画，目前数据源只包含 anime-offline-database
func (a *AnimeControllerV1) SearchV1(c *gin.Context) {
	var req struct {
		Term string `json:"term"`
		Pagination
	}
	err := c.BindJSON(&req)
	if err != nil {
		logging.GetLogger().Error("failed to bind json", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, BadRequest(err))
		return
	}

	by := service.OrderBy{}
	pagination := service.Pagination{Page: req.Page, PageSize: req.PageSize}
	anime, total, err := a.s.Search(req.Term, pagination, by)
	if err != nil {
		logging.GetLogger().Error("failed to search anime", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, ServerError(err))
		return
	}

	data := make([]models.AnimeOfflineDatabaseDTO, len(anime))
	for i := 0; i < len(anime); i++ {
		anime[i].AsDTO(&data[i])
	}

	c.JSON(http.StatusOK, OKList(data, total))
}
