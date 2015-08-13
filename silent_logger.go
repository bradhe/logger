package logger

import (
	"io"
	"log"
	"sync"
)

// This logger basically just ignores it's function call.
type SilentLogger struct {
	base *log.Logger
	mut  *sync.RWMutex
}

func (lo *SilentLogger) Fatal(v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Fatal(v...)
}

func (lo *SilentLogger) Fatalf(format string, v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Fatalf(format, v...)
}

func (lo *SilentLogger) Fatalln(v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Fatalln(v...)
}
func (lo *SilentLogger) Flags() int {
	mut.RLock()
	defer mut.RUnlock()
	return 0
}
func (lo *SilentLogger) Output(calldepth int, s string) error {
	mut.RLock()
	defer mut.RUnlock()
	return lo.base.Output(calldepth, s)
}
func (lo *SilentLogger) Panic(v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Panic(v...)
}
func (lo *SilentLogger) Panicf(format string, v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Panicf(format, v...)
}
func (lo *SilentLogger) Panicln(v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Panicln(v...)
}
func (lo *SilentLogger) Prefix() string {
	mut.RLock()
	defer mut.RUnlock()
	return lo.base.Prefix()
}
func (lo *SilentLogger) SetPrefix(prefix string) {
	// NOTE: We don't want to acquire the lock here because it'll be held by
	// upstream processes.
	lo.base.SetPrefix(prefix)
}

//
// NO-OP functions
//
func (lo *SilentLogger) Print(v ...interface{})                 {}
func (lo *SilentLogger) Printf(format string, v ...interface{}) {}
func (lo *SilentLogger) Println(v ...interface{})               {}
func (lo *SilentLogger) SetFlags(flag int)                      {}

func NewSilentLogger(writer io.Writer, mut *sync.RWMutex) Logger {
	lo := new(SilentLogger)
	lo.base = log.New(writer, "", log.LstdFlags)
	lo.mut = mut
	return lo
}
