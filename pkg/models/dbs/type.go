package dbs

import (
	"database/sql"
	"encoding/json"
	"time-todo/pkg/models"
)

type TypeModel struct {
	DB *sql.DB
}

func (m *TypeModel) GetAllTypes() ([]byte, error) {
	stmt := `SELECT idtype,type_rus FROM type`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []*models.Type
	for rows.Next() {
		t := &models.Type{}
		err := rows.Scan(&t.IdType, &t.TypeRus)
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	convertedTypes, err := json.Marshal(types)
	if err != nil {
		return nil, err
	}
	return convertedTypes, nil
}
