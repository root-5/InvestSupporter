package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// DB の初期化
var dsn = "host=" + os.Getenv("POSTGRES_HOST") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=" + os.Getenv("POSTGRES_PORT") + " sslmode=disable TimeZone=Asia/Tokyo"
fmt.Println(dsn)

// DB に接続
db, err := sql.Open("postgres", dsn)
if err != nil {
	fmt.Println(err)
	return
}
defer db.Close()

func test(){
	// テーブル削除
	_, err = db.Exec("DROP TABLE IF EXISTS jquants")
	if err != nil {
		fmt.Println(err)
		return
	}

	// テーブルの作成
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS jquants (id SERIAL PRIMARY KEY, email TEXT, pass TEXT)")
	if err != nil {
		fmt.Println(err)
		return
	}

	// データの挿入
	_, err = db.Exec("INSERT INTO jquants (email, pass) VALUES ($1, $2)", email, pass)
	if err != nil {
		fmt.Println(err)
		return
	}

	// データの取得
	rows, err := db.Query("SELECT * FROM jquants")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 取得したデータを表示
	for rows.Next() {
		var id int

		var email string
		var pass string

		err = rows.Scan(&id, &email, &pass)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("id: %d, email: %s, pass: %s\n", id, email, pass)
	}
	defer rows.Close()
}
