// 定期実行を行う関数をまとめたパッケージ
package scheduler

import (
	"fmt"
	"sync"
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

// 定期実行を行う関数
func schedulerExec(jobs Jobs) {
	var wg sync.WaitGroup
	for _, job := range jobs {

		// ExecuteFlag が true の場合のみ実行
		if job.ExecuteFlag {
			wg.Add(1)
			go func(job Job) {
				defer wg.Done()
				job.Function()
				time.Sleep(job.Duration)
			}(job)
		}
		// Jobs を確実に上から実行するために1秒待機
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}

// 定期実行を開始する関数
func SchedulerStart() {
	fmt.Println("EXECUTE SchedulerStart")
	schedulerExec(jobs)
}
