package models

// Email representa un correo electrónico.
type Email struct {
	ID        string `json:"id"`
	MessageID string `json:"message_id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Date      string `json:"date"`
	Body      string `json:"body"`
}
