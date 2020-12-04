package log

import (
	"log"
	"os"
	"strconv"
)

var enableLogging bool

func init() {
	enable, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err == nil && enable {
		enableLogging = true
	} else {
		enableLogging = false
	}
}

func Println(v ...interface{}) {
	if enableLogging {
		log.Println(v)
	}
}

func Printf(format string, v ...interface{}) {
	if enableLogging {
		log.Printf(format, v)
	}
}
