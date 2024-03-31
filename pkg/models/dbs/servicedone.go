package dbs

import (
	"database/sql"
	"time-todo/pkg/models"
)

type ServiceDoneModel struct {
	DB *sql.DB
}

// Insert adds a new ServiceDone record to the database.
func (m *ServiceDoneModel) Insert(s *models.ServiceDone) error {
	stmt := `INSERT INTO servicedone (idservicedone_idmachine, idservicedone_idservice, servicedonemotohour, servicedonekilometr, servicedonemiles, servicedonedate, servicedonename)
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := m.DB.Exec(stmt, s.MachineID, s.ServiceID, s.MotoHour, s.Kilometr, s.Miles, s.ServiceDate, s.ServiceName)
	return err
}

// GetByID retrieves a ServiceDone record by its ID.
func (m *ServiceDoneModel) GetByID(id int) (*models.ServiceDone, error) {
	stmt := `SELECT * FROM servicedone WHERE idservicedone = ?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.ServiceDone{}
	err := row.Scan(&s.ID, &s.MachineID, &s.ServiceName, &s.MotoHour, &s.Kilometr, &s.Miles, &s.ServiceDate, &s.ServiceName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	return s, nil
}

// Update modifies an existing ServiceDone record.
func (m *ServiceDoneModel) Update(s *models.ServiceDone) error {
	stmt := `UPDATE servicedone SET idservicedone_idmachine = ?, idservicedone_idservice = ?, servicedonemotohour = ?, servicedonekilometr = ?, servicedonemiles = ?, servicedonedate = ?, servicedonename = ? WHERE idservicedone = ?`
	_, err := m.DB.Exec(stmt, s.MachineID, s.ServiceID, s.MotoHour, s.Kilometr, s.Miles, s.ServiceDate, s.ServiceName, s.ID)
	return err
}

// Delete removes a ServiceDone record from the database.
func (m *ServiceDoneModel) Delete(id int) error {
	stmt := `DELETE FROM servicedone WHERE idservicedone = ?`
	_, err := m.DB.Exec(stmt, id)
	return err
}
