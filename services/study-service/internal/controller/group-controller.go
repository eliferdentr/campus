package controller

import (
	"net/http"

	request "campus-project.com/study-service/internal/request"
	"github.com/gin-gonic/gin"
)

func CreateGroupHandler (context *gin.Context) {
	var req request.CreateGroupRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest,  gin.H{"error": err.Error()})
		return
	}
	

}

func InviteToGroupHandler (context *gin.Context) {
	
}

func GetGroupsHandler (context *gin.Context) {
	
}

func DeleteGroupHandler (context *gin.Context) {
	
}