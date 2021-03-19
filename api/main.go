package main

import (
	"log"

	"github.com/maxstanley/fast_finder/logger"
)

func main() {
	logger.SetLogger(log.Printf)
	logger.SetLogLevel(logger.LevelInfo)

	logger.Info("Starting Fast Finder API.")
	logger.Info("Exiting Fast Finder API.")
}
