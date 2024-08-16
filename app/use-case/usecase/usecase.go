// 各コントローラーへの処理をまとめ、動作単位にまとめた関数を定義するパッケージ
package usecase

import (
	jquants "app/controller/jquants"
	postgres "app/controller/postgres"
	"fmt"
)

// 各コントローラーの初期化関数を呼び出す関数
func Init() {
	// DB の初期化
	postgres.Init()

	// Jquants の初期化
	jquants.Init()
}

// Jquants API から上場銘柄一覧を取得し、DB に保存する関数
func GetAndUpdateStocksInfo() (err error) {
	fmt.Println(">> GetAndUpdateStocksInfo")

	// 上場銘柄一覧を取得
	stocks, err := jquants.GetStocksInfo()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 取得した上場銘柄を DB に保存
	err = postgres.UpdateStocksInfo(stocks)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
