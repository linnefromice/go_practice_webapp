package main

import (
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

const (
	dbName = "test.db"
	statusDone = "Done"
	statusDoing = "Doing"
	statusReady = "Ready"
	statusPending = "Pending"
)

type Todo struct {
    gorm.Model
    Text   string
    Status string
}

func dbInit() {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database (Init)")
	}
	db.AutoMigrate(&Todo{})
}

func dbInsert(text string, status string) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database (Insert)")
	}
	db.Create(Todo{Text: text, Status: status})
}

func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database (Update)")
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
}

func dbDelete(id int) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database (Delete)")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
}

func dbSelectOne(id int) Todo {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database (SelectOne)")
	}
	var todo Todo
	db.First(&todo, id)
	return todo
}

func dbSelectAll() []Todo {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database (SelectAll)")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	return todos
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
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
		requestTime := time.Now().Format("2006/01/02 15:04:05.000")
		db.Order("created_at desc").Find(&todos)
		ctx.HTML(200, "index.html", gin.H{
			"requestTime": requestTime,
			"todos" : todos,
		})
	})
	router.Run()
}