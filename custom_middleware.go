package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		//Set example variable
		c.Set("example", "21232")

		//before request

		c.Next()

		//after request
		latency := time.Since(t)
		log.Print(latency)

		//access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		//it wpold pring: "21232"
		log.Println(example)
	})

	//Listen and serve on 0.0.0.0:8088
	r.Run(":8088")
}
