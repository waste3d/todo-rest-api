package handlers

import (
	"net/http"
	"strconv"
	"todoapp/db"
	"todoapp/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(ctx *gin.Context) {
	var todos []models.Todo
	db.DB.Find(&todos)
	ctx.JSON(http.StatusOK, todos)
}

func CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&todo)
	ctx.JSON(http.StatusOK, todo)
}

func UpdateTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var todo models.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.ID = uint(id)
	if err := db.DB.Save(&todo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func DeleteTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var todo models.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := db.DB.Delete(&todo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
