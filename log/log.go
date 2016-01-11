package log

import (
	"fmt"
	"github.com/toshaf/gobl/cli"
)

type LogLevel int

const (
	LogSilent LogLevel = iota
	LogNormal
	LogVerbose
)

type Logger interface {
	Log(format string, args ...interface{})
	Logv(format string, args ...interface{})
}

type consoleLogger struct {
	level LogLevel
}

func NewConsoleLogger(level LogLevel) Logger {
	return &consoleLogger{level}
}

func (l *consoleLogger) Log(format string, args ...interface{}) {
	if l.level == LogSilent {
		return
	}

	log(format, args...)
}

func (l *consoleLogger) Logv(format string, args ...interface{}) {
	if l.level == LogVerbose {
		log(format, args...)
	}
}

func log(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}

func NewLoggerFromFlags() Logger {
	args := cli.Parse()

	var silent, verbose bool

	for _, o := range args.Options {
		switch o {
		case "-s":
			silent = true
		case "-v":
			verbose = true
		}
	}

	level := LogNormal
	if silent {
		level = LogSilent
	} else if verbose {
		level = LogVerbose
	}

	return NewConsoleLogger(level)
}
