package main

import "net/http"

func (app *application) getAllModels(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	modelData, err := app.models.GetAllModels(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(modelData) // Ignoring error for brevity
}
