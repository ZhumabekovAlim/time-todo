package dbs

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time-todo/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, phone, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO customer (name, email, phone, password, role) VALUES ($1, $2, $3, $4, $5);`

	_, err = m.DB.Exec(stmt, name, email, phone, string(hashedPassword), "USER")
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && strings.Contains(pgErr.Message, "users_uc_email") {
			return models.ErrDuplicateEmail
		}
		return err
	}

	return nil
}

func (m *UserModel) GetAllUsers() ([]byte, error) {
	stmt := `SELECT * FROM customer`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var users []*models.User

	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Role)
		if err != nil {
			return nil, err
		}
		// Append it to the slice of snippets.
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	convertedAllUser, err := json.Marshal(users)
	return convertedAllUser, nil
}

func (m *UserModel) GetUserRoleById(id int) string {
	stmt := `SELECT role FROM customer WHERE id = $1`
	var role string
	err := m.DB.QueryRow(stmt, id).Scan(&role)
	if err != nil {
		return ""
	}

	return role
}

func (m *UserModel) GetUserById(id string) ([]byte, error) {
	stmt := `SELECT * FROM customer WHERE id = $1`

	userRow := m.DB.QueryRow(stmt, id)

	u := &models.User{}

	err := userRow.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.Password, &u.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	convertedUser, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return convertedUser, nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte
	stmt := "SELECT id, password FROM customer WHERE email = $1"
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

func (m *UserModel) UpdateUser(user *models.User, id int) error {
	stmt := `
    UPDATE customer
    SET
      name = $2,
      email = $3,
      phone = $4,
      role = $5
    WHERE
      id = $1`

	_, err := m.DB.Exec(stmt, id, user.Name, user.Email, user.Phone, user.Role)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) DeleteUserById(id int) {
	stmt := `DELETE FROM customer WHERE id = $1`
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return
	}
}
