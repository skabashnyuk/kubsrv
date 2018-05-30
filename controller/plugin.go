package controller

import (
	"github.com/skabashnyuk/kubsrv/storage"
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
)

type Plugin struct {
	Storage *storage.Storage
}

func (plugin *Plugin) GetPlugin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	obj, err := plugin.Storage.GetPlugin(&storage.ItemId{
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

func (plugin *Plugin) GetLatestPluginsList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	chePlugins, err := plugin.Storage.GetPlugins(1000, 0)
	if err != nil {
		msg, code := ToHTTPError(err)
		log.Printf("Error in  GetServiceByIdList %s", err.Error())
		http.Error(w, msg, code)

		return
	}
	w.WriteHeader(http.StatusOK)
	WriteJSON(w, chePlugins)
}
