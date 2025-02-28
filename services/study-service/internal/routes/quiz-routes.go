package routes

import "github.com/gin-gonic/gin"

func RegisterQuizRoutes(router *gin.Engine) {
	router.POST("/api/quizzes")
	router.GET("/api/quizzes")
}