package main

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/api/core/v1"
    "github.com/gin-gonic/gin"
    "github.com/skabashnyuk/kubsrv/render"
)

//type Product struct {
//	gorm.Model
//	Code string
//	Price uint
//}

func main() {
    claim := v1.PersistentVolumeClaim{
        TypeMeta: metav1.TypeMeta{
            Kind:       "PersistentVolumeClaim",
            APIVersion: "v1",
        },
        ObjectMeta: metav1.ObjectMeta{
            Name:      "claim",
            Namespace: "ns",
        }}

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.Render(200, render.GYAML{Data: claim})

    })
    r.Run() // listen and serve on 0.0.0.0:8080

}
