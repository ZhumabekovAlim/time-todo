package dbs

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time-todo/pkg/models"
)

type MhKmModel struct {
	DB *sql.DB
}

func (m *MhKmModel) Insert(mhkm *models.MhKm, flag string) error {

	if flag == "moto" {
		var maxMotoHour int
		err := m.DB.QueryRow("SELECT MAX(motohour) FROM mhkm WHERE idmhkm_idmachine = ?", mhkm.IdMHKMIdMachine).Scan(&maxMotoHour)
		if err != nil {
			return err
		}

		if mhkm.MotoHour > maxMotoHour {
			stmt := `
        INSERT INTO mhkm
        (idmhkm_idmachine, motohour, kilometr, miles, mhkmdate, mhkmname) 
        VALUES (?, ?, ?, ?, ?, ?);`

			_, err := m.DB.Exec(stmt, mhkm.IdMHKMIdMachine, mhkm.MotoHour, 0, 0, mhkm.MHKMDate, mhkm.MHKMName)
			if err != nil {
				return err
			}
		}
	} else if flag == "kilo" {
		var maxKilometr int
		err := m.DB.QueryRow("SELECT MAX(kilometr) FROM mhkm WHERE idmhkm_idmachine = ?", mhkm.IdMHKMIdMachine).Scan(&maxKilometr)
		if err != nil {
			return err
		}

		if mhkm.Kilometr > maxKilometr {
			stmt := `
        INSERT INTO mhkm
        (idmhkm_idmachine, motohour, kilometr, miles, mhkmdate, mhkmname) 
        VALUES (?, ?, ?, ?, ?, ?);`

			_, err := m.DB.Exec(stmt, mhkm.IdMHKMIdMachine, 0, mhkm.Kilometr, 0, mhkm.MHKMDate, mhkm.MHKMName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *MhKmModel) GetMhKmById(id string) ([]byte, error) {
	stmt := `SELECT * FROM mhkm WHERE idmhkm = ?`

	mhkmRow := m.DB.QueryRow(stmt, id)

	mhkm := &models.MhKm{}

	err := mhkmRow.Scan(&mhkm.IdMHKM, &mhkm.IdMHKMIdMachine, &mhkm.MotoHour, &mhkm.Kilometr, &mhkm.Miles, &mhkm.MHKMDate, &mhkm.MHKMName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	convertedMhKm, err := json.Marshal(mhkm)
	if err != nil {
		return nil, err
	}
	return convertedMhKm, nil
}

func (m *MhKmModel) UpdateMhKm(mhkm *models.MhKm) error {
	stmt := `
    UPDATE mhkm
    SET
      idmhkm_idmachine = ?,
      motohour = ?,
      kilometr = ?,
      miles = ?,
      mhkmdate = ?,
      mhkmname = ?
    WHERE
      idmhkm = ?`

	_, err := m.DB.Exec(stmt, mhkm.IdMHKMIdMachine, mhkm.MotoHour, mhkm.Kilometr, mhkm.Miles, mhkm.MHKMDate, mhkm.MHKMName, mhkm.IdMHKM)
	if err != nil {
		return err
	}

	return nil
}

// DeleteMhKmById deletes a MhKm record by its ID.
func (m *MhKmModel) DeleteMhKmById(id string) error {
	stmt := `DELETE FROM mhkm WHERE idmhkm = ?`
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
