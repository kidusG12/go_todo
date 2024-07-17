package main

import (
	"github.com/gin-gonic/gin"
	"x.com/todo/database"
	"x.com/todo/routes"
)

func main(){
	database.Init()

	server := gin.Default()

	routes.TodoRoutes(server)

	server.Run(":8000")
}