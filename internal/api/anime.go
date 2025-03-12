package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gitee.com/uniqptr/media-vault.git/internal/logging"
)

type AnimeControllerV1 struct {
	db *gorm.DB
}

// AnimeListV1 获取动画列表
func (controller *AnimeControllerV1) AnimeListV1(c *gin.Context) {
	logging.GetLogger().Panic("not implemented")
}

// AnimeDetailV1 获取指定动画的详细信息
func (controller *AnimeControllerV1) AnimeDetailV1(c *gin.Context) {
	logging.GetLogger().Panic("not implemented")
}

// AnimeAddV1 添加动画，指定类型和路径
func (controller *AnimeControllerV1) AnimeAddV1(c *gin.Context) {
	logging.GetLogger().Panic("not implemented")
}

// AnimeUpdateV1 更新动画基本信息
func (controller *AnimeControllerV1) AnimeUpdateV1(c *gin.Context) {
	logging.GetLogger().Panic("not implemented")
}
