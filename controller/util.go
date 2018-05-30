package controller

import (
	"net/http"
	"encoding/json"
)

// WriteJSON writes body as json to the response writer
func WriteJSON(w http.ResponseWriter, body interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(body)
}


// GetQueryArray returns a slice of strings for a given query key, plus
// a boolean value whether at least one value exists for the given key.
func GetQueryArray(r *http.Request, key string) ([]string, bool) {
	if values, ok := r.URL.Query()[key]; ok && len(values) > 0 {
		return values, true
	}
	return []string{}, false
}