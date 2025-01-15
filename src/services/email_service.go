package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tech-challenge/src/models"

	"gorm.io/gorm"
)

const (
	BatchSize = 1000 // Tamaño de lote para el envío
)

// ProcessAndSendEmails procesa y envía lotes de correos electrónicos a ZincSearch.
func ProcessAndSendEmails(batchChan chan []*models.Email, db *gorm.DB, zincURL, zincAPIKey string) {
	for batch := range batchChan {
		err := sendBatch(batch, zincURL, zincAPIKey)
		if err != nil {
			log.Printf("Error al enviar el lote a ZincSearch: %v\n", err)
		}

		// Guardar los correos electrónicos en la base de datos
		for _, email := range batch {
			if err := saveEmailToDB(db, email); err != nil {
				log.Printf("Error al guardar el correo en la base de datos: %v\n", err)
			}
		}
	}
}

func sendBatch(emails []*models.Email, zincURL, zincAPIKey string) error {
	bulkData := map[string]interface{}{
		"index":   "emails",
		"records": emails,
	}

	data, err := json.Marshal(bulkData)
	if err != nil {
		return fmt.Errorf("error al serializar los datos: %v", err)
	}

	req, err := http.NewRequest("POST", zincURL, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("error al crear la solicitud: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+zincAPIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error al enviar la solicitud: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error del servidor, código HTTP: %d", resp.StatusCode)
	}

	return nil
}

// saveEmailToDB guarda el correo electrónico en la base de datos PostgreSQL
func saveEmailToDB(db *gorm.DB, email *models.Email) error {
	if err := db.Create(&email).Error; err != nil {
		return fmt.Errorf("error al guardar el correo: %v", err)
	}
	return nil
}
