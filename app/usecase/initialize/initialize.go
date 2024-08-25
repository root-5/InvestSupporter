// アプリの初期化処理を行うパッケージ
package initialize

import (
	jquants "app/controller/jquants"
	log "app/controller/log"
	postgres "app/controller/postgres"
	"fmt"
)

// 各コントローラーの初期化関数を呼び出す関数
func Init() {
	// DB の初期化
	fmt.Println("DB の初期化")
	err := postgres.InitDB()
	if err != nil {
		log.Error(err)
	}

	// Jquants の初期化
	fmt.Println("Jquants の初期化")
	jquants.SchedulerStart()
}
