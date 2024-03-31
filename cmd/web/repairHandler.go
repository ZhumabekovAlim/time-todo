package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time-todo/pkg/models"
)

func (app *application) createRepair(w http.ResponseWriter, r *http.Request) {
	var newRepair models.Repair

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newRepair)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.repair.Insert(&newRepair)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) getRepair(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	repairData, err := app.repair.GetRepairById(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(repairData)
}

func (app *application) updateRepair(w http.ResponseWriter, r *http.Request) {
	var updatedRepair models.Repair

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&updatedRepair)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.repair.UpdateRepair(&updatedRepair)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteRepair(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.repair.DeleteRepairById(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
}
