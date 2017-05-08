package main

import (
	"net/http"
	"github.com/labstack/echo"

)

func main() {
	e := echo.New()
	e.Static("/", "public/dist")
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, This is my blog project!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}