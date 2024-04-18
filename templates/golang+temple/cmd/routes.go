package main

func RegisterRouter(app *App) {
	app.Router.GET("/users", app.Handlers.UsersHandler.GetUsers)
	app.Router.GET("/tasks", app.Handlers.TasksHandler.GetTasks)
}
