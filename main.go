package main

import (
	"github.com/daisakuhazui/four-quadrants-go/backend"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	backend.InitRoute(e)
}
