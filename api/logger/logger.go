package logger

import (
	params "api/parameters"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

const (
	INIT = "INIT"
	API  = "API"
)

const (
	logFileFormat      = ".log"
	logTimestampFormat = "2006-01-02_15-04-05"
)

var LogFile *os.File
var logFilePath string

func Init() {
	logWriter := initLogFile(params.LogPath)
	log.SetOutput(logWriter)
	gin.DefaultWriter = logWriter

	go syncLogFile()
}

func initLogFile(logPath string) io.Writer {
	if len(logPath) == 0 {
		return os.Stdout
	}

	logFilePath = logPath + time.Now().Format(logTimestampFormat) + logFileFormat

	err := os.MkdirAll(params.LogPath, 0666)
	if err != nil {
		log.Fatal(err)
	}

	LogFile, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return io.MultiWriter(os.Stdout, LogFile)
}

func syncLogFile() {
	for {
		time.Sleep(time.Second)
		if LogFile != nil {
			_ = LogFile.Sync()
		}
	}
}

func Println(prefix string, format string, v ...interface{}) {
	log.SetPrefix(fmt.Sprintf("[%v] ", prefix))
	log.Println(fmt.Sprintf(format, v...))
}
