package utils

import (
	"os"
	"path/filepath"
	"sync"
	"tech-challenge/models"
)

// ExploreDirectory recorre un directorio recursivamente y procesa archivos de correos electrónicos.
func ExploreDirectory(dir string, emailChan chan *models.Email, wg *sync.WaitGroup) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		fullPath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			wg.Add(1)
			go func(path string) {
				defer wg.Done()
				_ = ExploreDirectory(path, emailChan, wg)
			}(fullPath)
		} else {
			email := ParseEmailFile(fullPath)
			if email != nil {
				emailChan <- email
			}
		}
	}
	return nil
}
