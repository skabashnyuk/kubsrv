package storage

import (
	"os"
	"github.com/skabashnyuk/kubsrv/types"
	"strings"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"log"
	"io/ioutil"
	"github.com/ghodss/yaml"
	"os/exec"
)

var cheRegistryRepository = os.Getenv("CHE_REGISTRY_REPOSITORY")

type ItemId struct {
	Name    string
	Version string
}

func GetCheService(Id *ItemId) (*types.CheService, error) {
	name := strings.Replace(Id.Name, ".", string(os.PathSeparator), -1)
	cheServiceFile := filepath.Join(cheRegistryRepository, name, Id.Version, "CheService.yaml")
	if _, err := os.Stat(cheServiceFile); os.IsNotExist(err) {
		return nil, err
	}

	if (gin.IsDebugging()) {
		log.Printf("Requested CheService %s", cheServiceFile)
	}

	data, err := ioutil.ReadFile(cheServiceFile)
	if err != nil {
		return nil, err
	}
	obj := types.CheService{}
	err = yaml.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetCheFeature(Id *ItemId) (*types.CheFeature, error) {
	name := strings.Replace(Id.Name, ".", string(os.PathSeparator), -1)
	cheServiceFile := filepath.Join(cheRegistryRepository, name, Id.Version, "CheFeature.yaml")
	if _, err := os.Stat(cheServiceFile); os.IsNotExist(err) {
		return nil, err
	}

	if (gin.IsDebugging()) {
		log.Printf("Requested CheService %s", cheServiceFile)
	}

	data, err := ioutil.ReadFile(cheServiceFile)
	if err != nil {
		return nil, err
	}
	obj := types.CheFeature{}
	err = yaml.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func UpdateStorage() {
	log.Print("Before pull\n")

	cmd:= exec.Command("git", "pull")
	cmd.Dir = cheRegistryRepository
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Git pull %s\n", out)
}
