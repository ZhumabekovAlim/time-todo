package dbs

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"time-todo/pkg/models"
)

type ClientModel struct {
	DB *sql.DB
}

func (m *ClientModel) Insert(clientname, clientmail, clientpass, clientphone, clienttelegram string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(clientpass), 12)
	if err != nil {
		return err
	}

	stmt := `
        INSERT INTO client 
        (clientname, clientmail, clientpass, clientphone, clienttelegram) 
        VALUES (?, ?, ?, ?, ?);`

	_, err = m.DB.Exec(stmt, clientname, clientmail, string(hashedPassword), clientphone, clienttelegram)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return models.ErrDuplicateEmail
		}
		return err
	}

	return nil
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
	stmt := "SELECT idclient, clientpass FROM client WHERE clientmail = ?"
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
