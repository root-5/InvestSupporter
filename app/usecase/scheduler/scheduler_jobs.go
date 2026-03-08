package scheduler

import (
	"app/usecase/usecase"
	"time"
)

// 定期実行する関数とその設定をまとめたスライス
var jobs = []Job{
	{
		Name:        "UpdateStocksInfo",
		Duration:    6 * time.Hour,
		Function:    usecase.UpdateStocksInfo,
		ExecuteFlag: true,
	},
	{
		Name:        "UpdateStatementsInfo",
		Duration:    6 * time.Hour,
		Function:    usecase.UpdateStatementsInfo,
		ExecuteFlag: true,
	},
	{
		Name:        "UpdatePricesInfo",
		Duration:    6 * time.Hour,
		Function:    usecase.UpdatePricesInfo,
		ExecuteFlag: true,
	},
}
