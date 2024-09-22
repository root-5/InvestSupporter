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
Jquants API から昨日と今日に更新された財務情報を取得し、DB を更新する関数
- return) err	エラー
*/
func UpdateTodayStatementsInfo() (err error) {
	// fmt.Println("EXECUTE UpdateTodayStatementsInfo")

	// 昨日と今日の日付を取得
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	today := time.Now().Format("2006-01-02")

	// 上場銘柄一覧の財務情報を取得
	yesterdayStatements, err := jquants.FetchStatementsInfo(yesterday)
	if err != nil {
		log.Error(err)
		return err
	}
	todayStatements, err := jquants.FetchStatementsInfo(today)
	if err != nil {
		log.Error(err)
		return err
	}

	// 取得した財務情報を DB に保存
	if len(yesterdayStatements) != 0 {
		result, err := postgres.UpdateStatementsInfo(yesterdayStatements)
		if err != nil {
			log.Error(err)
			return err
		}
		// 影響を受けた行数を確認
		rowsAffected, err := postgres.RowsAffected(result)
		if err != nil {
			log.Error(err)
			return err
		}

		// 影響を受けた行数が0の場合はINSERTを行う
		if rowsAffected == 0 {
			err = postgres.InsertStatementsInfo(yesterdayStatements)
			if err != nil {
				// 銘柄一覧に存在しないのに財務情報には存在する形で Jquants API が返す場合があるため、エラーを無視
				// log.Error(err)
				return err
			}
		}
	}
	if len(todayStatements) != 0 {
		result, err := postgres.UpdateStatementsInfo(todayStatements)
		if err != nil {
			log.Error(err)
			return err
		}
		// 影響を受けた行数を確認
		rowsAffected, err := postgres.RowsAffected(result)
		if err != nil {
			log.Error(err)
			return err
		}

		// 影響を受けた行数が0の場合はINSERTを行う
		if rowsAffected == 0 {
			err = postgres.InsertStatementsInfo(todayStatements)
			if err != nil {
				// 銘柄一覧に存在しないのに財務情報には存在する形で Jquants API が返す場合があるため、エラーを無視
				// log.Error(err)
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
Jquants API から昨日と今日に更新された株価情報を取得し、DB を更新する関数
- return) err	エラー
*/
func UpdateTodayPricesInfo() (err error) {
	// fmt.Println("EXECUTE UpdateTodayPricesInfo")

	// 昨日と今日の日付を取得
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	today := time.Now().Format("2006-01-02")

	// DB から昨日と今日の株価情報を取得
	yesterdayPricesFromDb, err := postgres.GetPricesInfo("", yesterday)
	if err != nil {
		log.Error(err)
		return err
	}
	todayPricesFromDb, err := postgres.GetPricesInfo("", today)
	if err != nil {
		log.Error(err)
		return err
	}

	// 昨日の株価情報がない場合は取得して保存
	if len(yesterdayPricesFromDb) == 0 {
		// 昨日の株価情報を取得
		yesterdayPrices, err := jquants.FetchPricesInfo(yesterday)
		if err != nil {
			log.Error(err)
			return err
		}
		// 取得した株価情報を DB に保存
		if len(yesterdayPrices) != 0 {
			err = postgres.InsertPricesInfo(yesterdayPrices)
			if err != nil {
				log.Error(err)
				return err
			}
		}
	}

	// 今日の株価情報がない場合は取得して保存
	if len(todayPricesFromDb) == 0 {
		todayPrices, err := jquants.FetchPricesInfo(today)
		if err != nil {
			log.Error(err)
			return err
		}
		// 取得した株価情報を DB に保存
		if len(todayPrices) != 0 {
			err = postgres.InsertPricesInfo(todayPrices)
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
	statements, err := postgres.GetStatementInfoAll()
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
