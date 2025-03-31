package api

import (
	"log"
	"net/http"

	database "database.learning/start/internal/database/gorm"
	"github.com/gin-gonic/gin"
)

func GetAllJobs(c *gin.Context) {
	jobs, err := database.FindAllJobs()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Jobs size: ", len(jobs))
	c.IndentedJSON(http.StatusOK, jobs)
}
