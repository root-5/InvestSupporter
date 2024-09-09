// 定期実行を行う関数をまとめたパッケージ
package jquants

import (
	"app/controller/log"
	"time"
)

// 型定義
type Job struct {
	Name        string
	Duration    time.Duration
	Function    func() error
	ExecuteFlag bool
}
type Jobs []Job

// 定期実行する関数とその設定をまとめた構造体
var jobs = Jobs{
	{
		Name:        "setIdToken",
		Duration:    24 * time.Hour,
		Function:    setIdToken,
		ExecuteFlag: true,
	},
}

// 定期実行を行う関数
func schedulerExec(jobs Jobs) {
	// wg を使った待機は usecase/scheduler/scheduler.go にあるので、ここでは不要
	for _, job := range jobs {

		// ExecuteFlag が true の場合のみ実行
		if job.ExecuteFlag {
			go func(job Job) {
				for {
					log.Info("定期実行: " + job.Name)
					job.Function()
					time.Sleep(job.Duration)
				}
			}(job)
		}
		// Jobs を確実に上から実行するために1秒待機
		time.Sleep(1 * time.Second)
	}
}

// 定期実行を開始する関数
func SchedulerStart() {
	schedulerExec(jobs)
}
