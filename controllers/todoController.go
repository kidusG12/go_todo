package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"x.com/todo/models"
)

func GetTodos(context *gin.Context) {
	userId := 1

	todos, err := models.GetUsersTodos(int64(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "error",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"result": todos,
	})

}

func CreateTodos(context *gin.Context) {
	var todo models.Todo
	err := context.ShouldBind(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "title is required ",
		})
		return
	}

	todo.UserId = 1

	err = todo.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Could not save todo",
		})
		panic(fmt.Sprintf("Could not save todo\n%v", err))
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":  "ok",
		"message": "todo created",
	})
}

func UpdateTodos(context *gin.Context) {
	id, err := strconv.ParseInt((context.Param("id")), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "invalid todo id",
		})
		return
	}

	storedTodo, err := models.GetTodo(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "todo could not be found",
		})
		return
	}

	var updatedTodo models.Todo
	err = context.ShouldBind(&updatedTodo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "could not parse data",
		})
		return
	}

	updatedTodo.Id = storedTodo.Id

	err = updatedTodo.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "could not update todo",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "todo updated",
	})
}

func DeleteTodos(context *gin.Context) {
	id, err := strconv.ParseInt((context.Param("id")), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "invalid todo id",
		})
		return
	}

	storedTodo, err := models.GetTodo(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "todo could not be found",
		})
		return
	}

	err = storedTodo.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "could not delete todo",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "todo deleted",
	})
}
