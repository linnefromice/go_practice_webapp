package main

import (
	"linnefromice/go_practice_webapp/go_api_01/internal/pkg/models"
	openapi "linnefromice/go_practice_webapp/go_api_01/internal/pkg/oas-api"

	"github.com/labstack/echo/v4"
)

func main() {
	taskList := models.NewTaskList()
	// Dummy data
	taskList.AddTask("title 1", "description 1")
	taskList.AddTask("title 2", "description 2")
	taskList.AddTask("title 3", "description 3")

	handler := openapi.OasServerImpl{
		TaskList: *taskList,
	}
	e := echo.New()
	openapi.RegisterHandlers(e, handler)
	e.Logger.Fatal(e.Start(":5000"))
}
