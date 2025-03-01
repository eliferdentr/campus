package routes

import "github.com/gin-gonic/gin"


func RegisterFlashcardRoutes(router *gin.Engine) {
	router.POST("/api/flashcards")
	router.GET("/api/flashcards")

}