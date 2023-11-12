package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title" binding:"required"`
}

func main() {
	r := gin.Default()

	db := initDB()
	defer db.Close()

	r.GET("/todos", func(c *gin.Context) {

		rows, err := db.Query("SELECT * FROM todos")
		defer rows.Close()

		if err != nil {
			log.Fatal(err)
		}

		var todos []Todo

		for rows.Next() {
			var todo Todo
			err = rows.Scan(&todo.Id, &todo.Title)

			if err != nil {
				log.Fatal(err)
			}

			todos = append(todos, todo)
		}

		err = rows.Err()

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"todos": todos,
		})
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		value, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("User with id:%s not found", id),
			})
			return
		}

		rows := db.QueryRow("SELECT * FROM todos WHERE id = $1", value)

		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Title)

		for err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Error querying todo item"})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"todo": todo,
		})
	})

	r.POST("/todos", func(c *gin.Context) {
		var todo Todo

		if err := c.ShouldBindJSON(&todo); err != nil {

			//fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}

		sqlStatement := `INSERT INTO todos (title) VALUES ($1) RETURNING id`

		id := 0
		err := db.QueryRow(sqlStatement, todo.Title).Scan(&id)

		if err != nil {
			log.Fatalf("Unable to execute the query. %v", err)
		}

		c.JSON(http.StatusOK, gin.H{"message": "Todo inserted successfully", "id": id})

		//fmt.Printf("todo: %+v\n", todo)

	})

	r.Run(":8081")

}

func initDB() *sql.DB {

	const (
		host     = "localhost"
		port     = 5432 // Default port
		user     = "mirodil"
		password = "123456789"
		dbname   = "todos"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("Error pinging database: %q", err)
	}

	fmt.Println("Successfully connected!")
	return db
}
