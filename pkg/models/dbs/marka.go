package dbs

import (
	"database/sql"
	"encoding/json"
	"time-todo/pkg/models"
)

type MarkaModel struct {
	DB *sql.DB
}

func (m *MarkaModel) GetAllMarkas() ([]byte, error) {
	stmt := `SELECT * FROM marka`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var markas []*models.Marka
	for rows.Next() {
		mk := &models.Marka{}
		err := rows.Scan(&mk.IdMarka, &mk.MarkaRus)
		if err != nil {
			return nil, err
		}
		markas = append(markas, mk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	convertedMarkas, err := json.Marshal(markas)
	if err != nil {
		return nil, err
	}
	return convertedMarkas, nil
}
