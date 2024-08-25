// 各コントローラーへの処理をまとめ、動作単位にまとめた関数を定義するパッケージ
package usecase

import (
	jquants "app/controller/jquants"
	log "app/controller/log"
	postgres "app/controller/postgres"
	model "app/domain/model"
	"fmt"
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

	return nil
}
