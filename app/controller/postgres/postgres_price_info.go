// PostgreSQL を利用するための関数をまとめたパッケージ
package postgres

import (
	"app/controller/log"
	"app/domain/model"
	"database/sql"
	"fmt"
)

/*
株価テーブルに INSERT する関数
  - arg) price	株価一覧
  - return) err		エラー
*/
func InsertPricesInfo(price []model.PriceInfo) (err error) {
	// Prepare を利用して SQL 文を実行
	stmt, err := db.Prepare("INSERT INTO price_info (ymd, code, adjustment_open, adjustment_high, adjustment_low, adjustment_close, adjustment_volume) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		log.Error(err)
		return err
	}
	defer stmt.Close()

	// 株価テーブルに INSERT
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
株価テーブルを UPDATE する関数
  - arg) price	株価一覧
  - return) err		エラー
*/
func UpdatePricesInfo(prices []model.PriceInfo) (err error) {
	// Prepare を利用して SQL 文を実行
	stmt, err := db.Prepare("UPDATE price_info SET adjustment_open = $3, adjustment_high = $4, adjustment_low = $5, adjustment_close = $6, adjustment_volume = $7 WHERE ymd = $1 AND code = $2")
	if err != nil {
		log.Error(err)
		return err
	}
	defer stmt.Close()

	// 株価テーブルを UPDATE
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
株価テーブルをすべて取得する関数
  - arg) codeOrDate	銘柄コードまたは日付
  - return) prices	株価一覧
  - return) err		エラー
*/
func GetPricesInfo(code string, date string) (prices []model.PriceInfo, err error) {
	// code と date の値によって SQL 文を変更
	var rows *sql.Rows

	if code == "" && date == "" {
		rows, err = db.Query("SELECT * FROM price_info")
		if err != nil {
			log.Error(err)
			return nil, err
		}
	} else if code != "" && date == "" {
		rows, err = db.Query("SELECT * FROM price_info WHERE code = $1", code)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	} else if code == "" && date != "" {
		rows, err = db.Query("SELECT * FROM price_info WHERE ymd = $1", date)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	} else {
		rows, err = db.Query("SELECT * FROM price_info WHERE code = $1 AND ymd = $2", code, date)
		if err != nil {
			log.Error(err)
			return nil, err
		}
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
財務情報テーブルを DELETE する関数
  - return) err		エラー
*/
func DeletePriceInfoAll() (err error) {
	// 財務情報テーブルを DELETE
	_, err = db.Exec("DELETE FROM price_info")
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
