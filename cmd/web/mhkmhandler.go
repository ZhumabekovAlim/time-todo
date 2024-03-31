package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time-todo/pkg/models"
)

func (app *application) createMHKM(w http.ResponseWriter, r *http.Request) {
	var newMHKM models.MhKm

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newMHKM)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.mhkm.Insert(&newMHKM)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) getMHKM(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	mhkmData, err := app.mhkm.GetMhKmById(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(mhkmData)
}

func (app *application) updateMHKM(w http.ResponseWriter, r *http.Request) {
	var updatedMHKM models.MhKm

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&updatedMHKM)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.mhkm.UpdateMhKm(&updatedMHKM)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteMHKM(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.mhkm.DeleteMhKmById(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
