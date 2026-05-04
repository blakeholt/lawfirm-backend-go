package user

import (
	"database/sql"
	"fmt"
	"lawfirm-go-backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("User not found")
	}

	return u, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("User not found")
	}

	return u, nil
}

func (s *Store) GetAllUsers() ([]types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	users := make([]types.User, 0)
	for rows.Next() {
		u, err := scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, *u)
	}

	return users, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec(
		"INSERT INTO users (firstName, lastName, phoneNumber, email, password, avatar, role) VALUES (?,?,?,?,?,?,?)",
		user.FirstName, user.LastName, user.PhoneNumber, user.Email, user.Password, user.Avatar, user.Role,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateUser(user types.User) error {
	_, err := s.db.Exec(
		"UPDATE users SET firstName = ?, lastName = ?, phoneNumber = ?, email = ?, password = ?, avatar = ?, role = ? WHERE id = ?",
		user.FirstName, user.LastName, user.PhoneNumber, user.Email, user.Password, user.Avatar, user.Role, user.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteUser(userID int) error {
	_, err := s.db.Exec(
		"DELETE FROM users WHERE id = ?", userID,
	)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PhoneNumber,
		&user.Password,
		&user.Avatar,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
