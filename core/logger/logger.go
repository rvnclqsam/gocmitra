package logger

import (
	"fmt"
	"os"
	"time"
)

var logFile *os.File

func Init() {
	var err error
	logFile, err = os.OpenFile(".gocmitra.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[WARN] Could not open log file: %v\n", err)
		return
	}
}

func Info(msg string) {
	write("[INFO] ", msg)
	fmt.Fprintf(os.Stderr, "[ℹ️] %s\n", msg)
}

func Success(msg string) {
	write("[SUCCESS] ", msg)
	// Add a newline for spacing after success messages, to separate narrative from output
	fmt.Fprintf(os.Stderr, "[✅] %s\n\n", msg)
}

func Warn(msg string) {
	write("[WARN] ", msg)
	fmt.Fprintf(os.Stderr, "[⚠️] %s\n", msg)
}

func Error(msg string) {
	write("[ERROR] ", msg)
	fmt.Fprintf(os.Stderr, "[❌] %s\n", msg)
}

func write(prefix, msg string) {
	if logFile == nil {
		return
	}
	timestamp := time.Now().Format(time.RFC3339)
	_, _ = logFile.WriteString(fmt.Sprintf("%s %s%s\n", timestamp, prefix, msg))
}
