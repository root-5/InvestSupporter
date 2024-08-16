package scheduler

import (
	usecase "app/use-case/usecase"
	"time"
)

// 定期実行する関数とその設定をまとめた構造体
var jobs = Jobs{
	{
		Name: "SetIdToken",
		Duration: 24 * time.Hour,
		Function: usecase.SetIdToken,
		ExecuteFlag: true,
	},
	{
		Name: "GetAndUpdateStocksInfo",
		Duration: 30 * 24 * time.Hour,
		Function: usecase.GetAndUpdateStocksInfo,
		ExecuteFlag: true,
	},
}
