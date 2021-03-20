package handler

import (
	"net/http"

	"github.com/maxstanley/fast_finder/datastore"
	"github.com/maxstanley/fast_finder/models"
)

// NewCreateShortcutHandler creates a new shortcut.
func NewCreateShortcutHandler(c HandlerContext) (int, string) {
	// Retrieve shortcut data from the request body.
	var s models.Shortcut
	if err := c.UnmarshalJSONBody(&s); err != nil {
		return http.StatusBadRequest, "Request could not be parsed."
	}

	// Create a shortcut in the datastore using the information from the request
	// body.
	datastore.Connection().Create(&models.Shortcut{
		Keyword: s.Keyword,
		Link:    s.Link,
	})

	return http.StatusCreated, "OK"
}
