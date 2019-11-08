package backend

import (
	"encoding/json"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"net/http/httptest"
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
			Quadrant:     2,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "必要タスク2",
			Memo:         "必要メモ2",
			Quadrant:     2,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "必要タスク3",
			Memo:         "必要メモ3",
			Quadrant:     2,
			CompleteFlag: true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "錯覚タスク1",
			Memo:         "錯覚メモ1",
			Quadrant:     3,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "錯覚タスク2",
			Memo:         "錯覚メモ2",
			Quadrant:     3,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "錯覚タスク3",
			Memo:         "錯覚メモ3",
			Quadrant:     3,
			CompleteFlag: true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "無駄タスク1",
			Memo:         "無駄メモ1",
			Quadrant:     4,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "無駄タスク2",
			Memo:         "無駄メモ2",
			Quadrant:     4,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Name:         "無駄タスク3",
			Memo:         "無駄メモ3",
			Quadrant:     4,
			CompleteFlag: true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	InitDB()
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
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	setupTestData()
	defer teardownTestData()

	tests := []struct {
		name       string
		wantCount  int
		wantStatus int
		wantErr    bool
	}{
		{
			name:       "正常系",
			wantCount:  8,
			wantStatus: http.StatusOK,
			wantErr:    false,
		},
	}

	var tasks []common.Task
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handlerAllTasksGet(c); tt.wantErr {
				if err == nil {
					t.Fatalf("%q. wantErr %v, but actual err %+v", tt.name, tt.wantErr, err)
				}
			} else if err != nil {
				t.Fatalf("%q. wantErr %v, but actual err occured %+v", tt.name, tt.wantErr, err)
			} else {
				if tt.wantStatus != rec.Code {
					t.Errorf("%q. Expected status %v, but actual status %v", tt.name, tt.wantStatus, rec.Code)
				}
				data := rec.Body.Bytes()
				json.Unmarshal(data, &tasks)
				if tt.wantCount != len(tasks) {
					t.Errorf("%q. Expected Tasks count %v, but actual Tasks %+v", tt.name, tt.wantCount, len(tasks))
				}
			}
		})
	}
}
