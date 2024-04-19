package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"tempaltes-tutorial/db"
	"tempaltes-tutorial/handler"
)

func main() {
	app, err := NewApp()
	if err != nil {
		log.Fatal("Failed to create app:", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}

type App struct {
	DB       *sql.DB
	Handlers *handler.Handlers
	Router   *gin.Engine
}

func (app *App) Run() error {
	err := app.Router.Run(":7777")
	if err != nil {
		return err
	}

	return nil
}

func RegisterRouter(app *App) {
	app.Router.GET("/", app.Handlers.UsersHandler.Home)
	app.Router.GET("/users", app.Handlers.UsersHandler.GetUsers)
	app.Router.GET("/tasks", app.Handlers.TasksHandler.GetTasks)
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
