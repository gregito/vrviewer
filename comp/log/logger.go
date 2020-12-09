package log

import (
	"log"
	"os"
	"strconv"
)

var enableLogging bool

func init() {
	enable, err := strconv.ParseInt(os.Getenv("DEBUG"), 10, 0)
	if err == nil && enable == 1 {
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
