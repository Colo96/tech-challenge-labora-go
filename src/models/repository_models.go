package models

type Repository interface {
	getEmails() ([]Email, error)
	getEmailById(id int) (*Email, error)
	createEmail(email *Email) error
	updateEmail(email *Email) error
	deleteEmail(id int) error
}
