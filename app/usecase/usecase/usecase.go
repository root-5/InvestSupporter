// 各コントローラーへの処理をまとめ、動作単位にまとめた関数を定義するパッケージ
package usecase

import (
	"app/controller/jquants"
	"app/controller/log"
	"app/controller/postgres"
	"app/domain/model"
	"fmt"
	"sort"
	"time"
)

/*
Jquants API から上場銘柄一覧を取得し、DB に保存する関数
API に存在し DB に存在しない銘柄は INSERT し、既存銘柄の差分は UPDATE する
- return) err	エラー
*/
func UpdateStocksInfo() (err error) {
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

	stocksOldMap := make(map[string]model.StocksInfo, len(stocksOld))
	for _, stockOld := range stocksOld {
		stocksOldMap[stockOld.Code] = stockOld
	}

	var insertStocks []model.StocksInfo
	var updateStocks []model.StocksInfo
	for _, stockNew := range stocksNew {
		stockOld, ok := stocksOldMap[stockNew.Code]
		if !ok {
			insertStocks = append(insertStocks, stockNew)
			continue
		}

		if stockOld != stockNew {
			updateStocks = append(updateStocks, stockNew)
		}
	}

	if len(insertStocks) != 0 {
		err = postgres.InsertStocksInfo(insertStocks)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	if len(updateStocks) != 0 {
		err = postgres.UpdateStocksInfo(updateStocks)
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
func ResetStatementInfoAll() (err error) {
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
		// レート制限を考慮 (Light プラン: 1分間に60回まで)
		time.Sleep(1000 * time.Millisecond)

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
	stocks, err := postgres.GetStocksInfo()
	if err != nil {
		log.Error(err)
		return err
	}

	// stocks_info に存在する銘柄コードの集合を作成する
	stockCodeSet := make(map[string]struct{}, len(stocks))
	for _, stock := range stocks {
		stockCodeSet[stock.Code] = struct{}{}
	}

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

		var filteredStatements []model.StatementInfo
		skippedCount := 0
		for _, statement := range statements {
			if _, ok := stockCodeSet[statement.Code]; !ok {
				skippedCount++
				continue
			}

			filteredStatements = append(filteredStatements, statement)
		}
		if skippedCount != 0 {
			log.Info(fmt.Sprintf("stocks_info に存在しない銘柄コードの財務情報を %d 件スキップしました", skippedCount))
		}

		// 取得した財務情報を DB に保存
		if len(filteredStatements) != 0 {
			err = postgres.InsertStatementsInfo(filteredStatements)
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
func ResetPriceInfoAll() (err error) {
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
		prices, _, err = jquants.FetchPricesInfo(stock.Code)
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
	stocks, err := postgres.GetStocksInfo()
	if err != nil {
		log.Error(err)
		return err
	}

	// stocks_info に存在する銘柄コードの集合を作成する
	stockCodeSet := make(map[string]struct{}, len(stocks))
	for _, stock := range stocks {
		stockCodeSet[stock.Code] = struct{}{}
	}

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
	shortageDays := int(time.Since(lastestDateParsed).Hours() / 24)

	// 今日までの不足分の日付をリスト化
	var dates []time.Time
	for i := 1; i <= shortageDays; i++ {
		dates = append(dates, lastestDateParsed.AddDate(0, 0, i))
	}

	// 株式分割があった銘柄のコードを格納するスライス
	var splitStockCodesAll []string

	// 株価情報を取得し、DB に保存
	for _, date := range dates {
		// 日付を文字列に変換
		dateStr := date.Format("2006-01-02")

		// 株価情報を取得
		prices, splitStockCodes, err := jquants.FetchPricesInfo(dateStr)
		if err != nil {
			log.Error(err)
			return err
		}

		// stocks_info に存在しない銘柄コードは保存対象から除外する
		var filteredPrices []model.PriceInfo
		var filteredSplitStockCodes []string
		skippedCount := 0
		for _, price := range prices {
			if _, ok := stockCodeSet[price.Code]; !ok {
				skippedCount++
				continue
			}

			filteredPrices = append(filteredPrices, price)
		}
		for _, splitStockCode := range splitStockCodes {
			if _, ok := stockCodeSet[splitStockCode]; !ok {
				continue
			}

			filteredSplitStockCodes = append(filteredSplitStockCodes, splitStockCode)
		}
		if skippedCount != 0 {
			log.Info(fmt.Sprintf("stocks_info に存在しない銘柄コードの株価情報を %d 件スキップしました", skippedCount))
		}

		// 取得した株価情報を DB に保存
		if len(filteredPrices) != 0 {
			err = postgres.InsertPricesInfo(filteredPrices)
			if err != nil {
				log.Error(err)
				return err
			}
		}

		splitStockCodesAll = append(splitStockCodesAll, filteredSplitStockCodes...)
	}

	// 株式分割した銘柄の株価情報を削除して再取得
	for _, splitStockCode := range splitStockCodesAll {
		// 株価情報を削除
		err = postgres.DeletePriceInfo(splitStockCode)
		if err != nil {
			log.Error(err)
			return err
		}

		// 株価情報を取得
		prices, _, err := jquants.FetchPricesInfo(splitStockCode)
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
株価情報テーブルから複数のコードを指定して終値だけを600日（2年相当）分まとめて取得する関数
  - arg) codes			銘柄コードのスライス
  - arg) ymd			取得する日付（YYYY-MM-DD形式）
  - return) closePrices	株価情報の二次元スライス
  - return) err			エラー
*/
func GetClosePricesInfo(codes []string, ymd string) (closePrices [][]string, err error) {
	// 株価情報を取得
	prices, err := postgres.GetPricesInfo(codes, ymd)
	if err != nil {
		log.Error(err)
		return closePrices, err
	}

	// 日付ごとに価格情報をマップで整理
	priceMap := make(map[string]map[string]string)
	dateSet := make(map[string]bool)

	for _, price := range prices {
		dateStr := price.Date[:10] // YYYY-MM-DD形式に変換
		dateSet[dateStr] = true

		if priceMap[dateStr] == nil {
			priceMap[dateStr] = make(map[string]string)
		}

		if price.AdjustmentClose.Valid {
			priceMap[dateStr][price.Code] = fmt.Sprintf("%.6f", price.AdjustmentClose.Float64)
		} else {
			priceMap[dateStr][price.Code] = ""
		}
	}

	// 日付を降順でソート
	var dates []string
	for date := range dateSet {
		dates = append(dates, date)
	}
	sort.Slice(dates, func(i, j int) bool {
		return dates[i] > dates[j]
	})

	// CSVデータを構築
	for _, date := range dates {
		row := []string{date}
		for _, code := range codes {
			if price, exists := priceMap[date][code]; exists {
				row = append(row, price)
			} else {
				row = append(row, "")
			}
		}
		closePrices = append(closePrices, row)
	}

	// ヘッダー行（["日付", code1, code2,..]）を頭に追加
	header := append([]string{"日付"}, codes...)
	closePrices = append([][]string{header}, closePrices...)

	return closePrices, nil
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
		err = ResetStatementInfoAll()
		if err != nil {
			log.Error(err)
			return err
		}
	}

	// 株価情報を取得し、長さを確認し、0 の場合は再構築を行う
	prices, err := postgres.GetPricesInfo([]string{"72030"}, "")
	if err != nil {
		log.Error(err)
		return err
	}
	if len(prices) == 0 {
		fmt.Println("株価情報が存在しないため、再構築を行います")
		// 株価情報を全て取得し、DB に保存
		err = ResetPriceInfoAll()
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
func ResetDataAll() (err error) {
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
	err = ResetStatementInfoAll()
	if err != nil {
		log.Error(err)
		return err
	}
	// 3秒待機
	time.Sleep(3 * time.Second)

	// 株価情報を全て削除し、取得しなおして DB に保存
	err = ResetPriceInfoAll()
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
