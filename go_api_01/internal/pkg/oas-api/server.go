package openapi

import (
	"linnefromice/go_practice_webapp/go_api_01/internal/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OasServerImpl struct {
	TaskList models.TaskList
}

func NewDummyTask() Task {
	return Task{
		Description: "Description",
		Id:          1,
		IsDeleted:   false,
		IsFinished:  false,
		Title:       "Title",
		Version:     1,
	}
}

func (s OasServerImpl) GetTask(ctx echo.Context) error {
	tasks := s.TaskList.GetAllTasks()
	return ctx.JSON(http.StatusOK, tasks)
}

func (s OasServerImpl) PostTask(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, NewDummyTask())
}

func (s OasServerImpl) DeleteTaskTaskId(ctx echo.Context, taskId string) error {
	return ctx.JSON(http.StatusOK, NewDummyTask())
}

func (s OasServerImpl) GetTasksTaskId(ctx echo.Context, taskId string) error {
	return ctx.JSON(http.StatusOK, NewDummyTask())
}

func (s OasServerImpl) PatchTasksTaskId(ctx echo.Context, taskId string) error {
	return ctx.JSON(http.StatusOK, NewDummyTask())
}
