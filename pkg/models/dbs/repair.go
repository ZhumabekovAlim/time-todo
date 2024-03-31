package dbs

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time-todo/pkg/models"
)

type RepairModel struct {
	DB *sql.DB
}

func (m *RepairModel) Insert(repair *models.Repair) error {
	stmt := `
        INSERT INTO repair 
        (idrepair_idclient, repairname, repaircaption, repairstatus) 
        VALUES (?, ?, ?, ?);`

	_, err := m.DB.Exec(stmt, repair.IdRepairIdClient, repair.RepairName, repair.RepairCaption, repair.RepairStatus)
	if err != nil {
		return err
	}

	return nil
}

func (m *RepairModel) GetRepairById(id string) ([]byte, error) {
	stmt := `SELECT * FROM repair WHERE idrepair = ?`

	repairRow := m.DB.QueryRow(stmt, id)

	repair := &models.Repair{}

	err := repairRow.Scan(&repair.IdRepair, &repair.IdRepairIdClient, &repair.RepairName, &repair.RepairCaption, &repair.RepairStatus)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	convertedRepair, err := json.Marshal(repair)
	if err != nil {
		return nil, err
	}
	return convertedRepair, nil
}

func (m *RepairModel) UpdateRepair(repair *models.Repair) error {
	stmt := `
    UPDATE repair
    SET
      idrepair_idclient = ?,
      repairname = ?,
      repaircaption = ?,
      repairstatus = ?
    WHERE
      idrepair = ?`

	_, err := m.DB.Exec(stmt, repair.IdRepairIdClient, repair.RepairName, repair.RepairCaption, repair.RepairStatus, repair.IdRepair)
	if err != nil {
		return err
	}

	return nil
}

func (m *RepairModel) DeleteRepairById(id string) error {
	stmt := `DELETE FROM repair WHERE idrepair = ?`
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
