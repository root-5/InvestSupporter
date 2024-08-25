package main

import (
	initialize "app/usecase/initialize"
	scheduler "app/usecase/scheduler"
	"fmt"
	"time"
)

// Reset モード（新環境での再構築）の場合は true にする
var isResetMode = false

func main() {
	fmt.Println("")

	// 各コントローラーの初期化し、3秒待機
	initialize.Init()
	time.Sleep(3 * time.Second)

	// Scheduler の初期化
	fmt.Println("Scheduler の初期化")
	scheduler.SchedulerStart()

	// Reset モードの場合は Reset 関数を実行
	if isResetMode {
		initialize.Reset()
		return
	}

	// 起動したままにするため、無限ループ
	for {
		time.Sleep(1 * time.Hour)
	}
}
