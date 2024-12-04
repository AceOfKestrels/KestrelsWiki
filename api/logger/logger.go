package logger

import (
	params "api/parameters"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

const (
	INIT = "INIT"
	API  = "API"
)

var LogFile *os.File
var latestLogPath string

func Init() {
	logWriter := initLogFile(params.LogPath)
	log.SetOutput(logWriter)
	gin.DefaultWriter = logWriter
}

func initLogFile(logPath string) io.Writer {
	if len(logPath) == 0 {
		return os.Stdout
	}

	latestLogPath = logPath + "latest.log"
	err := RenameOldLogFile(logPath)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(params.LogPath, 0666)
	if err != nil {
		panic(err)
	}

	LogFile, err = os.OpenFile(latestLogPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	return io.MultiWriter(os.Stdout, LogFile)
}

func RenameOldLogFile(logPath string) error {
	file, err := os.Stat(latestLogPath)
	if err != nil {
		return nil
	}
	return os.Rename(latestLogPath, logPath+file.ModTime().Format("2006-01-02_15-04-05")+".log")
}

func Println(prefix string, format string, v ...interface{}) {
	log.SetPrefix(fmt.Sprintf("[%v] ", prefix))
	log.Println(fmt.Sprintf(format, v...))
}
