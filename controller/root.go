package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

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
		"services_url": baseURL + "/service",
		"service_url":  baseURL + "/service/{name}/{version}",
		"features_url": baseURL + "/feature",
		"feature_url":  baseURL + "/feature/{name}/{version}",
	}

	c.IndentedJSON(http.StatusOK, resources)
}
