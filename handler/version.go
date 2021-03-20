package handler

import (
	"encoding/json"
	"net/http"
)

// NewVersionHandler returns Fast Finder's current version.
func NewVersionHandler(_ HandlerContext) (int, string) {
	resMap := map[string]string{"version": "0.0.2"}
	res, _ := json.Marshal(resMap)

	return http.StatusOK, string(res)
}
