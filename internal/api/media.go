package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gitee.com/uniqptr/media-vault.git/internal/logging"
)

type MediaControllerV1 struct {
	db *gorm.DB
}

// MediaListV1 列出媒体
func (controller *MediaControllerV1) MediaListV1(c *gin.Context) {
	logging.GetLogger().Panic("not implemented")
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
