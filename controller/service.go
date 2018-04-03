package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/skabashnyuk/kubsrv/storage"
	"github.com/skabashnyuk/kubsrv/model/v1"
)

func CreateService(c *gin.Context) {

	db := storage.DBInstance(c)
	service := v1.CheService{}

	if err := c.Bind(&service); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Write("service",service.ObjectMeta.Name ,&service); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, service)
}