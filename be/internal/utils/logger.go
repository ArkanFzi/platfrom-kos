package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger wrapper untuk simple logging
type Logger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	warnLog  *log.Logger
	debugLog *log.Logger
}

// NewLogger membuat instance logger baru
func NewLogger() *Logger {
	return &Logger{
		infoLog:  log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile),
		errorLog: log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile),
		warnLog:  log.New(os.Stdout, "[WARN] ", log.LstdFlags|log.Lshortfile),
		debugLog: log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.infoLog.Printf(msg+"\n", args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.errorLog.Printf(msg+"\n", args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.warnLog.Printf(msg+"\n", args...)
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.debugLog.Printf(msg+"\n", args...)
}

// LogRequest logs HTTP request details
func (l *Logger) LogRequest(method, path, clientIP string, startTime time.Time, statusCode int) {
	duration := time.Since(startTime).Milliseconds()
	l.Info("%s %s from %s - Status: %d - Duration: %dms", method, path, clientIP, statusCode, duration)
}

// LogError logs error dengan context
func (l *Logger) LogError(operation string, err error, details ...interface{}) {
	msg := fmt.Sprintf("Operation: %s | Error: %v", operation, err)
	if len(details) > 0 {
		msg += fmt.Sprintf(" | Details: %v", details)
	}
	l.Error(msg)
}

// Global logger instance
var GlobalLogger = NewLogger()
