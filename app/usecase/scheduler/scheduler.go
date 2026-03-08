// 定期実行を行う関数をまとめたパッケージ
package scheduler

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

// 定期実行を行う関数
func schedulerExec(jobs []Job) {
	// 各ジョブの次回実行時刻を保持し、同時実行を避けつつ定義順で処理する
	nextExecuteTimes := make([]time.Time, len(jobs))
	for i, job := range jobs {
		if job.ExecuteFlag {
			nextExecuteTimes[i] = time.Now()
		}
	}

	for {
		executed := false
		now := time.Now()

		for i, job := range jobs {
			if !job.ExecuteFlag {
				continue
			}
			if now.Before(nextExecuteTimes[i]) {
				continue
			}

			log.Info("定期実行: " + job.Name)
			err := job.Function()
			if err != nil {
				log.Error(err)
			}

			nextExecuteTimes[i] = time.Now().Add(job.Duration)
			executed = true
		}

		if !executed {
			time.Sleep(1 * time.Second)
		}
	}
}

// 定期実行を開始する関数
func SchedulerStart() {
	// fmt.Println("EXECUTE SchedulerStart")
	schedulerExec(jobs)
}
