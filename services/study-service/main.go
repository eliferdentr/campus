package studyservice

import (
	"campus-project.com/study-service/internal/routes"
	"github.com/gin-gonic/gin"
)

func main () {
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run()
	
}