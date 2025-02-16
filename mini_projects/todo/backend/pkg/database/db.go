package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"to-do/config"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	var dsn = fmt.Sprintf(
		"postgres://%v:%v@%v:%v/todo?sslmode=disable",
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.Host,
		cfg.Db.Port,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	err = initializeDB(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initializeDB(db *sql.DB) (err error) {
	query := `CREATE TABLE IF NOT EXISTS todo (
		  id SERIAL PRIMARY KEY,
		  title VARCHAR (200) ,
		  created_at TIMESTAMP DEFAULT now() NOT NULL,
		  updated_at TIMESTAMP DEFAULT now()  NOT NULL
		);`

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
