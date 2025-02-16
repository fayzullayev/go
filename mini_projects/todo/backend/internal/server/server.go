package server

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"to-do/config"
	"to-do/internal/domain/repository"
	"to-do/internal/domain/services"
	"to-do/internal/http/handlers"
)

func StartServer(cfg *config.Config, db *sql.DB) error {
	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		return err
	}

	// Initialize repository & service
	todoRepo := repository.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	// Register routes
	v1 := router.Group("/api/v1")
	{
		v1.POST("/todos", todoHandler.CreateTodo)
		v1.GET("/todos", todoHandler.ListTodos)
		v1.GET("/todos/:id", todoHandler.GetTodo)
		v1.PUT("/todos/:id", todoHandler.UpdateTodo)
		v1.DELETE("/todos/:id", todoHandler.DeleteTodo)
	}

	addr := fmt.Sprintf(":%s", cfg.App.Port)
	return router.Run(addr)
}
