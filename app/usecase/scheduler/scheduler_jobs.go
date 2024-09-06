package scheduler

import (
	"app/usecase/usecase"
	"time"
)

// 定期実行する関数とその設定をまとめた構造体
var jobs = Jobs{
	{
		Name:        "GetAndSaveStocksInfo",
		Duration:    30 * 24 * time.Hour,
		Function:    usecase.GetAndSaveStocksInfo,
		ExecuteFlag: true,
	},
	{
		Name:        "GetAndUpdateFinancialInfoToday",
		Duration:    24 * time.Hour,
		Function:    usecase.GetAndUpdateFinancialInfoToday,
		ExecuteFlag: true,
	},
	{
		Name:        "GetAndUpdatePriceInfoToday",
		Duration:    24 * time.Hour,
		Function:    usecase.GetAndUpdatePriceInfoToday,
		ExecuteFlag: true,
	},
}
