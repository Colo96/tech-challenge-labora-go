package main

import (
	"fmt"
	"sync"
	"tech-challenge/src/models"
	"tech-challenge/src/services"
	"tech-challenge/src/utils"
)

var (
	wg sync.WaitGroup
)

func main() {
	rootDir := "./maildir"

	// Canal para recolectar correos procesados
	emailChan := make(chan *models.Email)
	// Canal para agrupar y enviar correos
	batchChan := make(chan []*models.Email)

	// Goroutine para procesar y enviar lotes de correos
	go services.ProcessAndSendEmails(batchChan)

	// Goroutine para recolectar correos y enviarlos en lotes
	go func() {
		var batch []*models.Email
		for email := range emailChan {
			batch = append(batch, email)
			if len(batch) >= services.BatchSize {
				batchChan <- batch
				batch = nil
			}
		}
		if len(batch) > 0 {
			batchChan <- batch
		}
		close(batchChan)
	}()

	// Exploración de directorios y procesamiento de correos
	if err := utils.ExploreDirectory(rootDir, emailChan, &wg); err != nil {
		fmt.Println("Error al explorar el directorio:", err)
		return
	}

	wg.Wait()
	close(emailChan)

	fmt.Println("Proceso completado.")
}
