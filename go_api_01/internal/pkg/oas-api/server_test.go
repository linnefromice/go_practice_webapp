package openapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"

	"linnefromice/go_practice_webapp/go_api_01/internal/pkg/models"
)

func TestGetTask(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/task", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := OasServerImpl{
		TaskList: createInitialTaskList(),
	}

	h.GetTask(c)

	if rec.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, rec.Code)
	}
	var got []Task
	err := json.NewDecoder(rec.Body).Decode(&got)
	if err != nil {
		t.Fatal("Unable to parse response from server")
	}
	expected := []Task{
		{
			Id:          1,
			Title:       "title 1",
			Description: "description 1",
			IsFinished:  false,
			IsDeleted:   false,
			Version:     1,
		},
		{
			Id:          2,
			Title:       "title 2",
			Description: "description 2",
			IsFinished:  false,
			IsDeleted:   false,
			Version:     1,
		},
		{
			Id:          3,
			Title:       "title 3",
			Description: "description 3",
			IsFinished:  false,
			IsDeleted:   false,
			Version:     1,
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("want %+v, but %+v", expected, got)
	}
}

func TestPostTask(t *testing.T) {
	e := echo.New()

	userJSON := `{"title":"title 4","description":"description 4"}`
	req := httptest.NewRequest(http.MethodPost, "/task", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := OasServerImpl{
		TaskList: createInitialTaskList(),
	}

	h.PostTask(c)

	if rec.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, rec.Code)
	}
	var got models.Task
	err := json.NewDecoder(rec.Body).Decode(&got)
	if err != nil {
		t.Fatal("Unable to parse response from server")
	}
	expected := models.Task{
		Id:          4,
		Title:       "title 4",
		Description: "description 4",
		IsFinished:  false,
		IsDeleted:   false,
		Version:     1,
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("want %+v, but %+v", expected, got)
	}
}

func TestDeleteTaskTaskId(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/task/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := OasServerImpl{
		TaskList: createInitialTaskList(),
	}

	h.DeleteTaskTaskId(c, "2")

	if rec.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, rec.Code)
	}
	var got models.Task
	err := json.NewDecoder(rec.Body).Decode(&got)
	if err != nil {
		t.Fatal("Unable to parse response from server")
	}
	expected := models.Task{
		Id:          2,
		Title:       "title 2",
		Description: "description 2",
		IsFinished:  false,
		IsDeleted:   true,
		Version:     2,
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("want %+v, but %+v", expected, got)
	}
}

func TestGetTasksTaskId(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/task/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := OasServerImpl{
		TaskList: createInitialTaskList(),
	}

	h.GetTasksTaskId(c, "3")

	if rec.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, rec.Code)
	}
	var got Task
	err := json.NewDecoder(rec.Body).Decode(&got)
	if err != nil {
		t.Fatal("Unable to parse response from server")
	}
	expected := Task{
		Id:          3,
		Title:       "title 3",
		Description: "description 3",
		IsFinished:  false,
		IsDeleted:   false,
		Version:     1,
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("want %+v, but %+v", expected, got)
	}
}

func TestPatchTasksTaskId(t *testing.T) {
	e := echo.New()
	userJSON := `{"title":"title updated","description":"description updated"}`
	req := httptest.NewRequest(http.MethodPost, "/task/1", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := OasServerImpl{
		TaskList: createInitialTaskList(),
	}

	h.PatchTasksTaskId(c, "2")

	if rec.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, rec.Code)
	}
	var got models.Task
	err := json.NewDecoder(rec.Body).Decode(&got)
	if err != nil {
		t.Fatal("Unable to parse response from server")
	}
	expected := models.Task{
		Id:          2,
		Title:       "title updated",
		Description: "description updated",
		IsFinished:  false,
		IsDeleted:   false,
		Version:     2,
	}

	if got != expected {
		t.Errorf("want %+v, but %+v", expected, got)
	}
}

func createInitialTaskList() models.TaskList {
	taskList := models.NewTaskList()
	taskList.AddTask("title 1", "description 1")
	taskList.AddTask("title 2", "description 2")
	taskList.AddTask("title 3", "description 3")
	return *taskList
}
