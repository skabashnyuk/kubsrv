package main

import (
    "github.com/skabashnyuk/kubsrv/server"
    "os"
    "strconv"
    "github.com/nanobox-io/golang-scribble"
    "fmt"
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


//// GinEngine is gin router.
//func GinEngine() *gin.Engine {
//    r := gin.Default()
//    claim := v1.PersistentVolumeClaim{
//        TypeMeta: metav1.TypeMeta{
//            Kind:       "PersistentVolumeClaim",
//            APIVersion: "v1",
//        },
//        ObjectMeta: metav1.ObjectMeta{
//            Name:      "claim",
//            Namespace: "ns",
//        }}
//    r.GET("/ping", func(c *gin.Context) {
//        c.Render(200, render.GYAML{Data: claim})
//
//    })
//
//    return r
//}
