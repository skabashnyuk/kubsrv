package server


import (

	"github.com/skabashnyuk/kubsrv/router"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	router.Initialize(r)
	return r
}
