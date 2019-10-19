package backend

import (
	"log"
	"net/http"
	"time"

	"github.com/daisakuhazui/four-quadrants-go/common"
	"github.com/labstack/echo"
)

func handlerTaskGet(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
}

func handlerTaskPost(c echo.Context) error {
	// bind request
	task := new(common.Task)
	if err := c.Bind(task); err != nil {
		log.Printf("Bad request: %+v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// open database
	db, openErr := OpenDB()
	if openErr != nil {
		panic(openErr)
	}

	result, err := db.Exec(
		`INSERT INTO TASKS (NAME, MEMO, QUADRANT, COMPLETEFLAG, CREATEDAT, UPDATEDAT) VALUES (?, ?, ?, ?, ?, ?)`,
		task.Name,
		task.Memo,
		task.Quadrant,
		task.CompleteFlag,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		log.Printf("Could not insert task: %+v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Could not get lastInserID: %+v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	task.ID = lastInsertID

	return c.JSON(http.StatusOK, task)
}

func handlerTaskPut(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"task": "put"})
}

func handlerTaskDelete(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"task": "delete"})
}
