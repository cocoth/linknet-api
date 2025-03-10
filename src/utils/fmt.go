package utils

import (
	"fmt"
	"time"
)

type LogType string

const (
	DEBUG   LogType = "DEBUG"
	WARN    LogType = "WARN"
	ERROR   LogType = "ERROR"
	INFO    LogType = "INFO"
	SUCCESS LogType = "SUCCESS"
)

var colorMap = map[LogType]string{
	DEBUG:   "\033[35m", // Magenta
	WARN:    "\033[33m", // Yellow
	ERROR:   "\033[31m", // Red
	INFO:    "\033[34m", // Blue
	SUCCESS: "\033[32m", // Green
}

func getTimeForLogFormat() string {
	return time.Now().Format("2025-02-01 15:04:05")
}

func logFmt(logType LogType, message string, funcName ...string) {
	color := colorMap[logType]
	if color == "" {
		color = "\033[0m" // Reset
	}
	currentTime := getTimeForLogFormat()
	functionName := ""
	if len(funcName) > 0 {
		functionName = fmt.Sprintf("funcName: %s", funcName[0])
	}
	logMessage := fmt.Sprintf("\033[2m[%s]\033[0m %s[%s]\033[0m %s: %s", currentTime, color, logType, functionName, message)

	fmt.Println(logMessage)

}

func Logger(msg string, funcName string) {
	logFmt(INFO, msg, funcName)
}

func Success(message string, funcName ...string) {
	logFmt(SUCCESS, message, funcName...)
}

func Error(message string, funcName ...string) {
	logFmt(ERROR, message, funcName...)
}

func Warn(message string, funcName ...string) {
	logFmt(WARN, message, funcName...)
}

func Info(message string, funcName ...string) {
	logFmt(INFO, message, funcName...)
}

func Debug(message string, funcName ...string) {
	logFmt(DEBUG, message, funcName...)
}

func Custom(logType string, message string, funcName ...string) {
	currentTime := getTimeForLogFormat()
	functionName := ""
	if len(funcName) > 0 {
		functionName = fmt.Sprintf("funcName: %s", funcName[0])
	}
	logMessage := fmt.Sprintf("\033[2m[%s]\033[0m \033[36m[%s]\033[0m %s: %s", currentTime, logType, functionName, message)
	fmt.Println(logMessage)
}

func ErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
