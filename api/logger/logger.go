package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	INIT = "INIT"
	API  = "API"
	DB   = "DB"
)

var logger = log.New(os.Stdout, "", log.LstdFlags)

func Init() {
	log.SetOutput(os.Stdout)
}

func Println(prefix string, format string, v ...interface{}) {
	logger.SetPrefix(fmt.Sprintf("[%v] ", prefix))
	logger.Println(fmt.Sprintf(format, v...))
}
