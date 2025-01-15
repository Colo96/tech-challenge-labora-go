package utils

import (
	"fmt"
	"os"
	"sync"
	"tech-challenge/src/models"
)

// Explora el directorio "maildir" para encontrar archivos
func ExploreDirectory(dirPath string, emailChan chan *models.Email, wg *sync.WaitGroup) error {
	defer close(emailChan)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("error al leer el directorio: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			wg.Add(1)
			go func(fileName string) {
				defer wg.Done()
				email, err := ParseEmailFile(fileName)
				if err != nil {
					fmt.Printf("Error al parsear el archivo %s: %v\n", fileName, err)
					return
				}
				emailChan <- email
			}(file.Name())
		}
	}
	return nil
}
