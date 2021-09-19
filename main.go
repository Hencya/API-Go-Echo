package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	echoApp := echo.New()

	echoApp.GET("/Hello",HelloworldController)

	echoApp.Start(":8080")
}

func HelloworldController(ctx echo.Context) error{
	response := struct {
		Data string `json:"data"`
	}{
		Data : "Hello World",
	}
	return ctx.JSON(http.StatusOK,response)
}

