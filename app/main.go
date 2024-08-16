package main

import (
	scheduler "app/use-case/scheduler"
	usercase "app/use-case/usecase"
	"fmt"
)

func main() {
	fmt.Println("")

	// DB の初期化
	usercase.InitDB()

	// 定期実行を開始
	scheduler.SchedulerStart()
}
