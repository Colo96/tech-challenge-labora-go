package routes

import (
	"tech-challenge/src/controllers"

	"github.com/gin-gonic/gin"
)

// SetupEmailRoutes configura las rutas de la API de correos electrónicos
func SetupEmailRoutes(router *gin.Engine) {
	emailRoutes := router.Group("/emails")
	{
		emailRoutes.GET("/", controllers.GetEmails)         // Obtener todos los correos electrónicos
		emailRoutes.GET("/:id", controllers.GetEmailByID)   // Obtener un correo electrónico por ID
		emailRoutes.POST("/", controllers.CreateEmail)      // Crear un nuevo correo electrónico
		emailRoutes.PUT("/:id", controllers.UpdateEmail)    // Actualizar un correo electrónico por ID
		emailRoutes.DELETE("/:id", controllers.DeleteEmail) // Eliminar un correo electrónico por ID
	}
}
