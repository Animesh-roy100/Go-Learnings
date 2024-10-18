package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// LogEntry represents a single log entry
type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
}

// logEntryPool is a sync.Pool for LogEntry objects
var logEntryPool = sync.Pool{
	New: func() interface{} {
		return &LogEntry{}
	},
}

// WriteLog writes a log entry to a hypothetical storage
func WriteLog(level, message string) error {
	// Get a LogEntry from the pool
	entry := logEntryPool.Get().(*LogEntry)

	// Reset and populate the entry
	entry.Timestamp = time.Now()
	entry.Level = level
	entry.Message = message

	// Convert to JSON (simulating writing to a log file)
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	fmt.Printf("Log written: %s\n", string(data))

	// Return the entry to the pool
	logEntryPool.Put(entry)

	return nil
}

func main() {
	// Simulate writing many log entries
	for i := 0; i < 1000000; i++ {
		WriteLog("INFO", fmt.Sprintf("This is log entry %d", i))
	}
}
