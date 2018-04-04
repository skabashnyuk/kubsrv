package server

import (
	"testing"
	"github.com/appleboy/gofight"
	"net/http"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"github.com/ghodss/yaml"
	"github.com/skabashnyuk/kubsrv/model/che.eclipse.org/v1"
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

func TestYaml(t *testing.T) {
	dat, err := ioutil.ReadFile("test-service.yaml")
	check(err)
	var cheService v1.CheService;
	err = yaml.Unmarshal(dat, &cheService)
	check(err)
}

func TestJson(t *testing.T) {
	dat, err := ioutil.ReadFile("test-service.json")
	check(err)
	var cheService v1.CheService
	err = json.Unmarshal(dat, &cheService)
	check(err)
}

func TestJson2(t *testing.T) {
	dat, err := ioutil.ReadFile("test-command.json")
	check(err)
	var cheService v1.CheServiceSpec

	err = json.Unmarshal(dat, &cheService)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
