package app

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

type App struct {
	DB   *sql.DB
	Port int
}

func (a *App) Run() error {
	router := gin.Default()

	err := router.SetTrustedProxies(nil)
	if err != nil {
		return err
	}

	return RegisterRoutes(router, a)
}

func RegisterRoutes(router *gin.Engine, a *App) error {

	v1 := router.Group("/api/v1")

	tasks := v1.Group("/tasks")

	{
		tasks.GET("/", func(context *gin.Context) {
			context.JSON(200, []int{1, 2, 3, 4, 5})
		})
	}

	err := router.Run(fmt.Sprintf(":%d", a.Port))
	if err != nil {
		return err
	}

	return nil
}
