package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// Logger struct with a single instance and file handle
type Logger struct {
	file *os.File
}

var (
	loggerInstance *Logger
	once           sync.Once
)

// GetLoggerInstance returns the singleton logger instance
func GetLoggerInstance() *Logger {
	once.Do(func() {
		file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		loggerInstance = &Logger{file: file}
	})
	return loggerInstance
}

// Log method to write messages to the log file
func (l *Logger) Log(message string) {
	_, err := l.file.WriteString(message + "\n")
	if err != nil {
		log.Fatalf("Failed to write to log file: %v", err)
	}
}

func main() {
	// Simulating multiple parts of the application trying to log messages
	logger1 := GetLoggerInstance()
	logger2 := GetLoggerInstance()

	// Both logger1 and logger2 should point to the same instance
	if logger1 == logger2 {
		fmt.Println("Logger1 and Logger2 are the same instance.")
	}

	// Log some messages
	logger1.Log("This is the first log message.")
	logger2.Log("This is the second log message.")

	fmt.Println("Messages logged to app.log file.")
}
