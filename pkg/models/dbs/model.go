package dbs

import (
	"database/sql"
	"encoding/json"
	"time-todo/pkg/models"
)

type ModelModel struct {
	DB *sql.DB
}

func (m *ModelModel) GetAllModels(idmarka string) ([]byte, error) {
	stmt := `SELECT * FROM model WHERE 	idmodel_idmarka = ?`

	rows, err := m.DB.Query(stmt, idmarka)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modelmachine []*models.Model
	for rows.Next() {
		md := &models.Model{}
		err := rows.Scan(&md.IdModel, &md.IdModelIdMarka, &md.Model, &md.ModelStatus)
		if err != nil {
			return nil, err
		}
		modelmachine = append(modelmachine, md)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	convertedModels, err := json.Marshal(modelmachine)
	if err != nil {
		return nil, err
	}
	return convertedModels, nil
}
