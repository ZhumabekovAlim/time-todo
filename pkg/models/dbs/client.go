package dbs

import (
	"database/sql"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time-todo/pkg/models"
)

type ClientModel struct {
	DB *sql.DB
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrDuplicatePhone = errors.New("duplicate phone")
)

func (m *ClientModel) Insert(clientname, clientmail, clientpass, clientphone, clienttelegram string) error {

	var exists int
	emailCheckQuery := "SELECT COUNT(*) FROM client WHERE clientmail = ?"
	phoneCheckQuery := "SELECT COUNT(*) FROM client WHERE clientphone = ?"

	err := m.DB.QueryRow(emailCheckQuery, clientmail).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return ErrDuplicateEmail
	}

	err = m.DB.QueryRow(phoneCheckQuery, clientphone).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return ErrDuplicatePhone
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(clientpass), 12)
	if err != nil {
		return err
	}

	stmt := `
        INSERT INTO client 
        (clientname, clientmail, clientpass, clientphone, clienttelegram, clientstatus) 
        VALUES (?, ?, ?, ?, ?, ?);`

	_, err = m.DB.Exec(stmt, clientname, clientmail, string(hashedPassword), clientphone, clienttelegram, 0)
	if err != nil {
		return err
	}

	return nil
}

func (m *ClientModel) VerifyClient(email string) (int, error) {
	var clientID int

	query := `SELECT idclient FROM client WHERE clientmail = ? AND clientstatus = 0;`
	err := m.DB.QueryRow(query, email).Scan(&clientID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("no matching unverified client found")
		}
		return 0, err
	}

	stmt := `
        UPDATE client
        SET clientstatus = 1
        WHERE idclient = ?;`
	_, err = m.DB.Exec(stmt, clientID)
	if err != nil {
		return 0, err
	}

	return clientID, nil
}

//	func (m *ClientModel) GetAllUsers() ([]byte, error) {
//		stmt := `SELECT * FROM client`
//
//		rows, err := m.DB.Query(stmt)
//		if err != nil {
//			return nil, err
//		}
//
//		var users []*models.Client
//
//		for rows.Next() {
//			user := &models.Client{}
//			err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Role)
//			if err != nil {
//				return nil, err
//			}
//			// Append it to the slice of snippets.
//			users = append(users, user)
//		}
//
//		if err = rows.Err(); err != nil {
//			return nil, err
//		}
//
//		convertedAllUser, err := json.Marshal(users)
//		return convertedAllUser, nil
//	}
func (m *ClientModel) GetUserRoleById(id int) string {
	stmt := `SELECT clientstatus FROM client WHERE idclient = $1`
	var role string
	err := m.DB.QueryRow(stmt, id).Scan(&role)
	if err != nil {
		return ""
	}

	return role
}

func (m *ClientModel) GetUserById(id string) ([]byte, error) {
	stmt := `SELECT * FROM client WHERE idclient = ?`

	userRow := m.DB.QueryRow(stmt, id)

	c := &models.Client{}

	err := userRow.Scan(&c.IdClient, &c.ClientName, &c.ClientMail, &c.ClientPass, &c.ClientPhone, &c.ClientTelegram, &c.ClientDateReg, &c.ClientTimeZone, &c.ClientTimeInfo, &c.ClientStatus, &c.IdcCient_IdClient)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	convertedUser, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return convertedUser, nil
}

func (m *ClientModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte
	stmt := "SELECT idclient, clientpass FROM client WHERE clientmail = ? AND clientstatus = 1"
	row := m.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	return id, nil
}

//
//func (m *ClientModel) UpdateUser(user *models.User, id int) error {
//	stmt := `
//    UPDATE customer
//    SET
//      name = $2,
//      email = $3,
//      phone = $4,
//      role = $5
//    WHERE
//      id = $1`
//
//	_, err := m.DB.Exec(stmt, id, user.Name, user.Email, user.Phone, user.Role)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (m *ClientModel) DeleteUserById(id int) {
//	stmt := `DELETE FROM customer WHERE id = $1`
//	_, err := m.DB.Exec(stmt, id)
//	if err != nil {
//		return
//	}
//}
