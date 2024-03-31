// dbs/repairdonemodel.go

package dbs

import (
	"database/sql"
	"time-todo/pkg/models"
)

type RepairDoneModel struct {
	DB *sql.DB
}

func (m *RepairDoneModel) Insert(rd *models.RepairDone) error {
	stmt := `INSERT INTO repairdone (idrepairdone_idmachine, idrepairdone_idrepair, repairdonemotohour, repairdonekilometr, repairdonemiles, repairdonecaption, repairdonedate, repairdonename)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := m.DB.Exec(stmt, rd.IdRepairDoneIdMachine, rd.IdRepairDoneIdRepair, rd.RepairDoneMotoHour, rd.RepairDoneKilometr, rd.RepairDoneMiles, rd.RepairDoneCaption, rd.RepairDoneDate, rd.RepairDoneName)
	return err
}

func (m *RepairDoneModel) Get(id int) (*models.RepairDone, error) {
	stmt := `SELECT * FROM repairdone WHERE idrepairdone = ?`
	row := m.DB.QueryRow(stmt, id)
	rd := &models.RepairDone{}
	err := row.Scan(&rd.IdRepairDone, &rd.IdRepairDoneIdMachine, &rd.IdRepairDoneIdRepair, &rd.RepairDoneMotoHour, &rd.RepairDoneKilometr, &rd.RepairDoneMiles, &rd.RepairDoneCaption, &rd.RepairDoneDate, &rd.RepairDoneName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	return rd, nil
}

func (m *RepairDoneModel) Update(rd *models.RepairDone) error {
	stmt := `UPDATE repairdone SET idrepairdone_idmachine = ?, idrepairdone_idrepair = ?, repairdonemotohour = ?, repairdonekilometr = ?, repairdonemiles = ?, repairdonecaption = ?, repairdonedate = ?, repairdonename = ? WHERE idrepairdone = ?`
	_, err := m.DB.Exec(stmt, rd.IdRepairDoneIdMachine, rd.IdRepairDoneIdRepair, rd.RepairDoneMotoHour, rd.RepairDoneKilometr, rd.RepairDoneMiles, rd.RepairDoneCaption, rd.RepairDoneDate, rd.RepairDoneName, rd.IdRepairDone)
	return err
}

func (m *RepairDoneModel) Delete(id int) error {
	stmt := `DELETE FROM repairdone WHERE idrepairdone = ?`
	_, err := m.DB.Exec(stmt, id)
	return err
}
