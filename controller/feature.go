package controller

import (
	"net/http"
	"github.com/skabashnyuk/kubsrv/storage"
	"strings"
	"github.com/skabashnyuk/kubsrv/types"
	"log"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

type Feature struct {
	Storage *storage.Storage
}

func (feature *Feature) GetFeature(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	obj, err := feature.Storage.GetCheFeature(&storage.ItemId{
		Name:    params.ByName("name"),
		Version: params.ByName("version")})

	if err != nil {
		msg, code := ToHTTPError(err)

		log.Printf("Error in  GetFeature %s", err.Error())
		http.Error(w, msg, code)
		return
	}
	w.WriteHeader(http.StatusOK)
	WriteJSON(w, obj)
}

func (feature *Feature) GetFeatureByIdList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ids, exists := GetQueryArray(r, "id")
	if exists {
		var cheFeatures []types.CheFeature
		for _, k := range ids {
			stringSlice := strings.Split(k, ":")

			obj, err := feature.Storage.GetCheFeature(&storage.ItemId{
				Name:    stringSlice[0],
				Version: stringSlice[1]})

			if err != nil {
				msg, code := ToHTTPError(err)
				log.Printf("Error in  GetFeatureByIdList %s", err.Error())
				http.Error(w, msg, code)

				return
			}
			cheFeatures = append(cheFeatures, *obj)
		}

		w.WriteHeader(http.StatusOK)
		WriteJSON(w, cheFeatures)

	} else {
		w.WriteHeader(400)
		fmt.Fprint(w, "Invalid request. No id query parameter provided")
	}
}
