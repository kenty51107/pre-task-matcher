package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBの起動時にgormパッケージの中で定義されているDBという構造体のアドレスが返ってくる
func NewDB() *gorm.DB {
	// 環境変数の読み込み
	if os.Getenv("GO_ENV") == "dev" {
		// .envファイルの読み込み
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	// DBに接続するためのURLを作成する
	// Sprintfはフォーマットを指定して文字列を作成する関数
	// os.Getenvは環境変数の値を取得する関数
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	// &gorm.Config{}はgormの設定を指定するためのもの
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected")
	// dbは*gorm.DB型の変数なので、そのまま返す
	return db
}

// DBをクローズする関数
func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
