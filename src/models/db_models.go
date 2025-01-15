package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

// InitDB inicializa la base de datos utilizando las variables de entorno.
func InitDB(host, port, user, password, dbName string) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	var err error
	DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	log.Println("Conectado a la base de datos PostgreSQL")
}

// Migrate realiza las migraciones para las tablas de la base de datos
func Migrate() {
	DB.AutoMigrate(&Email{})
}
