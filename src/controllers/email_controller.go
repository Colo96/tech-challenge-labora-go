package controllers

import (
	"net/http"
	"tech-challenge/src/models"

	"github.com/gin-gonic/gin"
)

// Obtener todos los correos electrónicos
func GetEmails(c *gin.Context) {
	var emails []models.Email
	if err := models.DB.Find(&emails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, emails)
}

// Obtener un correo electrónico por ID
func GetEmailByID(c *gin.Context) {
	var email models.Email
	id := c.Param("id")
	if err := models.DB.First(&email, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Correo no encontrado"})
		return
	}
	c.JSON(http.StatusOK, email)
}

// Crear un nuevo correo electrónico
func CreateEmail(c *gin.Context) {
	var email models.Email
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&email).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, email)
}

// Actualizar un correo electrónico por ID
func UpdateEmail(c *gin.Context) {
	var email models.Email
	id := c.Param("id")
	if err := models.DB.First(&email, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Correo no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Save(&email).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, email)
}

// Eliminar un correo electrónico por ID
func DeleteEmail(c *gin.Context) {
	var email models.Email
	id := c.Param("id")
	if err := models.DB.First(&email, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Correo no encontrado"})
		return
	}
	if err := models.DB.Delete(&email).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
