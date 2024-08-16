package main

import (
	usercase "app/use-case/usecase"
	scheduler "app/use-case/scheduler"
)

func main() {
	// DB の初期化
	usercase.InitDB()

	// 定期実行を開始
	scheduler.Start()
}
