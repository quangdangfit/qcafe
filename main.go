package main

import (
	"github.com/labstack/echo/v4"
	"qcafe/routers"
)

func main() {
	e := echo.New()

	routers.SetRouters(e)

	e.Logger.Fatal(e.Start(":8080"))
}
