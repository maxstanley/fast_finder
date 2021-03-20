package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Set routes for not found.
	r.NoRoute(handler.NewNotFoundHandler)

	// Creates the HTTP Server.
	server := &http.Server{
		Addr:    ":3000",
		Handler: r.Handler(),
	}

	// Starts the HTTP Server in a go routine so the interrupt signals can be
	// handled.	 
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Info("HTTP Server Error: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	// Wait for a selected signal to interrupt the program.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	signal := <-quit
	logger.Info("%s Signal has been caught.", signal.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the HTTP Server.
	logger.Info("HTTP Server Shutting down.")
	if err := server.Shutdown(ctx); err != nil {
		logger.Info("Server failed to shutdown gracefully: %s", err.Error())
	}

	logger.Info("Exiting Fast Finder API.")
}
