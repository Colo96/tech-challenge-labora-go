package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tech-challenge/src/models"
)

const (
	ZincURL      = "http://localhost:4080/api/_bulkv2" // Endpoint de ZincSearch
	ZincUser     = "admin"                             // Usuario de ZincSearch
	ZincPassword = "Complexpass#123"                   // Contraseña de ZincSearch
	BatchSize    = 1000                                // Tamaño de lote para el envío
)

// ProcessAndSendEmails procesa y envía lotes de correos electrónicos a ZincSearch.
func ProcessAndSendEmails(batchChan chan []*models.Email) {
	for batch := range batchChan {
		err := sendBatch(batch)
		if err != nil {
			log.Printf("Error al enviar el lote: %v\n", err)
		}
	}
}

func sendBatch(emails []*models.Email) error {
	bulkData := map[string]interface{}{
		"index":   "emails",
		"records": emails,
	}

	data, err := json.Marshal(bulkData)
	if err != nil {
		return fmt.Errorf("error al serializar los datos: %v", err)
	}

	req, err := http.NewRequest("POST", ZincURL, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("error al crear la solicitud: %v", err)
	}

	req.SetBasicAuth(ZincUser, ZincPassword)
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
