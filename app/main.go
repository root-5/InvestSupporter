package main

import (
	initialize "app/usecase/initialize"
	scheduler "app/usecase/scheduler"
	"fmt"
	"time"
)

func main() {
	fmt.Println("")

	// 各コントローラーの初期化
	initialize.Init()

	// コントローラーの初期化を3秒待機
	time.Sleep(3 * time.Second)

	// Scheduler の初期化
	fmt.Println("Scheduler の初期化")
	scheduler.SchedulerStart()
}
