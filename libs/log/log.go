package log

import (
	"fmt"
	"os"
	"time"
)

const (
	infoSeverity  = "Info"
	errorSeverity = "Error"
	warnSeverity  = "Warning"
	fatalSeverity = "Fatal"
)

// Info is used to print just info/debugging related information to console
func Info(format string, values ...interface{}) {
	print(infoSeverity, format, nil, values...)
}

// Warn is used to print just warn/debugging related information to console
func Warn(format string, values ...interface{}) {
	print(warnSeverity, format, nil, values...)
}

// Error is used to print error with it's exceptions
func Error(format string, err error, values ...interface{}) {
	print(errorSeverity, format, err, values...)
}

// Fatal is used to print error with it's exception and then halts the server execution
func Fatal(format string, err error, values ...interface{}) {
	print(fatalSeverity, format, err, values...)
	os.Exit(1)
}

// print handles the actual printing of the given information
func print(severity string, format string, err error, params ...interface{}) {
	prefix := fmt.Sprintf("%v [%v] -", time.Now().Format(time.RFC3339Nano), severity)
	var posfix = format
	if len(params) > 0 {
		posfix = fmt.Sprintf(posfix, params...)
		if err != nil {
			posfix = fmt.Sprintf("%v with downstream exception: \n%+v", posfix, err)
		}
	}
	fmt.Println(prefix, posfix)
}
