package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// NewNotFoundHandler returns Fast Finder's current version.
func NewNotFoundHandler(c HandlerContext) (int, string) {
	message := fmt.Sprintf("%s %s Not found!", c.Method(), c.Path())
	resMap := map[string]string{"message": message}
	res, _ := json.Marshal(resMap)

	return http.StatusNotFound, string(res)
}
