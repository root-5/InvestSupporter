// PostgreSQL を利用するための関数をまとめたパッケージ
package postgres

import (
	"app/controller/log"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// 型定義
var db *sql.DB

/*
DB の初期化をする関数
  - return) err	エラー
*/
func InitDB() (err error) {
	// 環境変数から接続情報を取得
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Tokyo"

	// DB に接続
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/*
直近操作のあった行の数を取得する関数
  - arg) result		操作結果
  - return) rows	操作のあった行の数
  - return) err		エラー
*/
func RowsAffected(result sql.Result) (rows int64, err error) {
	return result.RowsAffected()
}
