package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "res/test.html")
	e.GET("/d", download)

	e.Logger.Fatal(e.Start(":8080"))
}

func download(ctx echo.Context) error {
	b := make([]byte, 10240)

	return ctx.Blob(http.StatusOK, "blob", b)
}
