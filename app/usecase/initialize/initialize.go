// アプリの初期化処理を行うパッケージ
package initialize

import (
	jquants "app/controller/jquants"
	postgres "app/controller/postgres"
	scheduler "app/usecase/scheduler"
	log "app/controller/log"
	"fmt"
)

// 各コントローラーの初期化関数を呼び出す関数
func Init() {
	fmt.Println("Exec Init")

	// DB の初期化
	err :=postgres.InitDB()
	if err != nil {
		log.Error(err)
	}

	// Jquants の初期化
	jquants.SchedulerStart()

	// Scheduler の初期化
	scheduler.SchedulerStart()
}
