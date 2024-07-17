package routes

import (
	"github.com/gin-gonic/gin"
	"x.com/todo/controllers"
)

func TodoRoutes(server *gin.Engine){
	todoGroup :=server.Group("/todos")

	todoGroup.GET("", controllers.GetTodos)
	todoGroup.POST("", controllers.CreateTodos)
	todoGroup.PUT("/:id", controllers.UpdateTodos)
	todoGroup.DELETE("/:id", controllers.DeleteTodos)
}