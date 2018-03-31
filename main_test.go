package main

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/appleboy/gofight"
	"github.com/ghodss/yaml"
	"k8s.io/api/core/v1"
)

func TestGinHelloWorld(t *testing.T) {
	r := gofight.New()

	r.GET("/ping").
		SetDebug(true).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

		//assert.Equal(t, "Hello World", r.Body.String())
		assert.Equal(t, http.StatusOK, r.Code)

		var p2 v1.PersistentVolumeClaim
		assert.Nil(t, yaml.Unmarshal(r.Body.Bytes(), &p2))

	})
}
