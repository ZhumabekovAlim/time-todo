package models

import (
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type User struct {
	IdClient          string `json:"idclient"`
	ClientName        string `json:"clientname"`
	ClientMail        string `json:"clientmail"`
	ClientPass        string `json:"clientpass"`
	ClientPhone       string `json:"clientphone"`
	ClientTelegram    string `json:"clienttelegram"`
	ClientDateReg     string `json:"clientdatereg"`
	ClientTimeZone    string `json:"clienttimezone"`
	ClientTimeInfo    string `json:"clienttimeinfo"`
	ClientStatus      string `json:"clientstatus"`
	IdcCient_IdClient string `json:"idclient_idclient"`
}

//
//type Product struct {
//	ID          string `json:"productId"`
//	ProductName string `json:"productName"`
//	CategoryId  string `json:"categoryId"`
//	Price       string `json:"price"`
//	Quantity    string `json:"quantity"`
//	Type        string `json:"type"`
//	PhotoUrl    string `json:"photoUrl"`
//}
//
//type Category struct {
//	ID           string `json:"categoryId"`
//	CategoryName string `json:"categoryName"`
//}
