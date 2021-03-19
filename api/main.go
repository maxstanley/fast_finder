package main

import (
	"log"

	"github.com/maxstanley/fast_finder/handler"
	"github.com/maxstanley/fast_finder/logger"
	"github.com/maxstanley/fast_finder/router"
)

func main() {
	// Setup Logger.
	logger.SetLogger(log.Printf)
	logger.SetLogLevel(logger.LevelInfo)

	logger.Info("Starting Fast Finder API.")

	// Create new router.
	r := router.NewGinRouter()

	// Set router routes.
	r.GET("/version", handler.NewVersionHandler)

	// Start router.
	r.Start(":3000")

	logger.Info("Exiting Fast Finder API.")
}
