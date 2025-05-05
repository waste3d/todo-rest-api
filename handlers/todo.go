package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"todoapp/db"
	"todoapp/models"

	"github.com/gin-gonic/gin"
)

func getIDParam(ctx *gin.Context) (uint, error) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, errors.New("Invalid ID format")
	}
	return uint(id), nil
}

func GetTodos(ctx *gin.Context) {
	var todos []models.Todo

	completedParam := ctx.Query("completed")

	if completedParam != "" {
		completed, err := strconv.ParseBool(completedParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'completed' query param"})
			return
		}
		db.DB = db.DB.Where("completed = ?", completed)
	}

	if err := db.DB.Find(&todos).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read todos"})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func CreateTodo(ctx *gin.Context) {
	var input models.Todo

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&input).Error; err != nil {
		log.Print("db error:", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	ctx.JSON(http.StatusCreated, input)
}

func UpdateTodo(ctx *gin.Context) {
	id, err := getIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo models.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var input models.Todo
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.Title = input.Title
	todo.Completed = input.Completed

	if err := db.DB.Save(&todo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update todo"})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func DeleteTodo(ctx *gin.Context) {
	id, err := getIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo models.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	if err := db.DB.Delete(&todo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete todo"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func ToogleTodoStatus(ctx *gin.Context) {
	id, err := getIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo models.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := db.DB.Save(&todo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to toogle todo"})
	}

	ctx.JSON(http.StatusOK, todo)
}
