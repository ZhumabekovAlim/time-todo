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

func (app *application) deleteBalance(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	result := app.gormDB.Table("balance").Delete(&models.Balance{}, id)
	if result.RowsAffected == 0 {
		err := app.writeJSON(w, http.StatusNotFound, envelope{"message": "no balance deleted"}, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "balance deleted"}, nil)
	if err != nil {
		app.clientError(w, http.StatusInternalServerError)
	}
}

func (app *application) updateBalance(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var input struct {
		IdBalance         int64     `json:"idbalance"`
		IdBalanceIdClient int64     `json:"idbalance_idclient"`
		BalanceDateStart  time.Time `json:"balancedatestart"`
		BalanceDateStop   time.Time `json:"balancedatestop"`
		Balance           int64     `json:"balance"`
		BalanceCaption    string    `json:"balancecaption"`
		BalanceStatus     int64     `json:"balancestatus"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.gormDB.
		Table("balance").
		Model(&models.Balance{}).
		Select("*").
		Where("idbalance = ?", id).
		Updates(&models.Balance{
			IdBalance:         int64(id),
			IdBalanceIdClient: input.IdBalanceIdClient,
			BalanceDateStart:  time.Now().Format("2006-01-02"),
			BalanceDateStop:   time.Now().AddDate(1, 0, 0).Format("2006-01-02"),
			Balance:           input.Balance,
			BalanceCaption:    input.BalanceCaption,
			BalanceStatus:     input.BalanceStatus,
		})

	err = app.writeJSON(w, http.StatusOK, envelope{"balance": input}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
