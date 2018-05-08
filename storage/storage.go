package storage

import (
	"os"
	"github.com/skabashnyuk/kubsrv/types"
	"strings"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"log"
	"io/ioutil"
	"os/exec"
	"path"
	"fmt"
	"github.com/ghodss/yaml"
)

type Storage struct {
	CheRegistryRepository string
	CheRegistryGithubUrl  string
}

type ItemId struct {
	Name    string
	Version string
}

func (storage *Storage) GetPlugins(Limit int, Offset int) (*[]types.ChePlugin, error) {
	var result []types.ChePlugin

	err := filepath.Walk(storage.CheRegistryRepository, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, "CheMeta.yaml") {

			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			obj := types.ChePlugin{}
			err = yaml.Unmarshal(data, &obj)
			if err != nil {
				return err
			}
			result = append(result, obj)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}

	return &result, nil
}

func (storage *Storage) GetPlugin(Id *ItemId) (*types.ChePlugin, error) {
	name := strings.Replace(Id.Name, ".", string(os.PathSeparator), -1)
	cheServiceFile := filepath.Join(storage.CheRegistryRepository, name, Id.Version, "CheMeta.yaml")
	if _, err := os.Stat(cheServiceFile); os.IsNotExist(err) {
		return nil, err
	}

	if gin.IsDebugging() {
		log.Printf("Requested CheMeta %s", cheServiceFile)
	}

	data, err := ioutil.ReadFile(cheServiceFile)
	if err != nil {
		return nil, err
	}
	obj := types.ChePlugin{}
	err = yaml.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func (storage *Storage) GetCheService(Id *ItemId) (*types.CheService, error) {
	name := strings.Replace(Id.Name, ".", string(os.PathSeparator), -1)
	cheServiceFile := filepath.Join(storage.CheRegistryRepository, name, Id.Version, "CheService.yaml")
	if _, err := os.Stat(cheServiceFile); os.IsNotExist(err) {
		return nil, err
	}

	if gin.IsDebugging() {
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

func (storage *Storage) GetCheFeature(Id *ItemId) (*types.CheFeature, error) {
	name := strings.Replace(Id.Name, ".", string(os.PathSeparator), -1)
	cheServiceFile := filepath.Join(storage.CheRegistryRepository, name, Id.Version, "CheFeature.yaml")
	if _, err := os.Stat(cheServiceFile); os.IsNotExist(err) {
		return nil, err
	}

	if gin.IsDebugging() {
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

func (storage *Storage) UpdateStorage() {

	cmd := exec.Command("git", "pull")
	cmd.Dir = storage.CheRegistryRepository
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Storage update: %s\n", out)
}

func (storage *Storage) CloneStorage() {

	log.Printf("Cloning %s\n", storage.CheRegistryGithubUrl)

	cmd := exec.Command("git", "clone", storage.CheRegistryGithubUrl, ".")
	cmd.Dir = storage.CheRegistryRepository

	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Storage initialized: %s\n", out)
}

func (storage *Storage) EnsureExists() {

	if _, err := os.Stat(path.Join(storage.CheRegistryRepository, ".git")); os.IsNotExist(err) {
		storage.CloneStorage()
	} else {
		log.Print("Git storage setup and ready\n")
	}

}
