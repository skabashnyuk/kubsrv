package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"os"
)

var cheRegistryRepository  = os.Getenv("CHE_REGISTRY_REPOSITORY")

func APIEndpoints(c *gin.Context) {
	reqScheme := "http"

	if c.Request.TLS != nil {
		reqScheme = "https"
	}

	reqHost := "localhost:8080"

	if c.Request.Host != "" {
		reqHost = c.Request.Host
	}

	baseURL := fmt.Sprintf("%s://%s", reqScheme, reqHost)

	resources := map[string]string{
		"service_url":  baseURL + "/service/{name}/{version}",
		"feature_url":  baseURL + "/feature/{name}/{version}",
	}

	c.IndentedJSON(http.StatusOK, resources)
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