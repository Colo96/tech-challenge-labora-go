package main

import (
	"log"
	"sync"
	"tech-challenge/src/models"
	"tech-challenge/src/routes"
	"tech-challenge/src/services"
	"tech-challenge/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la base de datos y realizar la migración de las tablas
	models.InitDB()
	models.Migrate()

	// Inicializar el servidor y las rutas
	r := gin.Default()
	routes.SetupEmailRoutes(r)

	// Canal para procesar los correos electrónicos de la carpeta "maildir"
	emailChan := make(chan *models.Email)
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
	go services.ProcessAndSendEmails(emailChan, models.DB)

	// Esperar a que todas las goroutines finalicen
	wg.Wait()

	// Iniciar el servidor HTTP
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
