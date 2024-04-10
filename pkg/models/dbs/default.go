package dbs

import (
	"database/sql"
	"encoding/json"
	"time-todo/pkg/models"
)

type ConvoyInfoModel struct {
	DB *sql.DB
}

type MachineInfoModel struct {
	DB *sql.DB
}

func (m *ConvoyInfoModel) GetConvoyInfo(idconvoy_idclient string) ([]byte, error) {
	stmt := `select idconvoy from convoy where idconvoy_idclient=? AND convoystatus>0 AND convoystatus<9 order by convoyname;`

	rows, err := m.DB.Query(stmt, idconvoy_idclient)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var convoys []models.ConvoyInfo

	for rows.Next() {
		var info models.ConvoyInfo
		err := rows.Scan(&info.IdConvoy)
		if err != nil {
			return nil, err
		}

		convoys = append(convoys, info)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	convertedInfo, err := json.Marshal(convoys)
	if err != nil {
		return nil, err
	}

	return convertedInfo, nil
}

func (m *MachineInfoModel) GetMachineInfo(idmachine_idconvoy string) ([]byte, error) {
	stmt := `SELECT type_rus, marka, model, machineyear, machinegosnumber 
             FROM machine 
             INNER JOIN type ON idmachine_idtype=idtype 
             INNER JOIN model ON idmachine_idmodel=idmodel 
             INNER JOIN marka ON idmodel_idmarka=idmarka 
             WHERE idmachine_idconvoy=? AND machinestatus>0 AND machinestatus<9
             ORDER BY type_rus, marka, model, machineyear`

	rows, err := m.DB.Query(stmt, idmachine_idconvoy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var machines []models.MachineInfo

	for rows.Next() {
		var info models.MachineInfo
		err := rows.Scan(&info.TypeRus, &info.Marka, &info.Model, &info.MachineYear, &info.MachineGosnumber)
		if err != nil {
			return nil, err
		}

		machines = append(machines, info)
	}

	// Check for any error that occurred during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	convertedInfo, err := json.Marshal(machines)
	if err != nil {
		return nil, err
	}

	return convertedInfo, nil
}
