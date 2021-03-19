package handler

// Handler is the function signature for route handlers.
type Handler func(c HandlerContext) (int, string)

// HandlerContext contains the context from the request.
type HandlerContext interface {
	// Method retuns the Method the request was sent with.
	Method() string
	// Path contains the relative path the request was sent to.
	Path() string
}

// HandlerContextOptions is used to populate the HandlerContext.
type HandlerContextOptions struct {
	// Method contains the method the request was sent with.
	Method string
	// Path contains the relative path the request was sent to.
	Path string
}

// handleContext contains the request context.
type handlerContext struct {
	o *HandlerContextOptions
}

// NewHandlerContext creates a new handlerContext from the passed HandlerContextOptions.
func NewHandlerContext(o *HandlerContextOptions) HandlerContext {
	c := &handlerContext{o}
	return c
}

// Method retuns the Method the request was sent with.
func (c *handlerContext) Method() string {
	return c.o.Method
}

// Path contains the relative path the request was sent to.
func (c *handlerContext) Path() string {
	return c.o.Path
}
