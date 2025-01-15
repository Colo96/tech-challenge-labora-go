package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"tech-challenge/src/models"
)

// ExploreDirectory recorre un directorio recursivamente y procesa archivos de correos electrónicos.
func ExploreDirectory(dir string, emailChan chan []*models.Email, wg *sync.WaitGroup) error {
	defer wg.Done()
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	batch := []*models.Email{}
	batchSize := 100 // Número de correos a enviar por lote

	for _, file := range files {
		fullPath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			wg.Add(1)
			go func(path string) {
				defer wg.Done()
				err := ExploreDirectory(path, emailChan, wg)
				if err != nil {
					fmt.Println("Error al explorar subdirectorio:", err)
				}
			}(fullPath)
		} else {
			email := ParseEmailFile(fullPath)
			if email != nil {
				batch = append(batch, email)
				// Si alcanzamos el tamaño del lote, lo enviamos
				if len(batch) >= batchSize {
					emailChan <- batch
					batch = []*models.Email{} // Limpiar el lote
				}
			}
		}
	}

	// Enviar el último lote si no está vacío
	if len(batch) > 0 {
		emailChan <- batch
	}

	return nil
}
