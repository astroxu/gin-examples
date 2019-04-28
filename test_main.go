package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Goyuyan",
			"tag":  "<br>",
		}

	})
}
