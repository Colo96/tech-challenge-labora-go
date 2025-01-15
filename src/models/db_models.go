package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB inicializa la conexión a la base de datos PostgreSQL.
func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=password dbname=emaildb port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
}

// Email representa la estructura de un correo electrónico en la base de datos.
type Email struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	MessageID string `json:"message_id" gorm:"uniqueIndex"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Date      string `json:"date"`
	Body      string `json:"body"`
}

// Migrar las tablas.
func Migrate() {
	err := DB.AutoMigrate(&Email{})
	if err != nil {
		log.Fatalf("Error al migrar la base de datos: %v", err)
	}
}
