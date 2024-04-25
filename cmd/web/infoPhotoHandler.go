package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time-todo/pkg/models"
)

func (app *application) createInfoPhoto(w http.ResponseWriter, r *http.Request) {
	var newPhoto models.InfoPhoto

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newPhoto)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.infoPhoto.InsertPhoto(&newPhoto)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) getInfoPhoto(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	infoPhotoData, err := app.infoPhoto.GetPhotoById(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(infoPhotoData)
}

func (app *application) deleteInfoPhoto(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.infoPhoto.DeletePhotoById(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
