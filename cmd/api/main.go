package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/otherpirate/volume-challenge/internal/api"
)

func main() {
	router := gin.Default()
	router.POST("/calculate", api.Calculate)
	router.POST("/paths", api.Paths)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))

}
