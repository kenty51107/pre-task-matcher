package model

import "time"

// Taskのオブジェクト構造
type Task struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignkey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 外部キー
	UserId    uint      `json:"user_id" gorm:"not null"`
}

// 新しいTaskオブジェクトをクライアントに返す
type TaskResponse struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Title    string    `json:"title" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
