package v1

import (
	"testing"
	"io/ioutil"
	"encoding/json"
	"github.com/ghodss/yaml"
)

func TestYaml(t *testing.T) {
	dat, err := ioutil.ReadFile("test-cheservice.yaml")
	check(err)
	var cheService CheService
	err = yaml.Unmarshal(dat, &cheService)
	check(err)
}


func TestServiceYaml(t *testing.T) {
	dat, err := ioutil.ReadFile("test-service.yaml")
	check(err)
	var object Service
	err = yaml.Unmarshal(dat, &object)
	check(err)
}

func TestJson(t *testing.T) {
	dat, err := ioutil.ReadFile("test-service.json")
	check(err)
	var cheService CheService
	err = json.Unmarshal(dat, &cheService)
	check(err)
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}


