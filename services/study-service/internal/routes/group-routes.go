package routes

import ("github.com/gin-gonic/gin"
controller "campus-project.com/study-service/internal/controller")


func RegisterGroupRoutes(router *gin.Engine) {
	router.POST("/api/groups", controller.CreateGroup)
	router.GET("/api/groups")
	router.POST("/api/groups/:id/invite")
	router.DELETE("/api/groups/:id")

}