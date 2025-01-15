package models

import (
	"gorm.io/gorm"
)

// EmailRepository es un repositorio para manejar operaciones de correos electrónicos
type EmailRepository struct {
	DB *gorm.DB
}

// Obtener todos los correos electrónicos
func (r *EmailRepository) GetEmails() ([]Email, error) {
	var emails []Email
	err := r.DB.Find(&emails).Error
	return emails, err
}

// Obtener un correo electrónico por ID
func (r *EmailRepository) GetEmailByID(id uint) (*Email, error) {
	var email Email
	err := r.DB.First(&email, id).Error
	return &email, err
}
