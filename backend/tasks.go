package backend

import (
	"database/sql"
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
			log.Printf("Unexpected error occurs during rows.Scan(): %+v", err)
			return err
		}
		if !task.CompleteFlag {
			tasks = append(tasks, task)
		}
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
	task := new(common.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := updateTask(task); err != nil {
		log.Printf("Unexpected error occured during Task id:%v updating: ERROR %+v", task.ID, err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, map[string]string{"task": "put"})
}

func handlerTaskCheck(c echo.Context) error {
	taskID := c.Param("id")

	task, err := selectTask(taskID)
	if err == sql.ErrNoRows {
		log.Printf("Task id:%v doesn't exist in datastore", taskID)
		return c.JSON(http.StatusNotFound, nil)
	} else if err != nil {
		log.Printf("Unexpected error occured during Task id:%v row.Scan(): ERROR %+v", taskID, err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if task.CompleteFlag {
		task.CompleteFlag = false
	} else {
		task.CompleteFlag = true
	}
	if err := updateTask(task); err != nil {
		log.Printf("Unexpected error occured during Task id:%v CompleteFlag turning %v: ERROR %+v", task.ID, task.CompleteFlag, err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, task)
}

func handlerTaskDelete(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"task": "delete"})
}

func selectTask(taskID string) (*common.Task, error) {
	// open database
	db, openErr := OpenDB()
	if openErr != nil {
		panic(openErr)
	}

	// TDOO: この処理は後々切り出せるはず
	row := db.QueryRow(
		`SELECT * FROM TASKS WHERE ID=?`,
		taskID,
	)

	var task *common.Task
	err := row.Scan(
		&task.ID,
		&task.Name,
		&task.Memo,
		&task.Quadrant,
		&task.CompleteFlag,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	return task, err
}

func updateTask(task *common.Task) error {
	// open database
	db, openErr := OpenDB()
	if openErr != nil {
		panic(openErr)
	}

	task.UpdatedAt = time.Now()
	_, execErr := db.Exec(
		`UPDATE TASKS SET NAME=?, MEMO=?, QUADRANT=?, COMPLETEFLAG=?, CREATEDAT=?, UPDATEDAT=? WHERE ID=?`,
		task.Name,
		task.Memo,
		task.Quadrant,
		task.CompleteFlag,
		task.CreatedAt,
		task.UpdatedAt,
		task.ID,
	)
	if execErr != nil {
		return execErr
	}
	return nil
}
