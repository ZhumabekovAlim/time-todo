package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time-todo/pkg/models"
)

func (app *application) createService(w http.ResponseWriter, r *http.Request) {
	var newService models.Service

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newService)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.service.Insert(&newService)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) getService(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	serviceData, err := app.service.GetServiceById(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(serviceData) // Ignoring error for brevity
}

func (app *application) updateService(w http.ResponseWriter, r *http.Request) {
	var updatedService models.Service

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&updatedService)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.service.UpdateService(&updatedService)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteService(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.service.DeleteServiceById(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
