package main

import (
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/skabashnyuk/kubsrv/controller"
	"time"
	"github.com/skabashnyuk/kubsrv/storage"
	"fmt"
)


var cheRegistryGithubUrl = os.Getenv("CHE_REGISTRY_UPDATE_INTERVAL")

func main() {

	//periodically update storage with features and services
	go func() {
		storage.EnsureExists()
		i1, err := strconv.ParseInt(cheRegistryGithubUrl, 10, 64)
		if err == nil {
			fmt.Println(i1)
		}


		for range time.Tick(time.Second * time.Duration(i1)) {
			storage.UpdateStorage()
		}
	}()

	router := gin.Default()
	router.GET("/", controller.APIEndpoints)
	router.GET("/service/:name/:version", controller.GetService)
	router.GET("/service", controller.GetServiceByIdList)
	router.GET("/feature/:name/:version", controller.GetFeature)
	router.GET("/feature", controller.GetFeatureByIdList)
	port := "8080"

	if p := os.Getenv("PORT"); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			port = p
		}
	}

	router.Run(":" + port)

}
