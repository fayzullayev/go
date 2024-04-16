package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project-managememnt/models"
	"strconv"
)

func CreateProjectHandler(c *gin.Context) {
	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectId, err := project.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"projectId": projectId})
}

func GetProjectsHandler(c *gin.Context) {
	projects, err := models.GetProjects()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": projects,
	})

	return
}

func GetProjectHandler(c *gin.Context) {
	id := c.Param("id")
	projectId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Project id %s not found1", id),
		})
		return
	}

	project := models.Project{
		ID: projectId,
	}

	err = project.GetByID()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": project,
	})
}

func DeleteProjectHandler(c *gin.Context) {
	id := c.Param("id")
	projectId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Bad request",
		})
		return
	}

	project := models.Project{
		ID: projectId,
	}

	err = project.Delete()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Something went wrong: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project deleted",
	})

}
