package main

import (
	"log"
	"os"
	"sync"
	"tech-challenge/src/models"
	"tech-challenge/src/routes"
	"tech-challenge/src/services"
	"tech-challenge/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}
}

func main() {
	// Obtener las variables de entorno para la base de datos
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Inicializar la base de datos y realizar la migración de las tablas
	models.InitDB(dbHost, dbPort, dbUser, dbPassword, dbName)
	models.Migrate()

	// Obtener las variables de entorno para ZincSearch
	zincURL := os.Getenv("ZINCSEARCH_URL")
	zincAPIKey := os.Getenv("ZINCSEARCH_API_KEY")

	// Inicializar el servidor y las rutas
	r := gin.Default()
	routes.SetupEmailRoutes(r)

	// Canal para procesar lotes de correos electrónicos
	emailChan := make(chan []*models.Email)
	var wg sync.WaitGroup

	// Exploración del directorio "maildir"
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := utils.ExploreDirectory("maildir", emailChan, &wg)
		if err != nil {
			log.Printf("Error al explorar el directorio: %v\n", err)
		}
	}()

	// Enviar los correos electrónicos a ZincSearch y PostgreSQL
	go services.ProcessAndSendEmails(emailChan, models.DB, zincURL, zincAPIKey)

	// Esperar a que todas las goroutines finalicen
	wg.Wait()

	// Iniciar el servidor HTTP
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
