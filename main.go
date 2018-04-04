package main

import (
    "github.com/skabashnyuk/kubsrv/server"
    "os"
    "strconv"
    "fmt"

    "encoding/json"
)

func main() {
    var jsonBlob = []byte(`
	{
	"Name":"cute",
	"animals":[
		{"name": "Platypus", "order": "Monotremata", "something":"else"},
		{"name": "Quoll",    "order": "Dasyuromorphia", "something":"else"}
	]}`)

    type Animal struct {
        Name  string `json:"name"`
        Order string `json:"order"`
    }

    type Species struct {
        Name    string
        Animals []Animal `json:"animals"`
    }

    var species Species
    err := json.Unmarshal(jsonBlob, &species)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("%+v", species)



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
