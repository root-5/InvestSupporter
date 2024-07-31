package main

import (
	jquants "app/controller"

	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)


func main() {
	fmt.Println("Program started")

	// 環境変数からメールアドレスとパスワードを取得
	email := os.Getenv("JQUANTS_EMAIL")
	pass := os.Getenv("JQUANTS_PASS")



	// DB の初期化
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Tokyo"
	fmt.Println(dsn)
	
	// DB に接続
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// // テーブル削除
	// _, err = db.Exec("DROP TABLE IF EXISTS jquants")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // テーブルの作成
	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS jquants (id SERIAL PRIMARY KEY, email TEXT, pass TEXT)")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // データの挿入
	// _, err = db.Exec("INSERT INTO jquants (email, pass) VALUES ($1, $2)", email, pass)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // データの取得
	// rows, err := db.Query("SELECT * FROM jquants")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // 取得したデータを表示
	// for rows.Next() {
	// 	var id int

	// 	var email string
	// 	var pass string

	// 	err = rows.Scan(&id, &email, &pass)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Printf("id: %d, email: %s, pass: %s\n", id, email, pass)
	// }
	// defer rows.Close()

	// ID トークンをセット
	idToken, err := jquants.SetIdToken(email, pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = idToken
	// fmt.Printf("ID Token: %s\n", idToken)

	// 上場銘柄一覧を取得
	stocks, err := jquants.GetStockList(idToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stocks)
}
