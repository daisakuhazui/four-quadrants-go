package backend

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRoute(e *echo.Echo) {
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routing
	e.GET("/hello", handlerTasksGet)

	// start server 8080 port
	e.Logger.Fatal(e.Start(":8080"))
}
