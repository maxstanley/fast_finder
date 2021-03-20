package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxstanley/fast_finder/handler"
	"github.com/maxstanley/fast_finder/middleware"
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

// POST hanldes requests that are sent with the POST method.
func (r *ginRouter) POST(path string, h handler.Handler) {
	r.engine.POST(path, ginHandlerWrapper(h))
}

// NotFound handles requests that do not have an associated handler.
func (r *ginRouter) NoRoute(h handler.Handler) {
	r.engine.NoRoute(ginHandlerWrapper(h))
}

// Use allows middlewares to be called.
func (r *ginRouter) Use(h middleware.Handler) {
	r.engine.Use(ginMiddlewareWrapper(h))
}

// Handler returns the request handler for the router.
func (r *ginRouter) Handler() http.Handler {
	return r.engine
}

// ginHandlerWrapper wraps the gin.Context so that a custom context can be
// passed to the handlers.
func ginHandlerWrapper(h handler.Handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		requestContextOptions := ginRequestContext(c)
		requestContext := handler.NewHandlerContext(requestContextOptions)
		status, response := h(requestContext)

		// If the response status is a redirect, respond with a redirect.
		if status == http.StatusTemporaryRedirect {
			c.Redirect(status, response)
		} else {
			// Else return the response as string.
			c.String(status, response)
		}
	}
}

// ginMiddlewareWrapper wraps the gin.Context so that a custom context can be
// passed to the middleware handlers.
func ginMiddlewareWrapper(h middleware.Handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		requestContextOptions := ginRequestContext(c)
		requestContext := handler.NewHandlerContext(requestContextOptions)
		h(requestContext)
	}
}

// ginRequestContext converts gin context to handler context.
func ginRequestContext(c *gin.Context) *handler.HandlerContextOptions {
	return &handler.HandlerContextOptions{
		Method:            c.Request.Method,
		Path:              c.Request.URL.Path,
		Status:            c.Writer.Status(),
		Param:             c.Param,
		UnmarshalJSONBody: c.ShouldBindJSON,
		Next:              c.Next,
	}
}
