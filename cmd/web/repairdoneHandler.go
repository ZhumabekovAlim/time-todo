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

func (app *application) createRepairDone(w http.ResponseWriter, r *http.Request) {
	var newRepairDone models.RepairDone

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newRepairDone)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.repairDone.Insert(&newRepairDone)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) getRepairDone(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	repairDone, err := app.repairDone.Get(id)
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
	jsonData, err := json.Marshal(repairDone)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Write(jsonData)
}

func (app *application) updateRepairDone(w http.ResponseWriter, r *http.Request) {
	var updatedRepairDone models.RepairDone

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&updatedRepairDone)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.repairDone.Update(&updatedRepairDone)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteRepairDone(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.repairDone.Delete(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
