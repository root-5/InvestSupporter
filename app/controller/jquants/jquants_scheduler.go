// 定期実行を行う関数をまとめたパッケージ
package jquants

import (
	"fmt"
	"time"
)

// 型定義
type Job struct {
	Name string
	Duration time.Duration
	Function func() error
	ExecuteFlag bool
}
type Jobs []Job

// 定期実行する関数とその設定をまとめた構造体
var jobs = Jobs{
	{
		Name: "SetIdToken",
		Duration: 24 * time.Hour,
		Function: setIdToken,
		ExecuteFlag: true,
	},
}

// 定期実行を行う関数
func schedulerExec(jobs Jobs) {
	// wg を使った待機は usecase/scheduler/scheduler.go にあるので、ここでは不要
	for _, job := range jobs {
		// Jobs を確実に上から実行するために1秒待機
		time.Sleep(1 * time.Second)

		if job.ExecuteFlag {
			go func(job Job) {
				job.Function()
				time.Sleep(job.Duration)
			}(job)
		}
	}
}

// 定期実行を開始する関数
func schedulerStart() {
	fmt.Println("Exec jquant.schedulerStart")
	schedulerExec(jobs)
}
