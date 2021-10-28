package openapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGetTask(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/task", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := OasServerImpl{}

	h.GetTask(c)

	if rec.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, rec.Code)
	}
	var got Task
	err := json.NewDecoder(rec.Body).Decode(&got)
	if err != nil {
		t.Fatal("Unable to parse response from server")
	}
	expected := NewDummyTask()

	if got != expected {
		t.Errorf("want %+v, but %+v", expected, got)
	}
}
func TestPostTask(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/task", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := OasServerImpl{}

	h.PostTask(c)

	if rec.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, rec.Code)
	}
	var got Task
	err := json.NewDecoder(rec.Body).Decode(&got)
	if err != nil {
		t.Fatal("Unable to parse response from server")
	}
	expected := NewDummyTask()

	if got != expected {
		t.Errorf("want %+v, but %+v", expected, got)
	}
}

func DeleteTaskTaskId(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/task/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := OasServerImpl{}

	h.DeleteTaskTaskId(c, "1")

	if rec.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, rec.Code)
	}
	var got Task
	err := json.NewDecoder(rec.Body).Decode(&got)
	if err != nil {
		t.Fatal("Unable to parse response from server")
	}
	expected := NewDummyTask()

	if got != expected {
		t.Errorf("want %+v, but %+v", expected, got)
	}
}

func TestGetTasksTaskId(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/task/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := OasServerImpl{}

	h.GetTasksTaskId(c, "1")

	if rec.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, rec.Code)
	}
	var got Task
	err := json.NewDecoder(rec.Body).Decode(&got)
	if err != nil {
		t.Fatal("Unable to parse response from server")
	}
	expected := NewDummyTask()

	if got != expected {
		t.Errorf("want %+v, but %+v", expected, got)
	}
}

func PatchTasksTaskId(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/task/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := OasServerImpl{}

	h.PatchTasksTaskId(c, "1")

	if rec.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, rec.Code)
	}
	var got Task
	err := json.NewDecoder(rec.Body).Decode(&got)
	if err != nil {
		t.Fatal("Unable to parse response from server")
	}
	expected := NewDummyTask()

	if got != expected {
		t.Errorf("want %+v, but %+v", expected, got)
	}
}
