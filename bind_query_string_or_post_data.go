package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form"address"`
	Birthday time.Time `form:"nirthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", startPage)
	r.Run(":8088")
}

func startPage(c *gin.Context) {
	var person Person

	//If `GET`,only `Form` binding engine(`query`) used.
	//If `POST`, first checks the `content-type` for `JSON` or `XML`,then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(http.StatusOK, "Success")
}
