package logger

import "fmt"

type Logger func(f string, m ...interface{})

var (
	// logger is the log function used to output all log messages.
	logger Logger

	// logLevel is the maximum level that should be logged.
	logLevel int
)

// levels are the available levels that can be logged at.
const (
	LevelInfo = iota
)

var levelMap map[int]string = map[int]string{
	LevelInfo: "INFO",
}

// SetLogger sets the pacakge logger that will be used to output log messages.
func SetLogger(l Logger) {
	logger = l
}

// SetLogLevel sets the pacake log level that will determine which logs are outputted.
func SetLogLevel(level int) {
	logLevel = level
}

// log sets a standard format for all logger functions to call.
func log(level int, f string, a ...interface{}) {
	if level < logLevel {
		return
	}

	levelString := levelMap[level]
	m := fmt.Sprintf(f, a...)
	logger("%s: %s\n", levelString, m)
}

// Info allows log messages to be written at the Info level.
func Info(format string, a ...interface{}) {
	log(LevelInfo, format, a...)
}
