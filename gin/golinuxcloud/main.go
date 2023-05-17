package main

import (
	"api/api"
	"api/database"
)

func init() {
	database.NewPostgresSQLClient()
}

func main() {

	r := api.SetupRouter()

	r.Run(":5000")
}
