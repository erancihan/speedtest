package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
	"time"
)

func main() {
	e := echo.New()
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:63342", "http://speed.erancihan.xyz"},
	}))

	e.File("/", "res/test.html")
	e.GET("/d", download)
	e.POST("/u", upload)

	e.Logger.Fatal(e.Start(":8080"))
}

func download(ctx echo.Context) error {
	bSize, _ := strconv.Atoi(ctx.QueryParam("b"))

	b := make([]byte, bSize)

	return ctx.Blob(http.StatusOK, "blob", b)
}

func upload(ctx echo.Context) error {
	ts := time.Now().UnixNano()

	return ctx.JSON(http.StatusOK, ts)
}
