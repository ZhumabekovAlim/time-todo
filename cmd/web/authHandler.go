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

func (app *application) signupClient(w http.ResponseWriter, r *http.Request) {
	var newClient models.Client

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newClient)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.client.Insert(newClient.ClientName, newClient.ClientMail, newClient.ClientPass, newClient.ClientPhone, newClient.ClientTelegram, newClient.ClientDateReg, newClient.ClientTimeZone, newClient.ClientTimeInfo, newClient.ClientStatus, newClient.IdcCient_IdClient)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated) // 201
}

func (app *application) loginClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	clientId, err := app.client.Authenticate(client.ClientMail, client.ClientPass)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			app.clientError(w, http.StatusBadRequest)
			return
		} else {
			app.serverError(w, err)

			return
		}
	}

	responseUser, err := app.client.GetUserById(strconv.Itoa(clientId))
	//if app.session.Exists(r, "authenticatedUserID") {
	//	w.WriteHeader(http.StatusOK)
	//	_, err := w.Write(responseUser)
	//	if err != nil {
	//		return
	//	}
	//	return
	//}
	//app.session.Put(r, "authenticatedUserID", client)
	//if err != nil {
	//	return
	//}
	_, err = w.Write(responseUser)
	if err != nil {
		return
	}
}

//
//func (app *application) logOut(w http.ResponseWriter, r *http.Request) {
//	if !app.session.Exists(r, "authenticatedUserID") {
//		w.WriteHeader(http.StatusUnauthorized)
//		return
//	}
//	app.session.Pop(r, "authenticatedUserID")
//	w.WriteHeader(http.StatusOK)
//}
//
//func (app *application) testGin(c *gin.Context) {
//	user, err := app.user.GetUserById("1")
//	if err != nil {
//		app.clientError(c.Writer, http.StatusOK)
//	}
//	c.JSON(http.StatusOK, user)
//}
