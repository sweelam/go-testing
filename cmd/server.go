package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	api "database.learning/start/internal/handlers"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	routes(router)

	log.Println("Server started on port: 8080")

	if err := router.Run(":8080"); err != nil {
		log.Printf("Error: %v", err)
	}
}

func routes(router *gin.Engine) {
	router.GET("api/v1/jobs", api.GetAllJobs)
	router.GET("api/v1/person/:name", predictByName)
}

func predictByName(c *gin.Context) {
	name := c.Params.ByName("name")
	log.Println("Name passed is: ", name)
	response, err := http.Get("https://api.nationalize.io/?name=" + name)

	if err != nil {
		log.Printf("Error reading response body: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer response.Body.Close()

	bodyRes, err := io.ReadAll(response.Body)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, json.RawMessage(bodyRes))
}
