package main

import (
	"github.com/daisakuhazui/four-quadrants-go/backend"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// initialize database
	backend.InitDB()

	// initialize routing
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// TODO: echo によるログ出力がほぼデフォルトのままで見ずらい点を改善する
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	backend.InitRoute(e)
}
