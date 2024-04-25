package dbs

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time-todo/pkg/models"
)

type ConvoyModel struct {
	DB *sql.DB
}

func (m *ConvoyModel) Insert(idconvoy_idclient int, convoyname string) error {
	convoystatus := 1
	stmt := `
        INSERT INTO convoy 
        ( idconvoy_idclient, convoyname, convoystatus) 
        VALUES (?, ?, ?);`

	_, err := m.DB.Exec(stmt, idconvoy_idclient, convoyname, convoystatus)
	if err != nil {
		return err
	}

	return nil
}

func (m *ConvoyModel) GetConvoyById(id string) ([]byte, error) {
	stmt := `SELECT * FROM convoy WHERE idconvoy = ?`

	convoyRow := m.DB.QueryRow(stmt, id)

	c := &models.Convoy{}

	err := convoyRow.Scan(&c.IdConvoy, &c.IdConvoy_IdClient, &c.ConvoyName, &c.ConvoyStatus)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	convertedConvoy, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return convertedConvoy, nil
}

func (m *ConvoyModel) UpdateConvoy(convoy *models.Convoy) error {
	stmt := `
    UPDATE convoy
    SET
      idconvoy_idclient = ?,
      convoyname = ?,
      convoystatus = ?
    WHERE
      idconvoy = ?`

	_, err := m.DB.Exec(stmt, convoy.IdConvoy_IdClient, convoy.ConvoyName, 1, convoy.IdConvoy)
	if err != nil {
		return err
	}

	return nil
}

func (m *ConvoyModel) DeleteConvoyById(id string) error {
	stmt := `DELETE FROM convoy WHERE idconvoy = ?`
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
