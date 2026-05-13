package handlers

import (
	"net/http"
	"sr582021/services"
	"strconv"

	"github.com/gorilla/mux"
)

type ConfigHandler struct {
	service services.ConfigService
}

func NewConfigHandler(service services.ConfigService) ConfigHandler {
	return ConfigHandler{
		service: service,
	}
}

func (c ConfigHandler) Get(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	versionStr := mux.Vars(r)["version"]
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		renderJSON(w, nil, http.StatusBadRequest)
		return
	}
	config, err := c.service.Get(name, version)
	if err != nil {
		renderJSON(w, nil, http.StatusNotFound)
		return

	}
	renderJSON(w, config, http.StatusOK)
}
