package services

import (
	"fmt"
	"log"
	"tech-challenge/src/models"

	"gorm.io/gorm"
)

// Procesar y enviar correos electrónicos a ZincSearch y PostgreSQL
func ProcessAndSendEmails(emailChan chan *models.Email, db *gorm.DB) {
	for email := range emailChan {
		// Guardar en PostgreSQL
		err := db.Create(email).Error
		if err != nil {
			log.Printf("Error al guardar el correo: %v", err)
			continue
		}

		// Enviar a ZincSearch (Aquí iría el código para interactuar con ZincSearch)
		err = sendToZincSearch(email)
		if err != nil {
			log.Printf("Error al enviar el correo a ZincSearch: %v", err)
		}
	}
}

func sendToZincSearch(email *models.Email) error {
	// Simulando el envío a ZincSearch (aquí se incluiría la lógica de integración real)
	fmt.Printf("Enviando correo a ZincSearch: %v\n", email.MessageID)
	return nil
}
