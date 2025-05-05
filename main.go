package main

import (
	"todoapp/db"
	"todoapp/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/todos", handlers.GetTodos)
	r.POST("/todos", handlers.CreateTodo)
	r.PUT("/todos/:id", handlers.UpdateTodo)
	r.DELETE("/todos/:id", handlers.DeleteTodo)
	r.PATCH("/todos/:id/toggle", handlers.ToogleTodoStatus)

	r.Run(":8080")
}
