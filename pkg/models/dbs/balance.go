package dbs

import (
	"database/sql"
	"gorm.io/gorm"
	"time-todo/pkg/models"
)

type BalanceModel struct {
	DB     *sql.DB
	GormDB *gorm.DB
}

func (m *BalanceModel) Insert(balance *models.Balance) error {

	stmt := `
        INSERT INTO balance 
        (idbalance_idclient,balancedatestart,balancedatestop,balance,balancecaption,balancestatus) 
        VALUES (?, ?, ?, ?, ?, ?);`

	_, err := m.DB.Exec(stmt, balance.IdBalanceIdClient, balance.BalanceDateStart, balance.BalanceDateStop, balance.Balance, balance.BalanceCaption, balance.BalanceStatus)
	if err != nil {
		return err
	}

	return nil
}
