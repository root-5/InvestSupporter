package scheduler

import (
	"app/usecase/usecase"
	"time"
)

// 定期実行する関数とその設定をまとめた構造体
var jobs = Jobs{
	{
		Name:        "UpdateStocksInfo",
		Duration:    24 * time.Hour,
		Function:    usecase.UpdateStocksInfo,
		ExecuteFlag: true,
	},
	{
		Name:        "UpdateTodayStatementsInfo",
		Duration:    24 * time.Hour,
		Function:    usecase.UpdateTodayStatementsInfo,
		ExecuteFlag: true,
	},
	{
		Name:        "UpdateTodayPricesInfo",
		Duration:    1 * time.Hour,
		Function:    usecase.UpdateTodayPricesInfo,
		ExecuteFlag: true,
	},
}
