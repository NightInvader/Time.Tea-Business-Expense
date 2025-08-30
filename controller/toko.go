package controller

import (
	"net/http"

	"time.tea/config"
	"time.tea/structs"

	"github.com/gin-gonic/gin"
)

func CreateToko(c *gin.Context) {
	var toko structs.Toko

	if err := c.ShouldBindJSON(&toko); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&toko).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, toko)
}

func GetTokos(c *gin.Context) {
	var tokos []structs.Toko
	config.DB.Find(&tokos)
	c.JSON(http.StatusOK, tokos)
}

func GetTokoByID(c *gin.Context) {
	id := c.Param("id")
	var toko structs.Toko

	if err := config.DB.First(&toko, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "toko not found"})
		return
	}

	c.JSON(http.StatusOK, toko)
}

func UpdateToko(c *gin.Context) {
	var toko structs.Toko
	id := c.Param("id")

	if err := config.DB.First(&toko, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "toko Not found"})
		return
	}

	if err := c.ShouldBindJSON(&toko); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&toko)
	c.JSON(http.StatusOK, toko)
}

func DeleteToko(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&structs.Toko{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Toko Deleted"})
}
