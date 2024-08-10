package utils

import (
	"log"
)


func LogFatalError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}