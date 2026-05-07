package client

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

func (s *Store) GetClientByEmail(email string) (*types.Client, error) {
	rows, err := s.db.Query("SELECT * FROM clients WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	c := new(types.Client)
	for rows.Next() {
		c, err = scanRowsIntoClient(rows)
		if err != nil {
			return nil, err
		}
	}
	if c.ID == 0 {
		return nil, fmt.Errorf("Client not found")
	}

	return c, nil
}

func (s *Store) GetClientByID(id int) (*types.Client, error) {
	rows, err := s.db.Query("SELECT * FROM clients WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	c := new(types.Client)
	for rows.Next() {
		c, err = scanRowsIntoClient(rows)
		if err != nil {
			return nil, err
		}
	}
	if c.ID == 0 {
		return nil, fmt.Errorf("Client not found")
	}

	return c, nil
}

func (s *Store) GetAllClients() ([]types.Client, error) {
	rows, err := s.db.Query("SELECT * FROM clients")
	if err != nil {
		return nil, err
	}

	clients := make([]types.Client, 0)
	for rows.Next() {
		c, err := scanRowsIntoClient(rows)
		if err != nil {
			return nil, err
		}
		clients = append(clients, *c)
	}

	return clients, nil
}

func (s *Store) CreateClient(client types.Client) error {
	_, err := s.db.Exec(
		"INSERT INTO clients (name, email, phoneNumber) VALUES (?,?,?)",
		client.Name, client.Email, client.PhoneNumber,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateClient(client types.Client) error {
	_, err := s.db.Exec(
		"UPDATE clients SET name = ?, email = ?, phoneNumber = ? WHERE id = ?",
		client.Name, client.Email, client.PhoneNumber, client.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteClient(clientID int) error {
	_, err := s.db.Exec(
		"DELETE FROM clients WHERE id = ?",
		clientID,
	)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoClient(rows *sql.Rows) (*types.Client, error) {
	client := new(types.Client)
	err := rows.Scan(
		&client.ID,
		&client.Name,
		&client.Email,
		&client.PhoneNumber,
		&client.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}
