package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Cre4T3Tiv3/gocmitra/cmd/gocmitra/cmd"
	"github.com/Cre4T3Tiv3/gocmitra/core/logger"
)

func main() {
	// Initialize logger
	logger.Init()

	// Graceful shutdown on interrupt signals (e.g., Ctrl+C)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logger.Warn(fmt.Sprintf("Caught signal: " + sig.String() + ". Exiting gracefully..."))
		os.Exit(1)
	}()

	// Recover from unexpected panics
	defer func() {
		if r := recover(); r != nil {
			logger.Error(fmt.Sprintf("Unexpected panic: " + r.(string)))
			os.Exit(2)
		}
	}()

	// Run the main CLI command
	if err := cmd.Execute(); err != nil {
		logger.Error(fmt.Sprintf("Execution error: " + err.Error()))
		os.Exit(1)
	}
}
