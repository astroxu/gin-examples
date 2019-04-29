package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Goyuyan",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	//Listen and server on 0.0.0.0:8080
	r.Run(":8080")
}
