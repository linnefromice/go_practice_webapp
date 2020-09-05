package main

import (
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Todo struct {
    gorm.Model
    Text   string
    Status string
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}

	// db initialize
	db.AutoMigrate(&Todo{})
	var todos []Todo
	db.Order("created_at desc").Find(&todos)

	// add data
	title := "Homework - " + strconv.Itoa(len(todos) + 1)
	db.Create(&Todo{Text: title, Status: "DOING"})

	// routing
	router.GET("/", func(ctx *gin.Context) {
		request_time := time.Now().Format("2006/01/02 15:04:05.000")
		db.Order("created_at desc").Find(&todos)
		ctx.HTML(200, "index.html", gin.H{
			"request_time": request_time,
			"todos" : todos,
		})
	})
	router.Run()
}