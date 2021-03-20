package handler

import (
	"net/http"
)

// NewSearchHandler returns the path pararemeter that is sent with the request.
func NewSearchHandler(c HandlerContext) (int, string) {
	searchParam := c.Param("search")
	if searchParam == "" {
		return http.StatusBadRequest, "no search parameter provided"
	}

	return http.StatusOK, searchParam
}
