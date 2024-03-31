package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time-todo/pkg/models"
)

func (app *application) createMachine(w http.ResponseWriter, r *http.Request) {
	var newMachine models.Machine

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newMachine)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.machine.Insert(&newMachine)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) getMachine(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	machineData, err := app.machine.GetMachineById(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(machineData)
}

func (app *application) updateMachine(w http.ResponseWriter, r *http.Request) {
	var updatedMachine models.Machine

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&updatedMachine)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.machine.UpdateMachine(&updatedMachine)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
}

func (app *application) deleteMachine(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.machine.DeleteMachineById(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
}
