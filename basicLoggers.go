package basicLoggers

import (
	"fmt"
	"log"
	"os"
)

// Returns two loggers for standard formatting of messages and errors
// messageLogger prints given message on stdout in form: 'appName: message'.
// errorLogger prints given message on stderr in form: 'appName: message'.
func GetLoggers(appName string) (messageLogger *log.Logger, errorLogger *log.Logger) {
	messageLogger = log.New(os.Stdout, fmt.Sprintf("%s: ", appName), 0)
	errorLogger = log.New(os.Stderr, fmt.Sprintf("%s: ", appName), 0)

	return
}
