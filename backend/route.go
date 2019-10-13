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
	e.GET("/task/:id", handlerTaskGet)
	e.POST("/task/:id", handlerTaskPost)
	e.PUT("/task/:id", handlerTaskPut)
	e.DELETE("/task/:id", handlerTaskDelete)

	// start server 8080 port
	e.Logger.Fatal(e.Start(":8080"))
}
