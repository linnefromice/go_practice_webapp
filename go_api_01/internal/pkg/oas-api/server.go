package openapi

import (
	models "linnefromice/go_practice_webapp/go_api_01/internal/pkg/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OasServerImpl struct {
	TaskList models.TaskList
}

func NewDummyTask() models.Task {
	return models.Task{
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
	reqBody := new(PostTaskJSONBody)
	if err := ctx.Bind(reqBody); err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	task := s.TaskList.AddTask(*reqBody.Title, *reqBody.Description)
	return ctx.JSON(http.StatusOK, task)
}

func (s OasServerImpl) DeleteTaskTaskId(ctx echo.Context, taskId string) error {
	id, err := strconv.Atoi(taskId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "error occured when parse taskId because taskId is not positive number")
	}

	task := s.TaskList.DeleteTask(id)
	return ctx.JSON(http.StatusOK, task)
}

func (s OasServerImpl) GetTasksTaskId(ctx echo.Context, taskId string) error {
	id, err := strconv.Atoi(taskId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "error occured when parse taskId because taskId is not positive number")
	}

	task, err := s.TaskList.GetTask(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, task)
}

func (s OasServerImpl) PatchTasksTaskId(ctx echo.Context, taskId string) error {
	id, err := strconv.Atoi(taskId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "error occured when parse taskId because taskId is not positive number")
	}

	reqBody := new(PatchTasksTaskIdJSONBody)
	if err := ctx.Bind(reqBody); err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	task := s.TaskList.UpdateTask(id, *reqBody.Title, *reqBody.Description)
	return ctx.JSON(http.StatusOK, task)
}
