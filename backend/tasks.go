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
		log.Printf("Failed to get All Tasks from datastore: %+v", queryErr)
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
		// TODO: フロント側で「完了済みのタスクを表示させる」機能を実装したら下記を改修する
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
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	taskID, err := createTask(task)
	if err != nil {
		log.Printf("Could not insert task: %+v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	task.ID = taskID

	return c.JSON(http.StatusCreated, task)
}

func handlerTaskPut(c echo.Context) error {
	task := new(common.Task)
	if err := c.Bind(task); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := updateTask(task); err != nil {
		log.Printf("Unexpected error occured during Task id:%v updating: ERROR %+v", task.ID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, task)
}

func handlerTaskCheck(c echo.Context) error {
	taskID := c.Param("id")

	task, err := selectTask(taskID)
	if err == sql.ErrNoRows {
		log.Printf("Task id:%v doesn't exist in datastore", taskID)
		return echo.NewHTTPError(http.StatusNotFound, err)
	} else if err != nil {
		log.Printf("Unexpected error occured during Task id:%v row.Scan(): ERROR %+v", taskID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if task.CompleteFlag {
		task.CompleteFlag = false
	} else {
		task.CompleteFlag = true
	}
	if err := updateTask(&task); err != nil {
		log.Printf("Unexpected error occured during Task id:%v CompleteFlag turning %v: ERROR %+v", task.ID, task.CompleteFlag, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, task)
}

func handlerTaskDelete(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"task": "delete"})
}

func selectTask(taskID string) (common.Task, error) {
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

	var task common.Task
	if err := row.Scan(
		&task.ID,
		&task.Name,
		&task.Memo,
		&task.Quadrant,
		&task.CompleteFlag,
		&task.CreatedAt,
		&task.UpdatedAt,
	); err != nil {
		log.Printf("Unexpected error occured during select Task id: %v: ERROR %+v", taskID, err)
		return task, err
	}

	return task, nil
}

func createTask(task *common.Task) (int64, error) {
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
	taskID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Could not get taskID: %+v", err)
		return taskID, err
	}

	return taskID, nil
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
