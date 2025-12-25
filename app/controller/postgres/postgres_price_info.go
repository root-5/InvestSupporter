package postgres

import (
	"app/controller/log"
	"app/domain/model"
	"database/sql"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

/*
株価情報テーブルに INSERT する関数
  - arg) price	株価一覧
  - return) err		エラー
*/
func InsertPricesInfo(price []model.PriceInfo) (err error) {
	// Prepare を利用して SQL 文を実行
	stmt, err := db.Prepare("INSERT INTO prices_info (ymd, code, adjustment_open, adjustment_high, adjustment_low, adjustment_close, adjustment_volume) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		log.Error(err)
		return err
	}
	defer stmt.Close()

	// 株価情報テーブルに INSERT
	for _, stock := range price {
		_, err = stmt.Exec(
			stock.Date,
			stock.Code,
			stock.AdjustmentOpen,
			stock.AdjustmentHigh,
			stock.AdjustmentLow,
			stock.AdjustmentClose,
			stock.AdjustmentVolume,
		)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

/*
株価情報テーブルを UPDATE する関数
  - arg) price	株価一覧
  - return) err		エラー
*/
func UpdatePricesInfo(prices []model.PriceInfo) (err error) {
	// Prepare を利用して SQL 文を実行
	stmt, err := db.Prepare("UPDATE prices_info SET adjustment_open = $3, adjustment_high = $4, adjustment_low = $5, adjustment_close = $6, adjustment_volume = $7 WHERE ymd = $1 AND code = $2")
	if err != nil {
		log.Error(err)
		return err
	}
	defer stmt.Close()

	// 株価情報テーブルを UPDATE
	for _, price := range prices {
		_, err = stmt.Exec(
			price.Date,
			price.Code,
			price.AdjustmentOpen,
			price.AdjustmentHigh,
			price.AdjustmentLow,
			price.AdjustmentClose,
			price.AdjustmentVolume,
		)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

/*
株価情報テーブルを取得する関数
  - arg) ymd		日付
  - arg) code		コード
  - return) prices	株価一覧
  - return) err		エラー
*/
func GetPricesInfo(codes []string, ymd string) (prices []model.PriceInfo, err error) {
	// ymd を ~ で分割
	ymdRange := strings.Split(ymd, "~")
	isDateRange := len(ymdRange) == 2

	// code と ymd の値によって SQL 文を変更
	var rows *sql.Rows
	if len(codes) == 0 && ymd == "" {
		return nil, fmt.Errorf("codes and ymd are empty")
	} else if len(codes) > 0 && ymd == "" {
		// 銘柄指定あり、日付指定なし
		rows, err = db.Query("SELECT * FROM prices_info WHERE code = ANY($1) ORDER BY ymd DESC", pq.Array(codes))
	} else if len(codes) == 0 && ymd != "" {
		if isDateRange {
			// 銘柄指定なし、日付範囲指定あり
			rows, err = db.Query("SELECT * FROM prices_info WHERE ymd BETWEEN $1 AND $2 ORDER BY ymd, code ASC", ymdRange[0], ymdRange[1])
		} else {
			// 銘柄指定なし、単一日付指定あり
			rows, err = db.Query("SELECT * FROM prices_info WHERE ymd = (SELECT MAX(ymd) FROM prices_info WHERE ymd <= $1) ORDER BY code ASC", ymd)
		}
	} else { // len(codes) > 0 && ymd != ""
		if isDateRange {
			// 銘柄指定あり、日付範囲指定あり
			rows, err = db.Query("SELECT * FROM prices_info WHERE code = ANY($1) AND ymd BETWEEN $2 AND $3 ORDER BY ymd, code ASC", pq.Array(codes), ymdRange[0], ymdRange[1])
		} else {
			// 銘柄指定あり、単一日付指定あり
			rows, err = db.Query("SELECT * FROM prices_info WHERE code = ANY($1) AND ymd = (SELECT MAX(ymd) FROM prices_info WHERE ymd <= $2) ORDER BY code ASC", pq.Array(codes), ymd)
		}
	}

	if err != nil {
		log.Error(err)
		return nil, err
	}

	// 取得したデータを格納
	for rows.Next() {
		var price model.PriceInfo
		err := rows.Scan(
			&price.Date,
			&price.Code,
			&price.AdjustmentOpen,
			&price.AdjustmentHigh,
			&price.AdjustmentLow,
			&price.AdjustmentClose,
			&price.AdjustmentVolume,
		)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		prices = append(prices, price)
	}

	return prices, nil
}

/*
銘柄コードを指定して、同コードを持つレコードを削除する関数
  - arg) code		銘柄コード
  - return) err		エラー
*/
func DeletePriceInfo(code string) (err error) {
	// 銘柄コードを指定して、同コードを持つレコードを削除
	_, err = db.Exec("DELETE FROM prices_info WHERE code = $1", code)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/*
株価情報テーブルを DELETE する関数
  - return) err		エラー
*/
func DeletePriceInfoAll() (err error) {
	// 財務情報テーブルを DELETE
	_, err = db.Exec("DELETE FROM prices_info")
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/*
株価情報テーブルの最新の日付を取得する関数
  - return) date	最新の日付（例：2024-09-20T00:00:00Z）
  - return) err		エラー
*/
func GetPricesLatestDate() (date string, err error) {
	// データの取得
	err = db.QueryRow("SELECT MAX(ymd) FROM prices_info").Scan(&date)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return date, nil
}
