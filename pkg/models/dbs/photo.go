package dbs

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time-todo/pkg/models"
)

type InfoPhotoModel struct {
	DB *sql.DB
}

func (m *InfoPhotoModel) InsertPhoto(photo *models.InfoPhoto) error {
	stmt := `INSERT INTO infophoto ( idinfophoto_idinfo, infophoto) 
VALUES (?, ?, ?);`

	_, err := m.DB.Exec(stmt, photo.IdInfoPhotoIdInfo, photo.InfoPhoto)
	if err != nil {
		return err
	}

	return nil
}

func (m *InfoPhotoModel) GetPhotoById(id string) ([]byte, error) {
	stmt := `SELECT * FROM infophoto WHERE idinfophoto  = ?`

	photoRow := m.DB.QueryRow(stmt, id)

	p := &models.InfoPhoto{}

	err := photoRow.Scan(&p.IdInfoPhoto, &p.IdInfoPhotoIdInfo, &p.InfoPhoto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	convertedPhoto, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return convertedPhoto, nil
}

func (m *InfoPhotoModel) DeletePhotoById(id string) error {
	stmt := `DELETE FROM infophoto WHERE idinfophoto = ?`
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}
