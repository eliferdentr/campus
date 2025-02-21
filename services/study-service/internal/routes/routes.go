package routes
import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	RegisterGroupRoutes(router)
	RegisterLiveSessionRoutes(router)
	RegisterFlashcardRoutes(router)
	RegisterQuizRoutes(router)
}