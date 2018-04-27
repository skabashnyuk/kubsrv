package main

import (
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/skabashnyuk/kubsrv/controller"
	"time"
	"github.com/skabashnyuk/kubsrv/storage"
	"flag"
)

func main() {

	storage := storage.Storage{}
	cheRegistryUpdateInterval := flag.Int64("update", 0, "Storage update interval in seconds")
	flag.StringVar(&storage.CheRegistryRepository, "registry", "", "Location of repository on filesystem")
	flag.StringVar(&storage.CheRegistryGithubUrl, "github", "", "Git url of repository to clone")
	flag.Parse()

	if storage.CheRegistryRepository == "" || storage.CheRegistryGithubUrl == "" || *cheRegistryUpdateInterval == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	//periodically update storage with features and services
	go func() {
		storage.EnsureExists()
		if *cheRegistryUpdateInterval > 0 {
			for range time.Tick(time.Second * time.Duration(*cheRegistryUpdateInterval)) {
				storage.UpdateStorage()
			}
		}
	}()

	service := &controller.Service{Storage: &storage}
	feature := &controller.Feature{Storage: &storage}

	router := gin.Default()
	router.GET("/", controller.APIEndpoints)
	router.GET("/service/:name/:version", service.GetService)
	router.GET("/service", service.GetServiceByIdList)
	router.GET("/feature/:name/:version", feature.GetFeature)
	router.GET("/feature", feature.GetFeatureByIdList)
	port := "8080"

	if p := os.Getenv("PORT"); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			port = p
		}
	}

	router.Run(":" + port)

}
