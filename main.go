package main

import (
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/skabashnyuk/kubsrv/controller"
	"time"
	"github.com/skabashnyuk/kubsrv/storage"
	"flag"
	"net/http"
	"log"
	"os/signal"
	"context"
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
	plugin := &controller.Plugin{Storage: &storage}

	router := gin.Default()
	router.GET("/", controller.APIEndpoints)
	router.GET("/plugin/:name/:version", plugin.GetPlugin)
	router.GET("/plugin/", plugin.GetLatestPluginsList)
	router.GET("/service/:name/:version", service.GetService)
	router.GET("/service", service.GetServiceByIdList)
	router.GET("/feature/:name/:version", feature.GetFeature)
	router.GET("/feature", feature.GetFeatureByIdList)
	port := "3000"

	if p := os.Getenv("PORT"); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			port = p
		}
	}


	srv := &http.Server{
		Addr:    ":"+port,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")



}
