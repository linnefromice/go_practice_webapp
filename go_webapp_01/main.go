package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	content := "Hello World!"
	subContent := "Hello go-gin!"

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"content": content, "subContent": subContent})
	})

	router.Run()
}
