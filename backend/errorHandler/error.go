package errorHandler

import (
	"log"
	"os"
)

var (
	WarningLogger = log.New(os.Stderr, "[WARNING]: ", log.Ldate|log.Ltime)
	ErrorLogger   = log.New(os.Stderr, "[ERROR]: ", log.Ldate|log.Ltime)
)

func IsError(err error) bool {
	return err != nil
}
