package main

import (
	"net/http"
	"strconv"
	"time"
	"time-todo/pkg/models"
)

func (app *application) createBalance(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	currentTime := time.Now()
	Balance := &models.Balance{
		IdBalanceIdClient: idInt,
		BalanceDateStart:  currentTime.Format("2006-01-02"),
		BalanceDateStop:   currentTime.AddDate(1, 0, 0).Format("2006-01-02"),
		Balance:           16000,
		BalanceCaption:    "оплата за услуги сайта time-todo.com на 1 год",
		BalanceStatus:     0,
	}
	err = app.balance.Insert(Balance)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
