package main

import (
	jquants "app/controller/jquants"
	postgres "app/controller/postgres"

	"fmt"

	_ "github.com/lib/pq"
)


func main() {
	fmt.Println("Program started")

	// DB の初期化
	postgres.InitDB()

	// ID トークンをセット
	idToken, err := jquants.SetIdToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 上場銘柄一覧を取得
	stocks, err := jquants.GetStockList(idToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 取得した上場銘柄を DB に保存
	err = postgres.SaveStockList(stocks)
}
