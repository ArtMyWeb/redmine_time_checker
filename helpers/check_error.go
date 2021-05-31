package helpers

import (
	"log"
	"time"
)

// CheckError show pretty error message
func CheckError(message string, err error) {
	if err != nil {
		log.Fatalf("[%s] %s: %s", message, time.Now().Format("Mon, 02 Jan 2006 15:04:05 "), err.Error())
	}
}
