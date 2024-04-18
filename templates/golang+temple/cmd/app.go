package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"tempaltes-tutorial/db"
	"tempaltes-tutorial/handler"
)

type App struct {
	DB       *sql.DB
	Handlers *handler.Handlers
	Router   *gin.Engine
}

func (app *App) Run() error {
	err := app.Router.Run(":3333")
	if err != nil {
		return err
	}

	return nil
}

func NewApp() (*App, error) {

	DB, err := db.InitDB()

	if err != nil {
		return nil, err
	}

	router := gin.Default()

	err = router.SetTrustedProxies(nil)
	if err != nil {
		return nil, err
	}

	app := &App{
		DB:       DB,
		Router:   router,
		Handlers: handler.NewHandlers(DB),
	}

	RegisterRouter(app)

	return app, nil
}
