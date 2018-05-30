package controller

import (
	"github.com/skabashnyuk/kubsrv/storage"
	"net/http"
	"github.com/skabashnyuk/kubsrv/types"
	"strings"
	"log"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

type Service struct {
	Storage *storage.Storage
}

func (service *Service) GetService(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	obj, err := service.Storage.GetCheService(&storage.ItemId{
		Name:    params.ByName("name"),
		Version: params.ByName("version")})

	if err != nil {
		msg, code := ToHTTPError(err)

		log.Printf("Error in  GetService %s", err.Error())
		http.Error(w, msg, code)

		return
	}
	w.WriteHeader(http.StatusOK)
	WriteJSON(w, obj)
}

func (service *Service) GetServiceByIdList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ids, exists := GetQueryArray(r, "id")
	if exists {
		var cheServices []types.CheService
		for _, k := range ids {
			stringSlice := strings.Split(k, ":")

			obj, err := service.Storage.GetCheService(&storage.ItemId{
				Name:    stringSlice[0],
				Version: stringSlice[1]})

			if err != nil {
				msg, code := ToHTTPError(err)

				log.Printf("Error in  GetServiceByIdList %s", err.Error())
				http.Error(w, msg, code)
				return
			}
			cheServices = append(cheServices, *obj)
		}
		w.WriteHeader(http.StatusOK)
		WriteJSON(w, cheServices)

	} else {
		w.WriteHeader(400)
		fmt.Fprint(w, "Invalid request. No id query parameter provided")
	}
}
