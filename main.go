package main

import (
    "os"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/skabashnyuk/kubsrv/controller"
)

func main() {

    router := gin.Default()
    router.GET("/", controller.APIEndpoints)
    router.GET("/service/:name/:version", controller.GetService)
    port := "8080"

    if p := os.Getenv("PORT"); p != "" {
        if _, err := strconv.Atoi(p); err == nil {
            port = p
        }
    }

    router.Run(":" + port)

}
