package main

import (
    "github.com/skabashnyuk/kubsrv/server"
    "os"
    "strconv"
)

func main() {

    s := server.Setup()
    port := "8080"

    if p := os.Getenv("PORT"); p != "" {
        if _, err := strconv.Atoi(p); err == nil {
            port = p
        }
    }

    s.Run(":" + port)

}
