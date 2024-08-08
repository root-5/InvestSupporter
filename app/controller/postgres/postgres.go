package postgres

import (
	model "app/domain/model"

	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

// 型定義
var db *sql.DB
var err error

/* DB の初期化 */
func InitDB() {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Tokyo"
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}// defer db.Close()

	dbTest(0)
}

// デフォルト値を設定する関数
func defaultSmallInt(value string) int {
    if value == "" {
        return 0 // デフォルト値を0に設定
    }
    // 文字列を整数に変換
    intValue, err := strconv.Atoi(value)
    if err != nil {
        return 0 // 変換に失敗した場合もデフォルト値を0に設定
    }
    return intValue
}

func dbTest(num int) {
	// テスト実行フラグの判定
	if num == 0 {
		// テストを実行せず終了
		return
	}

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
	_, err = db.Exec("INSERT INTO jquants (email, pass) VALUES ($1, $2)", "email", "pass")
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

func SaveStockList(stocks []model.StockInfo) error {

	// 上場銘柄テーブルへの挿入
	for _, stock := range stocks {
		// ##############################################################
		// エラーが発生中
		// データ型が一致していないため、データの挿入に失敗している模様
		// pq: invalid input syntax for type smallint: ""
		// [{52530 カバー COVER Corporation 10 5250 - 0113 }]
		// ##############################################################
		_, err = db.Exec("INSERT INTO stocks_info (code, company_name, company_name_english, sector17_code, sector33_code, scale_category, market_code) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			stock.Code,
			stock.CompanyName,
			stock.CompanyNameEnglish,
			defaultSmallInt(stock.Sector17Code),
			defaultSmallInt(stock.Sector33Code),
			stock.ScaleCategory,
			defaultSmallInt(stock.MarketCode),
		)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	// データの取得
	rows, err := db.Query("SELECT * FROM stocks_info")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 取得したデータを表示
	fmt.Println("stocks_info")
	for rows.Next() {
		var code string
		var company_name string
		var company_name_english string
		var sector17_code string
		var sector33_code string
		var scale_category string
		var market_code string
		var margin_code string

		err = rows.Scan(&code, &company_name, &company_name_english, &sector17_code, &sector33_code, &scale_category, &market_code, &margin_code)
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Printf("code: %s, company_name: %s, company_name_english: %s, sector17_code: %s, sector33_code: %s, scale_category: %s, market_code: %s, margin_code: %s\n", code, company_name, company_name_english, sector17_code, sector33_code, scale_category, market_code, margin_code)
	}


	return nil
}
