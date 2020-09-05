package main

import (
	"time"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		request_time := time.Now().Format("2006/01/02 15:04:05.000")
		ctx.HTML(200, "index.html", gin.H{ "request_time": request_time })
	})
	router.Run()
}