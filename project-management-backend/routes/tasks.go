package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project-managememnt/models"
	"strconv"
)

func CreateTasksHandler(c *gin.Context) {
	task := models.Task{}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	projectId, err := task.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"data":    gin.H{"projectId": projectId},
	})

}

func GetTasksHandler(c *gin.Context) {
	id := c.Param("id")
	projectId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Project id %s not found1", id),
		})
		return
	}

	//fmt.Println("projectId", projectId)

	tasks, err := models.GetTasks(projectId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}
