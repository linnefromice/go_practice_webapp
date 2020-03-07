package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	title := "Hello World!"
	subtitle := "Hello go-gin!"

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"title": title, "subtitle": subtitle})
	})

	router.Run()
}
