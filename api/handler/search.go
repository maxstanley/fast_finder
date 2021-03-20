package handler

import (
	"net/http"

	"github.com/maxstanley/fast_finder/datastore"
	"github.com/maxstanley/fast_finder/models"
)

// NewSearchHandler returns the path pararemeter that is sent with the request.
func NewSearchHandler(c HandlerContext) (int, string) {
	searchParam := c.Param("search")
	if searchParam == "" {
		return http.StatusBadRequest, "no search parameter provided"
	}

	// Get datastore connection.
	db := datastore.Connection()
	var shortcut models.Shortcut
	// Find the associated url for the searched keyword.
	db.First(&shortcut, "keyword = ?", searchParam)

	// If no record was returned, respond with not found.
	if shortcut.Keyword == "" {
		return http.StatusNotFound, ""
	}

	// Return the found Link for the associate keyword.
	return http.StatusTemporaryRedirect, shortcut.Link
}
