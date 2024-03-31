package dbs

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time-todo/pkg/models"
)

type ServiceModel struct {
	DB *sql.DB
}

func (m *ServiceModel) Insert(service *models.Service) error {
	stmt := `
        INSERT INTO service 
        (idservice_idclient, servicename, motohourstandart, kilometrstandart, milesstandart, servicestatus) 
        VALUES (?, ?, ?, ?, ?, ?);`

	_, err := m.DB.Exec(stmt, service.IdServiceIdClient, service.ServiceName, service.MotoHourStandart, service.KilometrStandart, service.MilesStandart, service.ServiceStatus)
	if err != nil {
		return err
	}

	return nil
}

func (m *ServiceModel) GetServiceById(id string) ([]byte, error) {
	stmt := `SELECT * FROM service WHERE idservice = ?`

	serviceRow := m.DB.QueryRow(stmt, id)

	service := &models.Service{}

	err := serviceRow.Scan(&service.IdService, &service.IdServiceIdClient, &service.ServiceName, &service.MotoHourStandart, &service.KilometrStandart, &service.MilesStandart, &service.ServiceStatus)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	convertedService, err := json.Marshal(service)
	if err != nil {
		return nil, err
	}
	return convertedService, nil
}

func (m *ServiceModel) UpdateService(service *models.Service) error {
	stmt := `
    UPDATE service
    SET
      idservice_idclient = ?,
      servicename = ?,
      motohourstandart = ?,
      kilometrstandart = ?,
      milesstandart = ?,
      servicestatus = ?
    WHERE
      idservice = ?`

	_, err := m.DB.Exec(stmt, service.IdServiceIdClient, service.ServiceName, service.MotoHourStandart, service.KilometrStandart, service.MilesStandart, service.ServiceStatus, service.IdService)
	if err != nil {
		return err
	}

	return nil
}

func (m *ServiceModel) DeleteServiceById(id string) error {
	stmt := `DELETE FROM service WHERE idservice = ?`
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
