package logger

import (
	"io"
	"sync"
)

type LogLevel int

const (
	defaultNumPrefixes = 5
)

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
	LogLevelDebug
)

const (
	DefaultLogLevel = LogLevelInfo
)

var (
	Error   Logger
	Warning Logger
	Info    Logger
	Debug   Logger
)

var (
	setup sync.Once

	// A mutex used to lock prefix additions, etc., when we need to push on a new
	// prefix.
	mut sync.RWMutex

	// Prefix we're currently on.
	prefixCount int

	// Total depth we allow for here.
	prefixes [5]string
)

func getLoggerWithLevel(requested, needed LogLevel, writer io.Writer) Logger {
	if requested >= needed {
		return NewDefaultLogger(writer, &mut)
	}

	return NewSilentLogger(writer, &mut)
}

func SetupWith(level LogLevel, writer io.Writer) {
	setup.Do(func() {
		Error = getLoggerWithLevel(level, LogLevelError, writer)
		Warning = getLoggerWithLevel(level, LogLevelWarning, writer)
		Info = getLoggerWithLevel(level, LogLevelInfo, writer)
		Debug = getLoggerWithLevel(level, LogLevelDebug, writer)

		// This setups up the prefixes for the first time.
		resetPrefix()
	})
}

func Setup(writer io.Writer) {
	SetupWith(DefaultLogLevel, writer)
}

func resetPrefix() {
	Error.SetPrefix("ERROR ")
	Warning.SetPrefix("WARNING ")
	Info.SetPrefix("INFO ")
	Debug.SetPrefix("DEBUG ")
}

// TODO: Implement these things! I couldn't figure out a way to make this fast
// enough (see benchmark).

// Pushes a new prefix on to make tracing exact function location easier.
//
// **WARNING:** It's not implemented.
func Push(prefix string) {
	mut.Lock()
	defer mut.Unlock()

	// TODO: Implement me.
}

// Rmoves a prefix from the stack of prefixes to make tracing the exact
// location easier.
func Pop() {
	mut.Lock()
	defer mut.Unlock()

	// TODO: Implement me.
}
