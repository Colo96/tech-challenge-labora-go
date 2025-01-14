package utils

import (
	"bufio"
	"os"
	"strings"
	"tech-challenge/src/models"
)

// ParseEmailFile lee y parsea un archivo de correo electrónico.
func ParseEmailFile(filePath string) *models.Email {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	email := &models.Email{}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Message-ID:") {
			email.MessageID = strings.TrimSpace(strings.TrimPrefix(line, "Message-ID:"))
		} else if strings.HasPrefix(line, "Date:") {
			email.Date = strings.TrimSpace(strings.TrimPrefix(line, "Date:"))
		} else if strings.HasPrefix(line, "From:") {
			email.From = strings.TrimSpace(strings.TrimPrefix(line, "From:"))
		} else if strings.HasPrefix(line, "To:") {
			email.To = strings.TrimSpace(strings.TrimPrefix(line, "To:"))
		} else if strings.HasPrefix(line, "Subject:") {
			email.Subject = strings.TrimSpace(strings.TrimPrefix(line, "Subject:"))
		} else if line == "" {
			break
		}
	}

	body := []string{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		body = append(body, strings.TrimSpace(line))
	}
	email.Body = strings.Join(body, "\n")

	return email
}
