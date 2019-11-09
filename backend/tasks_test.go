package backend

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/daisakuhazui/four-quadrants-go/common"
)

func setupTestData() error {
	os.Setenv("RUNNING_ENV", "TEST")

	testTasks := []*common.Task{
		{
			ID:           1,
			Name:         "必須タスク1",
			Memo:         "必須メモ1",
			Quadrant:     1,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           2,
			Name:         "必須タスク2",
			Memo:         "必須メモ2",
			Quadrant:     1,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           3,
			Name:         "必須タスク3",
			Memo:         "必須メモ3",
			Quadrant:     1,
			CompleteFlag: true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           4,
			Name:         "必要タスク1",
			Memo:         "必要メモ1",
			Quadrant:     2,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           5,
			Name:         "必要タスク2",
			Memo:         "必要メモ2",
			Quadrant:     2,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           6,
			Name:         "必要タスク3",
			Memo:         "必要メモ3",
			Quadrant:     2,
			CompleteFlag: true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           7,
			Name:         "錯覚タスク1",
			Memo:         "錯覚メモ1",
			Quadrant:     3,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           8,
			Name:         "錯覚タスク2",
			Memo:         "錯覚メモ2",
			Quadrant:     3,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           9,
			Name:         "錯覚タスク3",
			Memo:         "錯覚メモ3",
			Quadrant:     3,
			CompleteFlag: true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           10,
			Name:         "無駄タスク1",
			Memo:         "無駄メモ1",
			Quadrant:     4,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           11,
			Name:         "無駄タスク2",
			Memo:         "無駄メモ2",
			Quadrant:     4,
			CompleteFlag: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           12,
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

func Test_handlerTaskPost(t *testing.T) {
	setupTestData()
	defer teardownTestData()

	tests := []struct {
		name       string
		message    string
		wantTask   common.Task
		wantStatus int
		wantErr    bool
	}{
		{
			name: "正常系",
			message: `{
				"name":"ダミータスク名",
				"memo":"ダミーメモ",
				"quadrant":1,
				"completeFlag":false
			}`,
			wantTask: common.Task{
				Name:         "ダミータスク名",
				Memo:         "ダミーメモ",
				Quadrant:     1,
				CompleteFlag: false,
			},
			wantStatus: http.StatusCreated,
			wantErr:    false,
		},
		{
			name:       "異常系 JSONデータが欠損",
			message:    ``,
			wantTask:   common.Task{},
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			byte := []byte(tt.message)
			buf.Write(byte)
			body := strings.NewReader(buf.String())

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/task", body)
			// TODO: 下記の意味を理解する
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			if err := handlerTaskPost(c); tt.wantErr {
				if err == nil {
					t.Fatalf("%q. wantErr %v, but actual err %+v", tt.name, tt.wantErr, err)
				}
			} else if err != nil {
				t.Fatalf("%q. wantErr %v, but actual err occured %+v", tt.name, tt.wantErr, err)
			} else {
				if tt.wantStatus != rec.Code {
					t.Fatalf("%q. Expected status %v, but actual status %v", tt.name, tt.wantStatus, rec.Code)
				}

				resTask := new(common.Task)
				data := rec.Body.Bytes()
				json.Unmarshal(data, &resTask)

				if tt.wantTask.Name != resTask.Name {
					t.Errorf("%q. Task Name: Expected %v, but actual %v", tt.name, tt.wantTask.Name, resTask.Name)
				}
				if tt.wantTask.Memo != resTask.Memo {
					t.Errorf("%q. Task Memo: Expected %v, but actual %v", tt.name, tt.wantTask.Memo, resTask.Memo)
				}
				if tt.wantTask.Quadrant != resTask.Quadrant {
					t.Errorf("%q. Task Quadrant: Expected %v, but actual %v", tt.name, tt.wantTask.Quadrant, resTask.Quadrant)
				}
				if tt.wantTask.CompleteFlag != resTask.CompleteFlag {
					t.Errorf("%q. Task CompleteFlag: Expected %v, but actual %v", tt.name, tt.wantTask.CompleteFlag, resTask.CompleteFlag)
				}
			}
		})
	}
}

func Test_handlerTaskPut(t *testing.T) {
	setupTestData()
	defer teardownTestData()

	tests := []struct {
		name       string
		message    string
		wantTask   common.Task
		wantStatus int
		wantErr    bool
	}{
		{
			name: "正常系",
			message: `{
				"id":1,
				"name":"updated必須タスク1",
				"memo":"updated必須メモ1",
				"quadrant":2,
				"completeFlag":false
			}`,
			wantTask: common.Task{
				ID:           1,
				Name:         "updated必須タスク1",
				Memo:         "updated必須メモ1",
				Quadrant:     2,
				CompleteFlag: false,
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "異常系 JSONデータが欠損",
			message:    ``,
			wantTask:   common.Task{},
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			byte := []byte(tt.message)
			buf.Write(byte)
			body := strings.NewReader(buf.String())

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/task", body)
			// TODO: 下記の意味を理解する
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			if err := handlerTaskPut(c); tt.wantErr {
				if err == nil {
					t.Fatalf("%q. wantErr %v, but actual err %+v", tt.name, tt.wantErr, err)
				}
			} else if err != nil {
				t.Fatalf("%q. wantErr %v, but actual err occured %+v", tt.name, tt.wantErr, err)
			} else {
				if tt.wantStatus != rec.Code {
					t.Fatalf("%q. Expected status %v, but actual status %v", tt.name, tt.wantStatus, rec.Code)
				}

				resTask := new(common.Task)
				data := rec.Body.Bytes()
				json.Unmarshal(data, &resTask)

				if tt.wantTask.Name != resTask.Name {
					t.Errorf("%q. Task Name: Expected %v, but actual %v", tt.name, tt.wantTask.Name, resTask.Name)
				}
				if tt.wantTask.Memo != resTask.Memo {
					t.Errorf("%q. Task Memo: Expected %v, but actual %v", tt.name, tt.wantTask.Memo, resTask.Memo)
				}
				if tt.wantTask.Quadrant != resTask.Quadrant {
					t.Errorf("%q. Task Quadrant: Expected %v, but actual %v", tt.name, tt.wantTask.Quadrant, resTask.Quadrant)
				}
				if tt.wantTask.CompleteFlag != resTask.CompleteFlag {
					t.Errorf("%q. Task CompleteFlag: Expected %v, but actual %v", tt.name, tt.wantTask.CompleteFlag, resTask.CompleteFlag)
				}
			}
		})

	}
}

func Test_handlerTaskCheck(t *testing.T) {
	setupTestData()
	defer teardownTestData()

	tests := []struct {
		name             string
		paramTaskID      string
		wantCompleteFlag bool
		wantStatus       int
		wantErr          bool
	}{
		{
			name:             "正常系 未完了のタスクを完了する",
			paramTaskID:      "1",
			wantCompleteFlag: true,
			wantStatus:       http.StatusOK,
			wantErr:          false,
		},
		{
			name:             "正常系 完了済のタスクを未完了にする",
			paramTaskID:      "1",
			wantCompleteFlag: false,
			wantStatus:       http.StatusOK,
			wantErr:          false,
		},
		{
			name:             "異常系 指定したタスクが存在しない",
			paramTaskID:      "100",
			wantCompleteFlag: false,
			wantStatus:       http.StatusNotFound,
			wantErr:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/task", nil)
			// TODO: 下記の意味を理解する
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.paramTaskID)

			if err := handlerTaskCheck(c); tt.wantErr {
				if err == nil {
					t.Fatalf("%q. wantErr %v, but actual err %+v", tt.name, tt.wantErr, err)
				}
			} else if err != nil {
				t.Fatalf("%q. wantErr %v, but actual err occured %+v", tt.name, tt.wantErr, err)
			} else {
				if tt.wantStatus != rec.Code {
					t.Fatalf("%q. Expected status %v, but actual status %v", tt.name, tt.wantStatus, rec.Code)
				}

				resTask := new(common.Task)
				data := rec.Body.Bytes()
				json.Unmarshal(data, &resTask)

				if tt.wantCompleteFlag != resTask.CompleteFlag {
					t.Errorf("%q. Task CompleteFlag: Expected %v, but actual %v", tt.name, tt.wantCompleteFlag, resTask.CompleteFlag)
				}
			}
		})
	}
}
