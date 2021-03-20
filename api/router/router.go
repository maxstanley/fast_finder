package router

import (
	"net/http"

	"github.com/maxstanley/fast_finder/handler"
	"github.com/maxstanley/fast_finder/middleware"
)

// Router will allow requests to he handled when requests are sent to specified
// paths.
type Router interface {
	// GET handles requests that are sent with the GET method.
	GET(path string, h handler.Handler)
	// NotFound handles requests that do not have an associated handler.
	NoRoute(h handler.Handler)

	// Use allows middlewares to be called.
	Use(middleware.Handler)
	// Handler returns the request handler for the router.
	Handler() http.Handler
}
