package controllers

import (
	"net/http"
	"pos-backend/config"
	"pos-backend/models"

	"github.com/gin-gonic/gin"
)

func GetSales(c *gin.Context) {
	var sales []models.Sale
	config.DB.Find(&sales)
	c.JSON(http.StatusOK, sales)
}

func CreateSale(c *gin.Context) {
	var sale models.Sale
	if err := c.BindJSON(&sale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&sale)
	c.JSON(http.StatusCreated, sale)
}
