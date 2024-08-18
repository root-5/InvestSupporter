// 各コントローラーへの処理をまとめ、動作単位にまとめた関数を定義するパッケージ
package usecase

import (
	jquants "app/controller/jquants"
	log "app/controller/log"
	postgres "app/controller/postgres"
	"fmt"
)

// 各コントローラーの初期化関数を呼び出す関数
func Init() {
	fmt.Println("Exec Init")

	// DB の初期化
	postgres.Init()

	// Jquants の初期化
	jquants.Init()
}

// Jquants API から上場銘柄一覧を取得し、DB に保存する関数
func GetAndSaveStocksInfo() (err error) {
	fmt.Println("Exec GetAndUpdateStocksInfo")

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
