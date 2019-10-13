package main

import (
	"github.com/daisakuhazui/four-quadrants-go/backend"
	"github.com/labstack/echo"
)

func main() {
	// initialize database
	backend.InitDB()

	// initialize routing
	e := echo.New()
	backend.InitRoute(e)
}
