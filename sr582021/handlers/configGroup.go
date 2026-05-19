package handlers

import (
	"encoding/json"
	
	"net/http"
	"sr582021/model"
	"sr582021/services"
	"strconv"
	
	"sr582021/utils"

	"github.com/gorilla/mux"
)

type ConfigGroupHandler struct {
	service services.ConfigGroupService
}

func NewConfigGroupHandler(service services.ConfigGroupService) ConfigGroupHandler {
	return ConfigGroupHandler{
		service: service,
	}
}

func (c ConfigGroupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	versionStr := mux.Vars(r)["version"]
	version, err := strconv.Atoi(versionStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	group, err := c.service.GetGroup(name, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c ConfigGroupHandler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := c.service.GetAllGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(groups)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c ConfigGroupHandler) PostGroup(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	versionStr := mux.Vars(r)["version"]
	version, err := strconv.Atoi(versionStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var configs []model.Config

	err = json.NewDecoder(r.Body).Decode(&configs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.service.PostGroup(name, version, configs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c ConfigGroupHandler) PutGroup(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	versionStr := mux.Vars(r)["version"]
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedGroup model.ConfigGroup

	err = json.NewDecoder(r.Body).Decode(&updatedGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.service.PutGroup(name, version, updatedGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c ConfigGroupHandler) DeleteGroupByVersion(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	versionStr := mux.Vars(r)["version"]
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.service.DeleteGroupByVersion(name, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c ConfigGroupHandler) DeleteConfigByVersion(w http.ResponseWriter, r *http.Request) {
	groupName := mux.Vars(r)["groupName"]
	groupVersionStr := mux.Vars(r)["groupVersion"]
	groupVersion, err := strconv.Atoi(groupVersionStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	configName := mux.Vars(r)["configName"]
	configVersionStr := mux.Vars(r)["configVersion"]
	configVersion, err := strconv.Atoi(configVersionStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.DeleteConfigByVersion(groupName, groupVersion, configName, configVersion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}


func (c ConfigGroupHandler) GetByLabels(w http.ResponseWriter, r *http.Request) {
    name := mux.Vars(r)["name"]
    versionStr := mux.Vars(r)["version"]
    version, err := strconv.Atoi(versionStr)
    if err != nil {
        renderJSON(w, nil, http.StatusBadRequest)
        return
    }

    raw := r.URL.Query().Get("labels")
    query, err := utils.ParseLabels(raw)
    if err != nil {
        renderJSON(w, nil, http.StatusBadRequest)
        return
    }

	configs, err := c.service.GetConfigsByLabels(name, version, query)
    if err != nil {
        renderJSON(w, nil, http.StatusNotFound)
        return
    }
    renderJSON(w, configs, http.StatusOK)
}

func (c ConfigGroupHandler) DeleteByLabels(w http.ResponseWriter, r *http.Request) {
    name := mux.Vars(r)["name"]
    versionStr := mux.Vars(r)["version"]
    version, err := strconv.Atoi(versionStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    raw := r.URL.Query().Get("labels")
    query, err := utils.ParseLabels(raw)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = c.service.DeleteConfigsByLabels(name, version, query)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}