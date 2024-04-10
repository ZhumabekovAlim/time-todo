package main

import (
	"errors"
	"net/http"
	"time-todo/pkg/models"
)

func (app *application) getConvoyInfo(w http.ResponseWriter, r *http.Request) {
	id_client := r.URL.Query().Get("id_client")
	if id_client == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	convoyInfoData, err := app.convoyInfo.GetConvoyInfo(id_client)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(convoyInfoData)
}

func (app *application) getMachineInfo(w http.ResponseWriter, r *http.Request) {
	id_convoy := r.URL.Query().Get("id_convoy")
	if id_convoy == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	machineInfoData, err := app.machineInfo.GetMachineInfo(id_convoy)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(machineInfoData)
}
