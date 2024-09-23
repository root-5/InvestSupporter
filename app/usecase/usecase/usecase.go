// 各コントローラーへの処理をまとめ、動作単位にまとめた関数を定義するパッケージ
package usecase

import (
	"app/controller/jquants"
	"app/controller/log"
	"app/controller/postgres"
	"app/domain/model"
	"fmt"
	"time"
)

/*
Jquants API から上場銘柄一覧を取得し、DB に保存する関数
API と DB の上場銘柄一覧の長さが不一致なら、テーブルを削除した上で INSERT する
API と DB の上場銘柄一覧のデータが不一致なら、テーブルを UPDATE する
- return) err	エラー
*/
func UpdateStocksInfo() (err error) {
	// fmt.Println("EXECUTE UpdateStocksInfo")

	// 上場銘柄一覧を API から取得
	stocksNew, err := jquants.FetchStocksInfo()
	if err != nil {
		log.Error(err)
		return err
	}

	// 上場銘柄一覧を DB から取得
	stocksOld, err := postgres.GetStocksInfo()
	if err != nil {
		log.Error(err)
		return err
	}

	// API と DB の上場銘柄の数が不一致なら、不足分を特定してその分だけ INSERT
	if len(stocksOld) != len(stocksNew) {
		// API と DB の上場銘柄の差分を特定
		var diffStocks []model.StocksInfo
		for _, stockNew := range stocksNew {
			isSame := false
			for _, stockOld := range stocksOld {
				if stockNew.Code == stockOld.Code {
					isSame = true
					break
				}
			}
			if !isSame {
				diffStocks = append(diffStocks, stockNew)
			}
		}

		// 上場銘柄一覧を追加
		err = postgres.InsertStocksInfo(diffStocks)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	// API と DB の上場銘柄データが不一致なら、不足分を特定してその分だけ UPDATE
	var diffStocks []model.StocksInfo
	var isSame bool
	for i := range stocksOld {
		if stocksOld[i] != stocksNew[i] {
			diffStocks = append(diffStocks, stocksNew[i])
			isSame = false
			break
		}
	}
	if !isSame {
		// 上場銘柄一覧を更新
		err = postgres.UpdateStocksInfo(diffStocks)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

/*
Jquants API から全ての財務情報を取得し、DB を一度削除したのち、全て保存する関数
！！！15分程度の実行時間が必要！！！
- return) err	エラー
*/
func FetchAndSaveStatementInfoAll() (err error) {
	// fmt.Println("EXECUTE FetchAndSaveStatementInfoAll")

	// 財務情報テーブルを全て削除
	err = postgres.DeleteStatementInfoAll()
	if err != nil {
		log.Error(err)
		return err
	}

	// 上場銘柄一覧を取得
	stocks, err := postgres.GetStocksInfo()
	if err != nil {
		log.Error(err)
		return err
	}

	// 上場銘柄一覧の財務情報を取得
	for _, stock := range stocks {
		statements, err := jquants.FetchStatementsInfo(stock.Code)
		if err != nil {
			log.Error(err)
			return err
		}

		// 取得した財務情報を DB に保存
		if len(statements) != 0 {
			err = postgres.InsertStatementsInfo(statements)
			if err != nil {
				log.Error(err)
				return err
			}
		}
	}

	return nil
}

/*
財務情報テーブルの最新の開示日と比較し、Jquants API から不足分を更新する関数
- return) err	エラー
*/
func UpdateStatementsInfo() (err error) {
	// fmt.Println("EXECUTE UpdateStatementsInfo")

	// 財務情報テーブルの最新の開示日（例：2024-09-20T00:00:00Z）を取得
	lastDisclosureDate, err := postgres.GetStatementsLatestDisclosedDate()
	if err != nil {
		log.Error(err)
		return err
	}

	// "T" 以降を削除し "2006-01-02" になるように整形
	lastDisclosureDate = lastDisclosureDate[:10]

	// time.Time に変換
	lastDisclosureDateParsed, err := time.Parse("2006-01-02", lastDisclosureDate)
	if err != nil {
		log.Error(err)
		return err
	}

	// 今日の日付と比較し、同じなら更新しない
	if lastDisclosureDateParsed.Format("2006-01-02") == time.Now().Format("2006-01-02") {
		return nil
	}

	// 今日の日付と比較し、日数を算出
	shortageDays := int(time.Since(lastDisclosureDateParsed).Hours()/24) - 1

	// 今日までの不足分の日付をリスト化
	var dates []time.Time
	for i := 1; i <= shortageDays; i++ {
		dates = append(dates, lastDisclosureDateParsed.AddDate(0, 0, i))
	}

	// 財務情報を取得し、DB に保存
	for _, date := range dates {
		// 日付を文字列に変換
		dateStr := date.Format("2006-01-02")

		// 財務情報を取得
		statements, err := jquants.FetchStatementsInfo(dateStr)
		if err != nil {
			log.Error(err)
			return err
		}

		// 取得した財務情報を DB に保存
		if len(statements) != 0 {
			err = postgres.InsertStatementsInfo(statements)
			if err != nil {
				log.Error(err)
				return err
			}
		}
	}

	return nil
}

/*
Jquants API からすべての株価情報を取得し、DB を一度削除したのち、全て保存する関数
！！！一時間半程度の実行時間が必要！！！
- return) err	エラー
*/
func FetchAndSavePriceInfoAll() (err error) {
	// fmt.Println("EXECUTE FetchAndSavePriceInfoAll")

	// 株価情報テーブルを全て削除
	err = postgres.DeletePriceInfoAll()
	if err != nil {
		log.Error(err)
		return err
	}

	// 上場銘柄一覧を取得
	stocks, err := postgres.GetStocksInfo()
	if err != nil {
		log.Error(err)
		return err
	}

	// 株価情報を格納するスライス
	var prices []model.PriceInfo

	for _, stock := range stocks {
		// 株価情報を取得
		prices, err = jquants.FetchPricesInfo(stock.Code)
		if err != nil {
			log.Error(err)
			return err
		}
		// 取得した株価情報を DB に保存
		err = postgres.InsertPricesInfo(prices)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

/*
株価情報テーブルの最新の日付と比較し、Jquants API から不足分を更新する関数
- return) err	エラー
*/
func UpdatePricesInfo() (err error) {
	// fmt.Println("EXECUTE UpdatePricesInfo")

	// 株価情報テーブルの最新の開示日（例：2024-09-20T00:00:00Z）を取得
	lastestDate, err := postgres.GetPricesLatestDate()
	if err != nil {
		log.Error(err)
		return err
	}

	// "T" 以降を削除し "2006-01-02" になるように整形
	lastestDate = lastestDate[:10]

	// time.Time に変換
	lastestDateParsed, err := time.Parse("2006-01-02", lastestDate)
	if err != nil {
		log.Error(err)
		return err
	}

	// 今日の日付と比較し、同じなら更新しない
	if lastestDateParsed.Format("2006-01-02") == time.Now().Format("2006-01-02") {
		return nil
	}

	// 今日の日付と比較し、日数を算出
	shortageDays := int(time.Since(lastestDateParsed).Hours()/24) - 1

	// 今日までの不足分の日付をリスト化
	var dates []time.Time
	for i := 1; i <= shortageDays; i++ {
		dates = append(dates, lastestDateParsed.AddDate(0, 0, i))
	}

	// 株価情報を取得し、DB に保存
	for _, date := range dates {
		// 日付を文字列に変換
		dateStr := date.Format("2006-01-02")

		// 株価情報を取得
		prices, err := jquants.FetchPricesInfo(dateStr)
		if err != nil {
			log.Error(err)
			return err
		}

		// 取得した株価情報を DB に保存
		if len(prices) != 0 {
			err = postgres.InsertPricesInfo(prices)
			if err != nil {
				log.Error(err)
				return err
			}
		}
	}

	return nil
}

/*
DB のデータを確認し、データがない場合はデータを取得し、保存する関数
  - return) エラー
*/
func CheckData() (err error) {
	// 上場銘柄を取得し、長さを確認し、0 の場合は再構築を行う
	stocks, err := postgres.GetStocksInfo()
	if err != nil {
		log.Error(err)
		return err
	}
	if len(stocks) == 0 {
		fmt.Println("上場銘柄が存在しないため、再構築を行います")
		// 上場銘柄を全て取得し、DB に保存
		err := UpdateStocksInfo()
		if err != nil {
			log.Error(err)
			return err
		}
	}

	// 財務情報を取得し、長さを確認し、0 の場合は再構築を行う
	statements, err := postgres.GetTodayStatementInfoAll()
	if err != nil {
		log.Error(err)
		return err
	}
	if len(statements) == 0 {
		fmt.Println("財務情報が存在しないため、再構築を行います")
		// 財務情報を全て取得し、DB に保存（15分程度の実行時間が必要）
		err = FetchAndSaveStatementInfoAll()
		if err != nil {
			log.Error(err)
			return err
		}
	}

	// 株価情報を取得し、長さを確認し、0 の場合は再構築を行う
	prices, err := postgres.GetPricesInfo("72030", "")
	if err != nil {
		log.Error(err)
		return err
	}
	if len(prices) == 0 {
		fmt.Println("株価情報が存在しないため、再構築を行います")
		// 株価情報を全て取得し、DB に保存
		err = FetchAndSavePriceInfoAll()
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

/*
全データを削除し、再構築する関数
  - return) エラー
*/
func RebuildData() (err error) {
	// 財務情報を全て削除
	err = postgres.DeleteStatementInfoAll()
	if err != nil {
		log.Error(err)
		return err
	}
	// 3秒待機
	time.Sleep(3 * time.Second)

	// 上場銘柄を全て削除
	err = postgres.DeleteStocksInfo()
	if err != nil {
		log.Error(err)
		return err
	}
	// 3秒待機
	time.Sleep(3 * time.Second)

	// 上場銘柄を全て取得し、DB に保存
	err = UpdateStocksInfo()
	if err != nil {
		log.Error(err)
		return err
	}
	// 3秒待機
	time.Sleep(3 * time.Second)

	// 財務情報を全て削除し、取得しなおして DB に保存（15分程度の実行時間が必要）
	err = FetchAndSaveStatementInfoAll()
	if err != nil {
		log.Error(err)
		return err
	}
	// 3秒待機
	time.Sleep(3 * time.Second)

	// 株価情報を全て削除し、取得しなおして DB に保存
	err = FetchAndSavePriceInfoAll()
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
