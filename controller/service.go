package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/skabashnyuk/kubsrv/storage"
	"github.com/skabashnyuk/kubsrv/render"
	"net/http"
	"github.com/skabashnyuk/kubsrv/types"
	"strings"
)

func GetService(c *gin.Context) {
	obj, err := storage.GetCheService(&storage.ItemId{
		Name:    c.Param("name"),
		Version: c.Param("version")})

	if err != nil {
		msg, code := ToHTTPError(err)
		//Error(w, msg, code)
		http.Error(c.Writer, msg, code)
		c.Abort()
		return
	}
	c.Render(200, render.GYAML{Data: obj})
}

func GetServiceByIdList(c *gin.Context) {
	ids, exists := c.GetQueryArray("id")
	if exists {
		var cheServices []types.CheService
		for _, k := range ids {
			stringSlice := strings.Split(k, ":")

			obj, err := storage.GetCheService(&storage.ItemId{
				Name:    stringSlice[0],
				Version: stringSlice[1]})

			if err != nil {
				msg, code := ToHTTPError(err)
				//Error(w, msg, code)
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
