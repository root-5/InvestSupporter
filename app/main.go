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

	// 上場銘柄一覧を取得
	stocks, err := jquants.GetStockList()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 取得した上場銘柄を DB に保存
	err = postgres.UpdateStocksInfo(stocks)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 上場銘柄を取得
	stocks, err = postgres.GetStocksInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", stocks)
}
