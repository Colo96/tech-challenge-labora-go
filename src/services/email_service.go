package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tech-challenge/src/models"

	"github.com/jinzhu/gorm"
)

const (
	// ZincSearch URL y autenticación
	ZincUser     = "admin"
	ZincPassword = "Complexpass#123"
)

// ProcessAndSendEmails procesa lotes de correos electrónicos y los envía a ZincSearch y PostgreSQL.
func ProcessAndSendEmails(batchChan chan []*models.Email, db *gorm.DB, zincURL, zincAPIKey string) {
	for batch := range batchChan {
		err := sendBatch(batch, db, zincURL, zincAPIKey)
		if err != nil {
			log.Printf("Error al enviar el lote de correos: %v\n", err)
		}
	}
}

// sendBatch envía un lote de correos electrónicos a ZincSearch y PostgreSQL.
func sendBatch(batch []*models.Email, db *gorm.DB, zincURL, zincAPIKey string) error {
	// Enviar a ZincSearch
	err := sendToZincSearch(batch, zincURL, zincAPIKey)
	if err != nil {
		return fmt.Errorf("error al enviar a ZincSearch: %v", err)
	}

	// Guardar en PostgreSQL
	err = saveToPostgreSQL(batch, db)
	if err != nil {
		return fmt.Errorf("error al guardar en PostgreSQL: %v", err)
	}

	return nil
}

// sendToZincSearch envía el lote de correos electrónicos a ZincSearch.
func sendToZincSearch(batch []*models.Email, zincURL, zincAPIKey string) error {
	bulkData := map[string]interface{}{
		"index":   "emails",
		"records": batch,
	}

	data, err := json.Marshal(bulkData)
	if err != nil {
		return fmt.Errorf("error al serializar los datos para ZincSearch: %v", err)
	}

	req, err := http.NewRequest("POST", zincURL, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("error al crear la solicitud para ZincSearch: %v", err)
	}

	req.SetBasicAuth(ZincUser, ZincPassword)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error al enviar la solicitud a ZincSearch: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error del servidor ZincSearch, código HTTP: %d", resp.StatusCode)
	}

	return nil
}

// saveToPostgreSQL guarda un lote de correos electrónicos en la base de datos PostgreSQL.
func saveToPostgreSQL(batch []*models.Email, db *gorm.DB) error {
	for _, email := range batch {
		err := db.Create(email).Error
		if err != nil {
			return fmt.Errorf("error al guardar el correo en PostgreSQL: %v", err)
		}
	}
	return nil
}
