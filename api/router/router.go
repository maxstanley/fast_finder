package router

import "github.com/maxstanley/fast_finder/handler"

// Router will allow requests to he handled when requests are sent to specified
// paths.
type Router interface {
	// GET hanldes requests that are sent with the GET method.
	GET(path string, h handler.Handler)
	// Start starts the router listening for requests.
	Start(address string)
}
