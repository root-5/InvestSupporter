// PostgreSQL を利用するための関数をまとめたパッケージ
package postgres

import (
	log "app/controller/log"
	model "app/domain/model"

	_ "github.com/lib/pq"
)

/*
上場銘柄テーブルに INSERT する関数
  - arg) stocks	上場銘柄一覧
  - return) err		エラー
*/
func InsertStocksInfo(stocks []model.StocksInfo) (err error) {
	// 上場銘柄テーブルに INSERT
	for _, stock := range stocks {
		_, err = db.Exec("INSERT INTO stocks_info (code, company_name, company_name_english, sector17_code, sector33_code, scale_category, market_code) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			stock.Code,
			stock.CompanyName,
			stock.CompanyNameEnglish,
			stock.Sector17Code,
			stock.Sector33Code,
			stock.ScaleCategory,
			stock.MarketCode,
		)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

/*
上場銘柄テーブルを UPDATE する関数
  - arg) stocks	上場銘柄一覧
  - return) err		エラー
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
			log.Error(err)
			return err
		}
	}

	return nil
}

/*
上場銘柄テーブルを取得する関数
  - return) stocks	上場銘柄一覧
  - return) err		エラー
*/
func GetStocksInfo() (stocks []model.StocksInfo, err error) {
	// データの取得
	rows, err := db.Query("SELECT * FROM stocks_info")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// 取得したデータを格納
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
		log.Error(err)
		return nil, err
	}

	return stocks, nil
}