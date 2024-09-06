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
func GetAndSaveStocksInfo() (err error) {
	fmt.Println("EXECUTE GetAndUpdateStocksInfo")

	// 上場銘柄一覧を API から取得
	stocksNew, err := jquants.GetStocksInfo()
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
func GetAndSaveFinancialInfoAll() (err error) {
	fmt.Println("EXECUTE GetAndSaveFinancialInfoAll")

	// 財務情報テーブルを全て削除
	err = postgres.DeleteFinancialInfoAll()
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

	// 全ての財務情報を格納するスライス
	var allFinancials []model.FinancialInfo

	// 一括挿入か分割挿入かを決める変数
	isDividedInsert := true

	if isDividedInsert {
		// 上場銘柄一覧の財務情報を取得
		for _, stock := range stocks {
			financial, err := jquants.GetFinancialInfo(stock.Code)
			if err != nil {
				log.Error(err)
				return err
			}

			// 取得した財務情報を DB に保存
			err = postgres.InsertFinancialInfo(financial[0])
			if err != nil {
				log.Error(err)
				return err
			}
		}
	} else {
		// 上場銘柄一覧の財務情報を取得
		for _, stock := range stocks {
			financial, err := jquants.GetFinancialInfo(stock.Code)
			if err != nil {
				log.Error(err)
				return err
			}

			// 取得した財務情報をスライスに追加
			allFinancials = append(allFinancials, financial...)
		}

		// 取得した財務情報を DB に保存
		err = postgres.InsertFinancialInfoAll(allFinancials)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

/*
Jquants API から昨日と今日に更新された財務情報を取得し、DB を更新する関数
- return) err	エラー
*/
func GetAndUpdateFinancialInfoToday() (err error) {
	fmt.Println("EXECUTE GetAndUpdateFinancialInfoToday")

	// 昨日と今日の日付を取得
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	today := time.Now().Format("2006-01-02")

	// 上場銘柄一覧の財務情報を取得
	yesterdayFinancials, err := jquants.GetFinancialInfo(yesterday)
	if err != nil {
		log.Error(err)
		return err
	}
	todayFinancials, err := jquants.GetFinancialInfo(today)
	if err != nil {
		log.Error(err)
		return err
	}

	// 取得した財務情報を DB に保存
	if len(yesterdayFinancials) != 0 {
		for _, financial := range yesterdayFinancials {
			result, err := postgres.UpdateFinancialInfo(financial)
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
				err = postgres.InsertFinancialInfo(financial)
				if err != nil {
					fmt.Println(financial)
					log.Error(err)
					return err
				}
			}
		}
	}
	if len(todayFinancials) != 0 {
		for _, financial := range todayFinancials {
			result, err := postgres.UpdateFinancialInfo(financial)
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
				err = postgres.InsertFinancialInfo(financial)
				if err != nil {
					fmt.Println(financial)
					log.Error(err)
					return err
				}
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
		err := GetAndSaveStocksInfo()
		if err != nil {
			log.Error(err)
			return err
		}
	}

	// 財務情報を取得し、長さを確認し、0 の場合は再構築を行う
	financials, err := postgres.GetFinancialInfoAll()
	if err != nil {
		log.Error(err)
		return err
	}
	if len(financials) == 0 {
		fmt.Println("財務情報が存在しないため、再構築を行います")
		// 財務情報を全て取得し、DB に保存（15分程度の実行時間が必要）
		err = GetAndSaveFinancialInfoAll()
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
	err = postgres.DeleteFinancialInfoAll()
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
	err = GetAndSaveStocksInfo()
	if err != nil {
		log.Error(err)
		return err
	}
	// 3秒待機
	time.Sleep(3 * time.Second)

	// 財務情報を全て取得し、DB に保存（15分程度の実行時間が必要）
	err = GetAndSaveFinancialInfoAll()
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
