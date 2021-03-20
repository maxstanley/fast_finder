package middleware

import (
	"time"

	"github.com/maxstanley/fast_finder/handler"
	"github.com/maxstanley/fast_finder/logger"
)

func NewLoggerMiddleware(c handler.HandlerContext) {
	// Get the time at the start of the request
	t := time.Now()

	c.Next()

	// Calculate the time between now and the start of the request.
	totalTime := time.Since(t)
	// Log the request information.
	logger.Info("HTTP %s %s - %d %s", c.Method(), c.Path(), c.Status(), totalTime)
}
