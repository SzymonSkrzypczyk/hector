package log

import (
	"log"
	"os"
)

var (
	verbose bool
	logger  = log.New(os.Stdout, "", 0)
)

func SetVerbose(v bool) {
	verbose = v
}

func Info(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

func Debug(format string, v ...interface{}) {
	if verbose {
		logger.Printf("[DEBUG] "+format, v...)
	}
}

func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}
