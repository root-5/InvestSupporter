package postgres

import (
	model "app/domain/model"

	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// 型定義
var db *sql.DB
var err error

/* DB の初期化をする関数 */
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
	}
}

/* 上場銘柄テーブルを UPDATE する関数
	- stocks	上場銘柄一覧
	> err		エラー
*/
func UpdateStocksInfo(stocks []model.StocksInfo) (err error) {
	// 上場銘柄テーブルを UPDATE
	for _, stock := range stocks {
		_, err = db.Exec("UPDATE stocks_info SET company_name = $2, company_name_english = $3, sector17_code = $4, sector33_code = $5, scale_category = $6, market_code = $7 WHERE code = $1",
			stock.Code,
			stock.CompanyName,
			stock.CompanyNameEnglish,
			stock.Sector17Code,
			stock.Sector33Code,
			stock.ScaleCategory,
			stock.MarketCode,
		)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

/* 上場銘柄テーブルを取得する関数
	- stocks	上場銘柄一覧
	> err		エラー
*/
func GetStocksInfo() (stocks []model.StocksInfo, err error) {
	// データの取得
	rows, err := db.Query("SELECT * FROM stocks_info")
	if err != nil {
		fmt.Println(err)
		return nil,err
	}

    for rows.Next() {
        var stock model.StocksInfo
        err := rows.Scan(&stock.Code, &stock.CompanyName, &stock.CompanyNameEnglish, &stock.Sector17Code, &stock.Sector33Code, &stock.ScaleCategory, &stock.MarketCode)
        if err != nil {
            return nil, err
        }
        stocks = append(stocks, stock)
    }

    // エラーチェック
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return stocks, nil
}
