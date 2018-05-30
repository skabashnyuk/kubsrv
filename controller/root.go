package controller

import (
	"fmt"
	"net/http"
	"os"
	"github.com/julienschmidt/httprouter"
)

var cheRegistryRepository = os.Getenv("CHE_REGISTRY_REPOSITORY")

func APIEndpoints(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	reqScheme := "http"

	if r.TLS != nil {
		reqScheme = "https"
	}

	reqHost := "localhost:8080"

	if r.Host != "" {
		reqHost = r.Host
	}

	baseURL := fmt.Sprintf("%s://%s", reqScheme, reqHost)

	resources := map[string]string{
		"service_url": baseURL + "/service/{name}/{version}",
		"feature_url": baseURL + "/feature/{name}/{version}",
	}
	w.WriteHeader(http.StatusOK)
	WriteJSON(w, resources)
}

// toHTTPError returns a non-specific HTTP error message and status code
// for a given non-nil error value. It's important that toHTTPError does not
// actually return err.Error(), since msg and httpStatus are returned to users,
// and historically Go's ServeContent always returned just "404 Not Found" for
// all errors. We don't want to start leaking information in error messages.
func ToHTTPError(err error) (msg string, httpStatus int) {
	if os.IsNotExist(err) {
		return "404  not found", http.StatusNotFound
	}
	if os.IsPermission(err) {
		return "403 Forbidden", http.StatusForbidden
	}
	// Default:
	return "500 Internal Server Error", http.StatusInternalServerError
}
