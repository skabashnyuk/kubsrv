package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/skabashnyuk/kubsrv/render"
	"github.com/skabashnyuk/kubsrv/storage"
	"strings"
	"github.com/skabashnyuk/kubsrv/types"
)

type Feature struct {
	Storage *storage.Storage
}

func (feature *Feature) GetFeature(c *gin.Context) {
	obj, err := feature.Storage.GetCheFeature(&storage.ItemId{
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

func (feature *Feature) GetFeatureByIdList(c *gin.Context) {
	ids, exists := c.GetQueryArray("id")
	if exists {
		var cheFeatures []types.CheFeature
		for _, k := range ids {
			stringSlice := strings.Split(k, ":")

			obj, err := feature.Storage.GetCheFeature(&storage.ItemId{
				Name:    stringSlice[0],
				Version: stringSlice[1]})

			if err != nil {
				msg, code := ToHTTPError(err)
				//Error(w, msg, code)
				http.Error(c.Writer, msg, code)
				c.Abort()
				return
			}
			cheFeatures = append(cheFeatures, *obj)
		}
		c.Render(200, render.GYAML{Data: cheFeatures})

	} else {
		c.String(400, "Invalid request. No id query parameter provided")
	}
}
