package routes

import (
	"campus-project.com/study-service/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/api/groups", controller.CreateGroupHandler)
	router.GET("/api/groups")
	router.POST("/api/groups/:id/invite")
	router.DELETE("/api/groups/:id")
}