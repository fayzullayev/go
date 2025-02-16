package main

import (
	"database/sql"
	"to-do/internal/server"

	//"database/sql"
	"log"
	"to-do/config"
	"to-do/pkg/database"
	//"to-do/internal/server"
	//"to-do/pkg/database"
)

func main() {
	/*--------------------------------------*/
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	/*--------------------------------------*/
	/*--------------------------------------*/
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}(db)
	
	/*--------------------------------------*/

	// 3. Start server
	if err = server.StartServer(cfg, db); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
