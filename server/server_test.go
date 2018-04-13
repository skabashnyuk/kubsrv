package server

import (
	"testing"
	"github.com/appleboy/gofight"
	"net/http"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestGinHelloWorld(t *testing.T) {
	r := gofight.New()

	r.GET("/").
		SetDebug(true).
		Run(Setup(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

		assert.Equal(t, http.StatusOK, r.Code)

		personMap := make(map[string]string)

		err := json.Unmarshal([]byte(r.Body.String()), &personMap)

		if err != nil {
			panic(err)
		}

		assert.Equal(t, "http://localhost:8080/feature/{name}/{version}", personMap["feature_url"])

	})
}

