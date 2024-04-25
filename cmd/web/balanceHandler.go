package main

import (
	"net/http"
	"strconv"
	"time"
	"time-todo/pkg/models"
)

func (app *application) createBalance(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	currentTime := time.Now()
	Balance := &models.Balance{
		IdBalanceIdClient: int64(id),
		BalanceDateStart:  currentTime.Format("2006-01-02"),
		BalanceDateStop:   currentTime.AddDate(1, 0, 0).Format("2006-01-02"),
		Balance:           16000,
		BalanceCaption:    "оплата за услуги сайта time-todo.com на 1 год",
		BalanceStatus:     0,
	}

	result := app.gormDB.Table("balance").Create(&Balance)
	if result.Error != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"balance": Balance}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
