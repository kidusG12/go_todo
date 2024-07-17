package routes

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(server *gin.Engine){
	server.POST("/login");
	server.POST("/signup")

}