package main

import (
	"bytes"
	"database/sql"
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

	err = app.client.Insert(newClient.ClientName, newClient.ClientMail, newClient.ClientPass, newClient.ClientPhone, newClient.ClientTelegram)
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

	_, err = w.Write(responseUser)
	if err != nil {
		return
	}
}

func (app *application) verifyClient(w http.ResponseWriter, r *http.Request) {
	type verificationRequest struct {
		Email string `json:"clientmail"`
		Code  string `json:"code"`
	}

	var req verificationRequest

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	if req.Code != "12345" {
		app.serverError(w, err)
		return
	}
	clientID, err := app.client.VerifyClient(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.clientError(w, http.StatusNotFound)
			return
		} else {
			app.serverError(w, err)
			return
		}
	}

	response := struct {
		ClientID int `json:"idclient"`
	}{
		ClientID: clientID,
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
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
