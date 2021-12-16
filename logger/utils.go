package logger

import (
	"runtime"
	"strconv"
)

func createCursor() string {
	_, file, line, _ := runtime.Caller(5)
	return file + ":" + strconv.Itoa(line)
}
func createFunctionName() string {
	pt, _, _, _ := runtime.Caller(5)
	return runtime.FuncForPC(pt).Name()
}
