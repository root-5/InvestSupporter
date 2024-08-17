package scheduler

import (
	usecase "app/usecase/usecase"
	"time"
)

// 定期実行する関数とその設定をまとめた構造体
var jobs = Jobs{
	{
		Name: "GetAndUpdateStocksInfo",
		Duration: 30 * 24 * time.Hour,
		Function: usecase.GetAndUpdateStocksInfo,
		ExecuteFlag: true,
	},
}
