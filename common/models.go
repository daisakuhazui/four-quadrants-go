package common

import (
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type User struct {
}

type Task struct {
	ID           int64     `json:"id"`           // テーブル上で割り当てられる一意な値
	Name         string    `json:"name"`         // タスクの名前
	Memo         string    `json:"memo"`         // タスクのメモ書き
	Quadrant     int64     `json:"quadrant"`     // 所属する象限
	CompleteFlag bool      `json:"completeFlag"` // 完了フラグ
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
