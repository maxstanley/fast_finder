package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maxstanley/fast_finder/handler"
)

// ginRouter implements the Router interface for gin-gonic.
type ginRouter struct {
	engine *gin.Engine
}

// NewGinRouter returns a new instance of Router for gin.
func NewGinRouter() Router {
	g := &ginRouter{}

	// Sets gin to release mode, this remove all debug statements.
	gin.SetMode(gin.ReleaseMode)
	// Create a new instance of the gin router.
	r := gin.New()
	// With recovery middleware.
	r.Use(gin.Recovery())
	g.engine = r

	return g
}

// GET hanldes requests that are sent with the GET method.
func (r *ginRouter) GET(path string, h handler.Handler) {
	r.engine.GET(path, ginHandlerWrapper(h))
}

// Start starts the router listening for requests.
func (r *ginRouter) Start(address string) {
	r.engine.Run(address)
}

// ginHandlerWrapper wraps the gin.Context so that a custom context can be
// passed to the handlers.
func ginHandlerWrapper(h handler.Handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		requestContext := &handler.HandlerContext{}
		status, response := h(requestContext)
		c.String(status, response)
	}
}
