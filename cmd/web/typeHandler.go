package main

import "net/http"

func (app *application) getAllTypes(w http.ResponseWriter, r *http.Request) {
	typeData, err := app.types.GetAllTypes()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(typeData)
}
