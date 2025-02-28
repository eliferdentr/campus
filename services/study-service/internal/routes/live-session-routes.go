package routes

import "github.com/gin-gonic/gin"

func RegisterLiveSessionRoutes(router *gin.Engine) {
	router.POST("/api/groups/:id/live-session")
	router.GET("/api/groups/:id/live-sessions")

}