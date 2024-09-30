package main

import (
	"database/sql"
	"fido-kpi/internal/app"
	"fido-kpi/internal/database"
	"github.com/spf13/viper"
	"log"
)

func main() {
	var err error

	db, err := database.Init()

	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	viper.SetConfigFile("config/config.yaml")

	if err = viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	myApp := app.App{
		DB:   db,
		Port: viper.GetInt("app.port"),
	}

	if err = myApp.Run(); err != nil {
		log.Fatal(err)
	}

}
