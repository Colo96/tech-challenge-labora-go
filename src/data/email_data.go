package data

import (
	"errors"
	"sync"
	"tech-challenge/src/models"
)

type EmailData struct {
	emails []models.Email
	nextID int
	mutex  sync.RWMutex
}

func NewUserRepository() *EmailData {
	return &EmailData{
		emails: make([]models.Email, 0),
		nextID: 1,
	}
}

func (ed *EmailData) getEmails() ([]models.Email, error) {
	ed.mutex.RLock()
	defer ed.mutex.RUnlock()
	return ed.emails, nil
}

func (ed *EmailData) getEmailById(id int) (*models.Email, error) {
	ed.mutex.RLock()
	defer ed.mutex.RUnlock()

	for _, e := range ed.emails {
		if e.ID == id {
			return &e, nil
		}
	}
	return nil, errors.New("user not found")
}

func (ed *EmailData) createEmail(email *models.Email) error {
	ed.mutex.Lock()
	defer ed.mutex.Unlock()

	email.ID = ed.nextID
	ed.nextID++
	ed.emails = append(ed.emails, *email)
	return nil
}

func (ed *EmailData) updateEmail(email *models.Email) error {
	ed.mutex.Lock()
	defer ed.mutex.Unlock()

	for i, existing := range ed.emails {
		if existing.ID == email.ID {
			ed.emails[i] = *email
			return nil
		}
	}
	return errors.New("email not found")
}

func (ed *EmailData) deleteEmail(id int) error {
	ed.mutex.Lock()
	defer ed.mutex.Unlock()

	for i, e := range ed.emails {
		if e.ID == id {
			ed.emails = append(ed.emails[:i], ed.emails[i+1:]...)
			return nil
		}
	}
	return errors.New("email not found")
}
