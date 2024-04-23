package app

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

func (app *App) Run() error {
	err := app.Router.Run(":5555")
	if err != nil {
		return err
	}

	return nil
}

func RegisterRouter(app *App) {
	app.Router.GET("/", app.Handlers.HomeHandler.HomePage)
	app.Router.GET("/users", app.Handlers.UsersHandler.GetUsers)
	app.Router.GET("/tasks", app.Handlers.TasksHandler.GetTasks)
}
