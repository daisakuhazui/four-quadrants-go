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

func handlerAllTasksGet(c echo.Context) error {
	db, openErr := OpenDB()
	if openErr != nil {
		panic(openErr)
	}

	// TODO: ログインユーザーのタスクのみ取得されるようにする
	rows, queryErr := db.Query(
		`SELECT * FROM TASKS`,
	)
	if queryErr != nil {
		log.Fatal("DBからの取得に失敗した")
		panic(queryErr)
	}
	defer rows.Close()

	var tasks []common.Task
	for rows.Next() {
		var task common.Task
		if err := rows.Scan(
			&task.ID,
			&task.Name,
			&task.Memo,
			&task.Quadrant,
			&task.CompleteFlag,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			log.Fatalf("Unexpected error occurs during rows.Scan(): %+v", err)
			return err
		}
		tasks = append(tasks, task)
	}

	return c.JSON(http.StatusOK, tasks)
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
