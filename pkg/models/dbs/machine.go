package dbs

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time-todo/pkg/models"
)

type MachineModel struct {
	DB *sql.DB
}

func (m *MachineModel) Insert(machine *models.Machine) error {
	stmt := `INSERT INTO machine (idmachine_idconvoy, idmachine_idmodel, idmachine_idtype, machineyear, machinegosnumber, machineoption, machinedatecome, machineseason, machinemotohour, machinekilometr, machinemiles, machinestatus) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	_, err := m.DB.Exec(stmt, machine.IdMachineIdConvoy, machine.IdMachineIdModel, machine.IdMachineIdType, machine.MachineYear, machine.MachineGosNumber, machine.MachineOption, machine.MachineDateCome, machine.MachineSeason, machine.MachineMotoHour, machine.MachineKilometr, machine.MachineMiles, machine.MachineStatus)
	if err != nil {
		return err
	}

	return nil
}

func (m *MachineModel) GetMachineById(id string) ([]byte, error) {
	stmt := `SELECT * FROM machine WHERE machine.idmachine = ?`

	machineRow := m.DB.QueryRow(stmt, id)

	machine := &models.Machine{}

	err := machineRow.Scan(&machine.IdMachine, &machine.IdMachineIdConvoy, &machine.IdMachineIdModel, &machine.IdMachineIdType, &machine.MachineYear, &machine.MachineGosNumber, &machine.MachineOption, &machine.MachineDateCome, &machine.MachineDateOut, &machine.MachineSeason, &machine.MachineMotoHour, &machine.MachineKilometr, &machine.MachineMiles, &machine.MachineStatus)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	convertedMachine, err := json.Marshal(machine)
	if err != nil {
		return nil, err
	}
	return convertedMachine, nil
}

func (m *MachineModel) UpdateMachine(machine *models.Machine) error {
	stmt := `
    UPDATE machine
    SET
      idmachine_idmodel = ?,
      idmachine_idtype = ?,
      machineyear = ?,
      machinegosnumber = ?,
      machineoption = ?,
      machinedatecome = ?,
      machineseason = ?,
      machinemotohour = ?,
      machinekilometr = ?,
      machinemiles = ?,
    WHERE
      idmachine = ?`

	_, err := m.DB.Exec(stmt, machine.IdMachineIdModel, machine.IdMachineIdType, machine.MachineYear, machine.MachineGosNumber, machine.MachineOption, machine.MachineDateCome, machine.MachineDateOut, machine.MachineSeason, machine.MachineMotoHour, machine.MachineKilometr, machine.MachineMiles, machine.MachineStatus, machine.IdMachine)
	if err != nil {
		return err
	}

	return nil
}

func (m *MachineModel) DeleteMachineById(id string) error {
	stmt := `DELETE FROM machine WHERE idmachine = ?`
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
