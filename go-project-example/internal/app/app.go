package app

import (
	"database/sql"
	"github.com/fayzullayev/go-project-example/api/handlers"
	"github.com/gin-gonic/gin"
)

type App struct {
	db   *sql.DB
	port string
}

func (app *App) Run() error {
	var err error

	router := gin.Default()

	err = router.SetTrustedProxies(nil)
	if err != nil {
		return err
	}

	handlers.Register(router, app)

	err = router.Run(":" + app.port)
	if err != nil {
		return err
	}

	return nil
}

func New(port string, db *sql.DB) *App {
	return &App{
		db:   db,
		port: port,
	}
}
