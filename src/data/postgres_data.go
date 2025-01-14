package data

import (
	"database/sql"
	"errors"
	"tech-challenge/src/models"
)

type PostgresEmailData struct {
	db *sql.DB
}

func NewPosrgresEmailData(db *sql.DB) *PostgresEmailData {
	return &PostgresEmailData{db: db}
}

func (ped *PostgresEmailData) getEmails() ([]models.Email, error) {
	query := `SELECT message_id, from, to FROM emails`
	rows, err := ped.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emails []models.Email
	for rows.Next() {
		var e models.Email
		if err := rows.Scan(&e.MessageID, &e.From, &e.To); err != nil {
			return nil, err
		}
		emails = append(emails, e)
	}
	return emails, nil
}

func (ped *PostgresEmailData) getEmailById(id int) (*models.Email, error) {
	query := `SELECT message_id, from, to FROM emails WHERE id = $1`
	var email models.Email
	err := ped.db.QueryRow(query, id).Scan(&email.MessageID, &email.From, &email.To)
	if err == sql.ErrNoRows {
		return nil, errors.New("email not found")
	}
	if err != nil {
		return nil, err
	}
	return &email, nil
}

func (ped *PostgresEmailData) createEmail(email *models.Email) error {
	query := `INSERT INTO emails (message_id, from, to) VALUES ($1, $2, $3) RETURNING id`
	return ped.db.QueryRow(query, email.MessageID, email.From, email.To).Scan(&u.ID)
}

func (ped *PostgresEmailData) updateEmail(email *models.Email) error {
	query := `UPDATE emails SET message_id = $1, from = $2, to = $3 WHERE id = $4`
	result, err := ped.db.Exec(query, email.MessageID, email.From, email.To)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("email not found")
	}
	return nil
}

func (ped *PostgresEmailData) deleteEmail(id int) error {
	query := `DELETE FROM emails WHERE id = $1`
	result, err := ped.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("email not found")
	}
	return nil
}
