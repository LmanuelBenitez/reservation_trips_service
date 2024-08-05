package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

const (
	DEBUG = 0
	INFO  = 1
	WARN  = 2
	ERROR = 3
	FATAL = 4
)

func setLogLevel() {
	var logLevel int
	if _, existsEnv := os.LookupEnv("LOG_LEVEL"); !existsEnv {
		logLevel = DEBUG
	}
	logger := zerolog.New(os.Stdout).Level(zerolog.Level(logLevel))

	logger.Debug().Int("port", 8080).Msg("Get ready...")
}

func Request(method string, status int, requestUri string, time time.Time) string {
	return fmt.Sprintf("[%s][%d][%s][%v]", method, status, requestUri, time)
}

func Info(fileName string, moduleName string, message string) string {
	return fmt.Sprintf("[%s][%s][%s]", fileName, moduleName, message)
}

func Warm(fileName string, moduleName string, message string) string {
	return fmt.Sprintf("[%s][%s][%s]", fileName, moduleName, message)
}

func Error(fileName string, moduleName string, message string) string {
	return fmt.Sprintf("[%s][%s][%s]", fileName, moduleName, message)
}

func Fatal(fileName string, moduleName string, message string) string {
	return fmt.Sprintf("[%s][%s][%s]", fileName, moduleName, message)
}
