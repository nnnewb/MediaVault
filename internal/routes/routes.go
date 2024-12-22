package routes

import "github.com/gin-gonic/gin"

func RegisterAppRoutes(router gin.IRouter) {
	api := router.Group("/api")
	RegisterAnimeRoutes(api)
}
