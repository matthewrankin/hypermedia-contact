package models

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5"
)

// Contact defines the model to hold the data for an individual contact.
type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type ContactModel struct {
	DB *sql.DB
}

// Insert inserts a new contact into the database.
func (m *ContactModel) Insert(contact Contact) (int, error) {
	query := `
    INSERT INTO contacts (first_name, last_name, email, phone)
    VALUES($1, $2, $3, $4)
    RETURNING id;`
	id := 0

	err := m.DB.QueryRow(query, contact.FirstName, contact.LastName, contact.Email, contact.Phone).
		Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get returns a specific contact based on the given ID.
func (m *ContactModel) Get(id int) (Contact, error) {
	return Contact{}, nil
}
