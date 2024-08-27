// 各コントローラーへの処理をまとめ、動作単位にまとめた関数を定義するパッケージ
package usecase

import (
	jquants "app/controller/jquants"
	log "app/controller/log"
	postgres "app/controller/postgres"
	model "app/domain/model"
	"fmt"
	"reflect"
	"time"
)

/*
Jquants API から上場銘柄一覧を取得し、DB に保存する関数
- return) err	エラー
*/
func GetAndSaveStocksInfo() (err error) {
	fmt.Println("EXECUTE GetAndUpdateStocksInfo")

	// 上場銘柄一覧を取得
	stocksNew, err := jquants.GetStocksInfo()
	if err != nil {
		log.Error(err)
		return err
	}

	// 上場銘柄一覧を取得する
	stocksOld, err := postgres.GetStocksInfo()
	if err != nil {
		log.Error(err)
		return err
	}

	// DB から取得した上場銘柄一覧が空の場合は INSERT 空でない場合は UPDATE
	if len(stocksOld) == 0 {
		err = postgres.InsertStocksInfo(stocksNew)
		if err != nil {
			log.Error(err)
			return err
		}
	} else {
		err = postgres.UpdateStocksInfo(stocksNew)
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

	// 分割挿入モードを指定する変数
	isDivisionalInsert := true

	if isDivisionalInsert {
		// 分割挿入モード
		// 上場銘柄一覧の財務情報を取得
		for _, stock := range stocks {
			financials, err := jquants.GetFinancialInfo(stock.Code)
			if err != nil {
				log.Error(err)
				return err
			}

			// 財務情報がない場合はスキップ
			if financials[0].DisclosedDate == nil {
				continue
			}

			// 取得した財務情報を DB に保存
			for _, financial := range financials {
				err = postgres.InsertFinancialInfo(financial)
				if err != nil {
					log.Error(err)
					return err
				}
			}
		}
	} else {
		// 一括挿入モード
		// 全ての財務情報を格納するスライス
		var allFinancials []model.FinancialInfo

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

	// 取得した財務情報を DB の財務情報と比較し、更新
	for _, financial := range yesterdayFinancials {
		// 現在の DB の財務情報を取得
		financialOld, err := postgres.GetFinancialInfo(financial.Code)
		if err != nil {
			log.Error(err)
			return err
		}

		// 構造体の各項目に更新（nil 以外の項目）がある場合は更新
		m := reflect.ValueOf(financialOld)
		merged := reflect.ValueOf(&financial).Elem()
		for i := 0; i < m.NumField(); i++ {
			if merged.Field(i).IsNil() {
				merged.Field(i).Set(m.Field(i))
			} else {
				m.Field(i).Set(merged.Field(i))
			}
		}

		err = postgres.UpdateFinancialInfo(financial)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	for _, financial := range todayFinancials {
		// 現在の DB の財務情報を取得
		financialOld, err := postgres.GetFinancialInfo(financial.Code)
		if err != nil {
			log.Error(err)
			return err
		}

		// 構造体の各項目に更新（nil 以外の項目）がある場合は更新
		m := reflect.ValueOf(financialOld)
		merged := reflect.ValueOf(&financial).Elem()
		for i := 0; i < m.NumField(); i++ {
			if merged.Field(i).IsNil() {
				merged.Field(i).Set(m.Field(i))
			} else {
				m.Field(i).Set(merged.Field(i))
			}
		}

		err = postgres.UpdateFinancialInfo(financial)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}
