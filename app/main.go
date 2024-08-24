package main

import (
	initialize "app/usecase/initialize"
	"fmt"
)

func main() {
	fmt.Println("")

	// 各コントローラーの初期化
	initialize.Init()

	// 定期実行を開始
	// defer scheduler.SchedulerStart()
}
