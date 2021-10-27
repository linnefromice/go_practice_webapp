package main

import (
	"github.com/labstack/echo/v4"

	openapi "linnefromice/go_practice_webapp/go_api_01/internal/pkg/oas-api"
)

func main() {
	handler := openapi.OasServerImpl{}
	e := echo.New()
	openapi.RegisterHandlers(e, handler)
	e.Logger.Fatal(e.Start(":5000"))
}
