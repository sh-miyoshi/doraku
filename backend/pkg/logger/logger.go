package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

func writeLog(level string, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	const layout = "2006/01/02 15:04:05"
	_, fname, line, _ := runtime.Caller(2)
	fname = filepath.Base(fname)

	fmt.Printf("%s %s:%d [%s] %s\n", now.Format(layout), fname, line, level, msg)
}

// Debug method outputs log as DEBUG Level
func Debug(format string, a ...interface{}) {
	writeLog("DEBUG", format, a...)
}

// Info method outputs log as INFO Level
func Info(format string, a ...interface{}) {
	writeLog("INFO", format, a...)
}

// Error method outputs log as ERROR Level
func Error(format string, a ...interface{}) {
	writeLog("ERROR", format, a...)
}
