package logs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	infoLogger  = log.New(os.Stdout, "INFO : ", 0)
	warnLogger  = log.New(os.Stdout, "WARN : ", 0)
	errorLogger = log.New(os.Stderr, "ERROR : ", 0)
	debugLogger = log.New(os.Stdout, "DEBUG : ", 0)
)

func formatLogMessage(level, message string, err error, fields ...interface{}) string {
	timestamp := time.Now().Format(time.RFC3339)
	_, file, line, ok := runtime.Caller(2)
	if ok {
		file = filepath.Base(file)
		if err != nil {
			return fmt.Sprintf("%s [%s:%d] %s : %v", timestamp, file, line, fmt.Sprintf(message, fields...), err)
		}
		return fmt.Sprintf("%s [%s:%d] %s", timestamp, file, line, fmt.Sprintf(message, fields...))
	}
	if err != nil {
		return fmt.Sprintf("[%s] %s: %s : %v", timestamp, level, fmt.Sprintf(message, fields...), err)
	}
	return fmt.Sprintf("[%s] %s: %s", timestamp, level, fmt.Sprintf(message, fields...))
}

func Error(err error, message string, fields ...interface{}) error {
	logMessage := formatLogMessage("ERROR", message, err, fields...)
	errorLogger.Println(logMessage)
	return err
}

func Warn(message string, fields ...interface{}) {
	logMessage := formatLogMessage("WARN", message, nil, fields...)
	warnLogger.Println(logMessage)
}

func Info(message string, fields ...interface{}) {
	logMessage := formatLogMessage("INFO", message, nil, fields...)
	infoLogger.Println(logMessage)
}

func Debug(message string, fields ...interface{}) {
	logMessage := formatLogMessage("DEBUG", message, nil, fields...)
	debugLogger.Println(logMessage)
}
