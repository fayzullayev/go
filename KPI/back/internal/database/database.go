package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Init() (*sql.DB, error) {
	var db *sql.DB
	var err error

	viper.SetConfigFile("config/config.yaml")

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var dns = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"),
		viper.GetString("database.sslmode"),
	)

	db, err = sql.Open("postgres", dns)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to database")

	return db, nil
}
