package applog

// Package applog adds support to output leveled logs in Go

import (
	"fmt"
	"time"
)

// LOGTIME : time format
var LOGTIME = "2006-01-02 15:04:05"

//Trace : Print trace string
func Trace(a ...interface{}) {
	a = append([]interface{}{"TRACE:", time.Now().Format(LOGTIME)}, a...)
	fmt.Println(a...)
}

//Tracef : Print formatted trace string
func Tracef(format string, a ...interface{}) {
	format = "TRACE: " + time.Now().Format(LOGTIME) + " " + format + "\033[0m"
	fmt.Printf(format, a...)
}

//Info : Print info string
func Info(a ...interface{}) {
	a = append([]interface{}{"INFO: ", time.Now().Format(LOGTIME)}, a...)
	fmt.Println(a...)
}

//Infof : Print formatted info string
func Infof(format string, a ...interface{}) {
	format = "INFO:  " + time.Now().Format(LOGTIME) + " " + format
	fmt.Printf(format, a...)
}

//Warn : Print warning string - highlight with Yellow
func Warn(a ...interface{}) {
	Warnf("%s\n", a...)
}

//Warnf : Print formatted warning string
func Warnf(format string, a ...interface{}) {
	format = "\033[33m" + "WARN:  " + "\033[0m" + time.Now().Format(LOGTIME) + " " + "\033[33m" + format + "\033[0m"
	fmt.Printf(format, a...)
}

//Error : Print error string
func Error(a ...interface{}) {
	Errorf("%s\n", a...)
}

//Errorf : Print formatted error string - highlight with Red
func Errorf(format string, a ...interface{}) {
	format = "\033[31m" + "ERROR: " + "\033[0m" + time.Now().Format(LOGTIME) + " " + "\033[31m" + format + "\033[0m"
	fmt.Printf(format, a...)
}
