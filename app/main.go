package main

import (
	scheduler "app/use-case/scheduler"
	usercase "app/use-case/usecase"
	"fmt"
)

func main() {
	fmt.Println("")

	// 各コントローラーの初期化
	usercase.Init()

	// 定期実行を開始
	defer scheduler.SchedulerStart()
}
