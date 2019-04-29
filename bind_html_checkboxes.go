package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(http.StatusOK, gin.H{
		"color": fakeForm.Colors,
	})
}

func main() {
	r := gin.Default()
	//r.LoadHTMLFiles("./views/form.html")
	r.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/github.com/chengango/gin-examples/views/*"))
	r.GET("/", indexHandler)
	r.POST("/", formHandler)
	r.Run(":8088")
}
