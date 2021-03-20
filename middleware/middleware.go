package middleware

import "github.com/maxstanley/fast_finder/handler"

// Handler is the function signature for route handlers.
type Handler func(c handler.HandlerContext)
