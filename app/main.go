package main

import (
	api "app/controller/api"
	"app/controller/jquants"
	"app/controller/log"
	"app/controller/postgres"
	"app/usecase/scheduler"
	"app/usecase/usecase"
	"fmt"
	"time"
)

func main() {
	fmt.Println("")

	// DB の初期化
	fmt.Println("DB の初期化")
	err := postgres.InitDB()
	if err != nil {
		log.Error(err)
	}

	// Jquants の初期化
	fmt.Println("Jquants の初期化")
	jquants.SchedulerStart()
	time.Sleep(3 * time.Second)

	// 財務情報を取得し、長さを確認し、0 の場合は再構築を行う
	financials, err := postgres.GetFinancialInfoAll()
	if err != nil {
		log.Error(err)
		return
	}
	if len(financials) == 0 {
		fmt.Println("財務情報が存在しないため、再構築を行います")
		// 財務情報を全て取得し、DB に保存（15分程度の実行時間が必要）
		err := usecase.GetAndSaveFinancialInfoAll()
		if err != nil {
			log.Error(err)
			return
		}
	}

	// Scheduler の初期化
	fmt.Println("Scheduler の初期化")
	scheduler.SchedulerStart()

	// api の初期化
	fmt.Println("API の初期化")
	api.StartServer()
}
