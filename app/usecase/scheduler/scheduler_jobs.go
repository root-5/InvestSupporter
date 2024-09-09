package scheduler

import (
	"app/usecase/usecase"
	"time"
)

// 定期実行する関数とその設定をまとめた構造体
var jobs = Jobs{
	{
		Name:        "FetchAndSaveStocksInfo",
		Duration:    24 * time.Hour,
		Function:    usecase.UpdateStocksInfo,
		ExecuteFlag: true,
	},
	{
		Name:        "FetchAndUpdateFinancialInfoToday",
		Duration:    24 * time.Hour,
		Function:    usecase.UpdateTodayFinancialsInfo,
		ExecuteFlag: true,
	},
	{
		Name:        "FetchAndUpdatePriceInfoToday",
		Duration:    1 * time.Hour,
		Function:    usecase.UpdateTodayPricesInfo,
		ExecuteFlag: true,
	},
}
