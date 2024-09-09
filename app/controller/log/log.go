package log

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

func Info(msg string) {
	fmt.Printf("\x1b[34m%s\x1b[0m	%s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func Error(err error) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf(`
======================== Error ============================
%s	%s/%d %s
			%v
============================================================
`,
			time.Now().Format("2006-01-02 15:04:05"),
			filepath.Base(file),
			line,
			runtime.FuncForPC(pc).Name(),
			err,
		)
	}
}
