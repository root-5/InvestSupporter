package main

import (
	// scheduler "app/usecase/scheduler"
	usecase "app/usecase/usecase"
	test "app/usecase/test"
	"fmt"
)

func main() {
	fmt.Println("")

	// テストモード
	isTestMode := true

	// テストモードの場合はテストを実行
	if isTestMode {
		test.Test()
		return
	}

	// 各コントローラーの初期化
	usecase.Init()

	// 定期実行を開始
	// defer scheduler.SchedulerStart()
}
