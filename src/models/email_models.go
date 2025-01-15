package models

import (
	"gorm.io/gorm"
)

// Estructura de un correo electrónico
type Email struct {
	gorm.Model
	MessageID string `json:"message_id"`
	Date      string `json:"date"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}
