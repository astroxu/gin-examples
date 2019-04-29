package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//Disable log`s color
	//gin.DisableConsoleColor()

	//Creates a gin router with default middleware:
	//logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.Run(":8088")
}
