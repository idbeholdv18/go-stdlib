package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Logger struct {
	logFile *os.File
	mu      sync.Mutex
}

func NewLogger(logFilePath string) *Logger {
	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return &Logger{
		logFile: file,
	}
}

func (l *Logger) Close() error {
	if err := l.logFile.Close(); err != nil {
		return err
	}
	return nil
}

func (l *Logger) Write(msg string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	msgWithTime := fmt.Sprintf("%v: %s\n", time.Now().Format(time.RFC3339), msg)
	_, err := l.logFile.Write([]byte(msgWithTime))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	logger := NewLogger("test.log")
	defer func() {
		if err := logger.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := logger.Write(fmt.Sprintf("%s from %s", r.Method, r.URL.Path))
		if err != nil {
			log.Println("log write error:", err)
		}
		json.NewEncoder(w).Encode("response")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
