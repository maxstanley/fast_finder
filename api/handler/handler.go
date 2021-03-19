package handler

// Handler is the function signature for route handlers.
type Handler func(c *HandlerContext) (int, string)

// HandleContext contains the request context.
type HandlerContext struct{}
