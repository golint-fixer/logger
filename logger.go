/*Package logger implements levels for logging

To get package:

	go get github.com/baijum/logger

Environment variable "LOG_LEVEL" set the log level.  These are the
possible values: DEBUG, INFO, WARNING, and ERROR

Log level can be set from the code like this:

	logger.SetLevel(logger.WARNING)

To log a message, compare the current level with available levels:

	if logger.Level <= logger.INFO {
		log.Printf("Some log message")
	}

	if logger.Level <= logger.WARNING {
		log.Printf("Some log message")
	}

*/
package logger

import (
	"os"
	"strings"
)

//Available levels for logging
const (
	DEBUG = 1 << iota
	INFO
	WARNING
	ERROR
)

//Level hold the current log level
var Level int8 = INFO

//SetLevel set the log level
func SetLevel(l int8) {
	Level = l
}

func initLogLevel() {
	logLevel := os.Getenv("LOG_LEVEL")
	logLevel = strings.ToUpper(logLevel)
	switch {
	case logLevel == "DEBUG":
		Level = DEBUG
	case logLevel == "INFO":
		Level = INFO
	case logLevel == "WARNING":
		Level = WARNING
	case logLevel == "ERROR":
		Level = ERROR
	}
}

func init() {
	initLogLevel()
}
