package logger

import (
	"io"
	"log"
	"sync"
)

// This logger basically just ignores it's function call.
type DefaultLogger struct {
	base *log.Logger
	mut  *sync.RWMutex
}

func (lo *DefaultLogger) Fatal(v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Fatal(v...)
}
func (lo *DefaultLogger) Fatalf(format string, v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Fatalf(format, v...)
}
func (lo *DefaultLogger) Fatalln(v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Fatalln(v...)
}
func (lo *DefaultLogger) Flags() int {
	mut.RLock()
	defer mut.RUnlock()
	return 0
}
func (lo *DefaultLogger) Output(calldepth int, s string) error {
	mut.RLock()
	defer mut.RUnlock()
	return lo.base.Output(calldepth, s)
}
func (lo *DefaultLogger) Panic(v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Panic(v...)
}
func (lo *DefaultLogger) Panicf(format string, v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Panicf(format, v...)
}
func (lo *DefaultLogger) Panicln(v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Panicln(v...)
}

func (lo *DefaultLogger) Prefix() string {
	mut.RLock()
	defer mut.RUnlock()
	return lo.base.Prefix()
}

func (lo *DefaultLogger) Print(v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Print(v...)
}
func (lo *DefaultLogger) Printf(format string, v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Printf(format, v...)
}
func (lo *DefaultLogger) Println(v ...interface{}) {
	mut.RLock()
	defer mut.RUnlock()
	lo.base.Println(v...)
}

func (lo *DefaultLogger) SetFlags(flag int) {
	lo.base.SetFlags(flag)
}

func (lo *DefaultLogger) SetPrefix(prefix string) {
	// NOTE: We don't want to acquire the lock here because it'll be held by
	// upstream processes.
	lo.base.SetPrefix(prefix)
}

func NewDefaultLogger(writer io.Writer, mut *sync.RWMutex) Logger {
	lo := new(DefaultLogger)
	lo.base = log.New(writer, "", log.LstdFlags)
	lo.mut = mut
	return lo
}
