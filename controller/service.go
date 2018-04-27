package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/skabashnyuk/kubsrv/storage"
	"github.com/skabashnyuk/kubsrv/render"
	"net/http"
	"github.com/skabashnyuk/kubsrv/types"
	"strings"
	"log"
)

type Service struct {
	Storage *storage.Storage
}

func (service *Service) GetService(c *gin.Context) {
	obj, err := service.Storage.GetCheService(&storage.ItemId{
		Name:    c.Param("name"),
		Version: c.Param("version")})

	if err != nil {
		msg, code := ToHTTPError(err)
		if gin.IsDebugging() {
			log.Printf("Error in  GetService %s", err.Error())
		}
		http.Error(c.Writer, msg, code)
		c.Abort()
		return
	}
	c.Render(200, render.GYAML{Data: obj})
}

func (service *Service) GetServiceByIdList(c *gin.Context) {
	ids, exists := c.GetQueryArray("id")
	if exists {
		var cheServices []types.CheService
		for _, k := range ids {
			stringSlice := strings.Split(k, ":")

			obj, err := service.Storage.GetCheService(&storage.ItemId{
				Name:    stringSlice[0],
				Version: stringSlice[1]})

			if err != nil {
				msg, code := ToHTTPError(err)
				if gin.IsDebugging() {
					log.Printf("Error in  GetServiceByIdList %s", err.Error())
				}
				http.Error(c.Writer, msg, code)
				c.Abort()
				return
			}
			cheServices = append(cheServices, *obj)
		}
		c.Render(200, render.GYAML{Data: cheServices})

	} else {
		c.String(400, "Invalid request. No id query parameter provided")
	}
}
