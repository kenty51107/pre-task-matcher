package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

// User関係の処理に対してリポジトリのインターフェースを定義
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// 以下リポジトリインターフェースの実装

type userRepository struct {
	db *gorm.DB
}

// DBのインスタンスをディペンデンシーインジェクションしている
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
