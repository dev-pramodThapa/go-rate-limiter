package utils

import (
	"log"
	"os"
	"time"
)

type CustomLogger struct {
	logger *log.Logger
}

func NewLogger() *CustomLogger {
	return &CustomLogger{
		logger: log.New(os.Stdout, "", 0),
	}
}

// Info logs an informational message with a timestamp
func (l *CustomLogger) Info(message string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	l.logger.Printf("[INFO] %s: %s", currentTime, message)
}

// Error logs an error message with a timestamp
func (l *CustomLogger) Error(message string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	l.logger.Printf("[ERROR] %s: %s", currentTime, message)
}

// Request logs the details of an HTTP request
func (l *CustomLogger) Request(method, path string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	l.logger.Printf("[REQUEST] %s: %s %s", currentTime, method, path)
}
