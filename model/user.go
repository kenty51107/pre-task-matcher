package model

import "time"

// Userのオブジェクト構造
type User struct {
	// ``で囲まれた部分はタグと呼ばれるもので、jsonの形式でデータを返す際にどのような形式で返すかを指定する
	// gormのタグはデータベースのカラム名を指定する
	ID        uint      `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 新しいUserオブジェクトをクライアントに返す
type UserResponse struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique"`
}
