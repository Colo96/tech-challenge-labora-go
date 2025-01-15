package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tech-challenge/src/models"
)

// ParseEmailFile lee y parsea un archivo de correo electrónico.
func ParseEmailFile(filePath string) (*models.Email, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	email := &models.Email{}
	scanner := bufio.NewScanner(file)
	var body strings.Builder
	var readingBody bool

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "From: ") {
			email.From = strings.TrimPrefix(line, "From: ")
		} else if strings.HasPrefix(line, "To: ") {
			email.To = strings.TrimPrefix(line, "To: ")
		} else if strings.HasPrefix(line, "Subject: ") {
			email.Subject = strings.TrimPrefix(line, "Subject: ")
		} else if strings.HasPrefix(line, "Date: ") {
			email.Date = strings.TrimPrefix(line, "Date: ")
		}

		if line == "" {
			readingBody = true
		}
		if readingBody {
			body.WriteString(line + "\n")
		}
	}

	if scanner.Err() != nil {
		return nil, fmt.Errorf("error al leer el archivo: %v", scanner.Err())
	}

	email.Body = body.String()
	return email, nil
}
