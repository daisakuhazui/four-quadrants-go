package backend

import (
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
	"time"

	"github.com/daisakuhazui/four-quadrants-go/common"
)

func setupTestData() error {
	os.Setenv("RUNNING_ENV", "TEST")

	testTasks := []*common.Task{
		{
			Name:         "必須タスク1",
			Memo:         "必須メモ1",
			Quadrant:     1,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "必須タスク2",
			Memo:         "必須メモ2",
			Quadrant:     1,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "必須タスク3",
			Memo:         "必須メモ3",
			Quadrant:     1,
			CompleteFlag: true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "必要タスク1",
			Memo:         "必要メモ1",
			Quadrant:     1,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "必要タスク2",
			Memo:         "必要メモ2",
			Quadrant:     1,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "必要タスク3",
			Memo:         "必要メモ3",
			Quadrant:     1,
			CompleteFlag: true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "錯覚タスク1",
			Memo:         "錯覚メモ1",
			Quadrant:     1,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "錯覚タスク2",
			Memo:         "錯覚メモ2",
			Quadrant:     1,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "錯覚タスク3",
			Memo:         "錯覚メモ3",
			Quadrant:     1,
			CompleteFlag: true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "無駄タスク1",
			Memo:         "無駄メモ1",
			Quadrant:     1,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "無駄タスク2",
			Memo:         "無駄メモ2",
			Quadrant:     1,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "無駄タスク3",
			Memo:         "無駄メモ3",
			Quadrant:     1,
			CompleteFlag: true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	for _, task := range testTasks {
		_, err := createTask(task)
		if err != nil {
			return err
		}
	}

	return nil
}

func teardownTestData() error {
	// open database
	db, openErr := OpenDB()
	if openErr != nil {
		panic(openErr)
	}

	_, err := db.Exec(
		`DELETE FROM TASKS`,
	)
	if err != nil {
		panic(err)
	}

	os.Setenv("RUNNING_ENV", "DEV")

	return nil
}

func Test_handlerAllTasksGet(t *testing.T) {

}
