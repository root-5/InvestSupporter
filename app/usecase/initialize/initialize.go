// アプリの初期化処理を行うパッケージ
package initialize

import (
	jquants "app/controller/jquants"
	log "app/controller/log"
	postgres "app/controller/postgres"
	usecase "app/usecase/usecase"
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

// コンテナ外部へ永続化されたデータすら無くなった状態からの再構築を行う関数
func Reset() {
	// 財務情報を全て取得し、DB に保存（15分程度の実行時間が必要）
	err := usecase.GetAndSaveFinancialInfoAll()
	if err != nil {
		log.Error(err)
		return
	}
}
