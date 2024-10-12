package logger

import (
	"log"
	"os"
)

type ILogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fileName    string
}

func Constructor(fileName string) ILogger {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	obj := ILogger{
		infoLogger:  log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLogger:  log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		fileName:    fileName,
	}
	return obj

}

func (w *ILogger) Info(logText string) {
	w.infoLogger.Println(logText)
}

func (w *ILogger) Warn(logText string) {
	w.warnLogger.Println(logText)
}

func (w *ILogger) Error(logText string) {
	w.errorLogger.Println(logText)
}
