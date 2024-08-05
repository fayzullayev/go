package handlers

import (
	"advice/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

type Advice struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetAdvices(c *gin.Context) {
	query := "SELECT * FROM advice"

	rows, err := db.DB.Query(query)

	var advices []Advice

	for rows.Next() {
		var advice Advice
		err = rows.Scan(&advice.Id, &advice.Name)
		if err != nil {
			fmt.Println("err", err)
			return
		}
		advices = append(advices, advice)
	}

	randInt := rand.Intn(len(advices))

	//fmt.Println("advices", advices)

	c.JSON(200, gin.H{
		"data": advices[randInt],
	})

	return
}
