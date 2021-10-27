package openapi

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

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
	expected := `{"description":"Description","id":1,"isDeleted":false,"isFinished":false,"title":"Title","version":1}`
	if reflect.DeepEqual(rec.Body.String(), expected) {
		t.Errorf("want %s, but %s", expected, rec.Body.String())
	}
}
