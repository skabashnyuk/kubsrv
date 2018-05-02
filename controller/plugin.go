package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/skabashnyuk/kubsrv/storage"
	"net/http"
	"github.com/skabashnyuk/kubsrv/types"
	"strings"
	"log"
)

type Plugin struct {
	Storage *storage.Storage
}

func (plugin *Plugin) GetPlugin(c *gin.Context) {
	obj, err := plugin.Storage.GetPlugin(&storage.ItemId{
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
	c.JSON(200, obj)
}

func (plugin *Plugin) GetLatestPluginsList(c *gin.Context) {
	ids, exists := c.GetQueryArray("id")
	if exists {
		var cheServices []types.CheService
		for _, k := range ids {
			stringSlice := strings.Split(k, ":")

			obj, err := plugin.Storage.GetCheService(&storage.ItemId{
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
		c.JSON(200, cheServices)

	} else {
		c.String(400, "Invalid request. No id query parameter provided")
	}
}
