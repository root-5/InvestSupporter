// 定期実行を行う関数をまとめたパッケージ
package scheduler

import (
	usecase "app/use-case/usecase"
	"fmt"
	"sync"
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

var jobs = Jobs{
	{
		Name: "SetIdToken",
		Duration: 24 * time.Hour,
		Function: usecase.SetIdToken,
		ExecuteFlag: true,
	},
	{
		Name: "GetAndUpdateStocksInfo",
		Duration: 30 * 24 * time.Hour,
		Function: usecase.GetAndUpdateStocksInfo,
		ExecuteFlag: true,
	},
}

// 定期実行を行う関数
func Run(jobs Jobs) {
	fmt.Println("Run Start")
	var wg sync.WaitGroup
	for _, job := range jobs {
		// Jobs を確実に上から実行するために1秒待機
		time.Sleep(1 * time.Second)

		if job.ExecuteFlag {
			wg.Add(1)
			go func(job Job) {
				defer wg.Done()
				fmt.Println(job.Name)
				job.Function()
				time.Sleep(job.Duration)
			}(job)
		}
	}
	wg.Wait()
	fmt.Println("Run Finish")
}

// 定期実行を開始する関数
func Start() {
	Run(jobs)
}