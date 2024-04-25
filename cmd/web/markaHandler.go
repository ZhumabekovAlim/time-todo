package main

import "net/http"

func (app *application) getAllMarkas(w http.ResponseWriter, r *http.Request) {
	markaData, err := app.marka.GetAllMarkas()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(markaData)
}
