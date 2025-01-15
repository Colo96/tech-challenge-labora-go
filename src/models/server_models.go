package models

import (
	"log"

	"gorm.io/gorm"
)

// Inicializar el servidor y manejar las dependencias
func InitServer(db *gorm.DB) {
	log.Println("Servidor inicializado correctamente")
}
