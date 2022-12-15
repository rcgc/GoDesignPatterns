package main

// TODO: use the "sync" package for the Once API
import (
	"fmt"
	"sync"
)

// MyLogger is the struct we want to make a singleton
type MyLogger struct {
	loglevel int
}

// Log a message using the Logger
func (l *MyLogger) Log(s string){
	fmt.Println(l.loglevel, ":", s)
}

// SetLogLevel sets the Log Level of the Logger
func (l *MyLogger) SetLogLevel(level int){
	l.loglevel = level
}

// the Logger instance
var logger *MyLogger


// TODO: use the sync package to enforce goroutine safety
var once sync.Once

// TODO: the getLoggerInstance function provides global acces to the
// Logger class instance
func getLoggerInstance() *MyLogger{
	once.Do(func() {
		fmt.Println("Creating Logger Instance")
		logger = &MyLogger{}
	})
	fmt.Println("Returning Logger Instance")
	return logger
}