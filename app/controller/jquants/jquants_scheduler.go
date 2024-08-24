// 定期実行を行う関数をまとめたパッケージ
package jquants

import (
	log "app/controller/log"
	"fmt"
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
func schedulerExec(jobs Jobs) (err error) {
	errChan := make(chan error, len(jobs))

	// wg を使った待機は usecase/scheduler/scheduler.go にあるので、ここでは不要
	for _, job := range jobs {

		if job.ExecuteFlag {
			go func(job Job) {
				err = job.Function()
				if err != nil {
					log.Error(err)
					errChan <- err
					return
				}
				time.Sleep(job.Duration)
				errChan <- nil
			}(job)
		} else {
			errChan <- nil
		}
		// Jobs を確実に上から実行するために1秒待機
		time.Sleep(1 * time.Second)
	}
	for range jobs {
		if err = <-errChan; err != nil {
			return err
		}
	}

	return nil
}

// 定期実行を開始する関数
func SchedulerStart() {
	fmt.Println("Exec jquant.schedulerStart")
	schedulerExec(jobs)
}
