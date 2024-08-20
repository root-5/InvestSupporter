package main

import (
	// scheduler "app/usecase/scheduler"
	usecase "app/usecase/usecase"
	"fmt"
)

func main() {
	fmt.Println("")

	// 各コントローラーの初期化
	usecase.Init()

	// 定期実行を開始
	// defer scheduler.SchedulerStart()
}
