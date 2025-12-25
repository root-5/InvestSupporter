package main

import (
	"app/controller/api"
	"app/controller/log"
	"app/controller/postgres"
	"app/usecase/scheduler"
	"app/usecase/usecase"
	"time"
)

func main() {
	log.Info("")

	// DB の初期化
	log.Info("DB の初期化")
	err := postgres.InitDB()
	if err != nil {
		log.Error(err)
	}
	time.Sleep(1 * time.Second)

	// DB のデータを確認し、問題がある場合は修正
	err = usecase.CheckData()
	if err != nil {
		log.Error(err)
		return
	}

	// Scheduler の初期化
	log.Info("Scheduler の初期化")
	scheduler.SchedulerStart()

	// api の初期化
	log.Info("API の初期化")
	api.StartServer()
}
