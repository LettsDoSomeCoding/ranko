package logging

import (
	"log"
	"os"
)

// todo - where should logs go to?
var logger = log.New(os.Stderr, "API Server: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC|log.Llongfile|log.Lmsgprefix)

// GetLogger returns a singleton logger to be shared by all modules
func GetLogger() *log.Logger {
	return logger
}
