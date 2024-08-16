package log

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func Error(err error) {
	fmt.Println("")
	fmt.Println("========================= Error ============================")

	pc, file, line, ok := runtime.Caller(1)
	if ok {
		// エラーが発生したファイルパス、行数、関数名を表示
		fmt.Printf("ファイル名/ %s\n", filepath.Base(file))
		fmt.Printf("行数/ %d\n", line)
		fmt.Printf("関数名/ %s\n", runtime.FuncForPC(pc).Name())
	}

	// エラーメッセージを表示
	fmt.Printf("エラーメッセージ/ %v\n", err)
	fmt.Println("============================================================")
	fmt.Println("")
}