package handler

import "net/http"

// NewVersionHandler returns Fast Finder's current version.
func NewVersionHandler(_ *HandlerContext) (int, string) {
	return http.StatusOK, "0.0.1"
}
