package backend

import (
	"net/http"

	"github.com/labstack/echo"
)

func handlerTaskGet(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
}

func handlerTaskPost(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"task": "post"})
}

func handlerTaskPut(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"task": "put"})
}

func handlerTaskDelete(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"task": "delete"})
}
