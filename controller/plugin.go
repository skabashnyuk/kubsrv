package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/skabashnyuk/kubsrv/storage"
	"net/http"
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

	chePlugins, err := plugin.Storage.GetPlugins(1000, 0)
	if err != nil {
		msg, code := ToHTTPError(err)
		if gin.IsDebugging() {
			log.Printf("Error in  GetServiceByIdList %s", err.Error())
		}
		http.Error(c.Writer, msg, code)
		c.Abort()
		return
	}
	c.JSON(200, chePlugins)
}
