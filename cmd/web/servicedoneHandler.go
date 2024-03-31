package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time-todo/pkg/models"
)

func (app *application) createServiceDone(w http.ResponseWriter, r *http.Request) {
	var newServiceDone models.ServiceDone

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newServiceDone)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.serviceDone.Insert(&newServiceDone)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) getServiceDone(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	serviceDone, err := app.serviceDone.GetByID(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
			return
		} else {
			app.serverError(w, err)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(serviceDone)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Write(jsonData) // Ignoring error for brevity
}

func (app *application) updateServiceDone(w http.ResponseWriter, r *http.Request) {
	var updatedServiceDone models.ServiceDone

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&updatedServiceDone)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.serviceDone.Update(&updatedServiceDone)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteServiceDone(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.serviceDone.Delete(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
