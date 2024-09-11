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

func Init() {
	log.SetOutput(os.Stdout)
}

func Println(prefix string, format string, v ...interface{}) {
	log.SetPrefix(fmt.Sprintf("[%v] ", prefix))
	log.Println(fmt.Sprintf(format, v...))
}
