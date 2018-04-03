package server


import (

	"github.com/skabashnyuk/kubsrv/router"
	"github.com/skabashnyuk/kubsrv/storage"
	"github.com/gin-gonic/gin"
	"github.com/nanobox-io/golang-scribble"
	"fmt"
)

func Setup() *gin.Engine {
	r := gin.Default()
	dir := "./"

	db, err := scribble.New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	r.Use(storage.SetDBtoContext(db))
	router.Initialize(r)
	return r
}
