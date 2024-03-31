package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time-todo/pkg/models"
)

func (app *application) createConvoy(w http.ResponseWriter, r *http.Request) {
	var newConvoy models.Convoy

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newConvoy)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.convoy.Insert(newConvoy.IdConvoy, newConvoy.IdConvoy_IdClient, newConvoy.ConvoyName, newConvoy.ConvoyStatus)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) getConvoy(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	convoyData, err := app.convoy.GetConvoyById(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(convoyData) // Ignoring error for brevity
}

func (app *application) updateConvoy(w http.ResponseWriter, r *http.Request) {
	var updatedConvoy models.Convoy

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&updatedConvoy)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.convoy.UpdateConvoy(&updatedConvoy)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
}

func (app *application) deleteConvoy(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.convoy.DeleteConvoyById(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
}
