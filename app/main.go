package main

import (
	api "app/controller/api"
	"app/controller/jquants"
	"app/controller/log"
	"app/controller/postgres"
	scheduler "app/usecase/scheduler"
	"app/usecase/usecase"
	"fmt"
	"time"
)

// Reset モード（新環境での再構築）の場合は true にする
var isResetMode = true

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

	// Reset モードの場合は Reset 関数を実行
	if isResetMode {
		Reset()
		return
	}

	// Scheduler の初期化
	fmt.Println("Scheduler の初期化")
	scheduler.SchedulerStart()

	// api の初期化
	fmt.Println("API の初期化")
	api.StartServer()
}

// コンテナ外部へ永続化されたデータすら無くなった状態からの再構築を行う関数
func Reset() {
	// 財務情報を全て取得し、DB に保存（15分程度の実行時間が必要）
	err := usecase.GetAndSaveFinancialInfoAll()
	if err != nil {
		log.Error(err)
		return
	}
}
