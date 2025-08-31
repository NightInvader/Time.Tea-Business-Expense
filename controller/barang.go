package controller

import (
	"net/http"

	"time.tea/config"
	"time.tea/structs"

	"github.com/gin-gonic/gin"
)

func CreateBarang(c *gin.Context) {
	var barang structs.Barang

	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, barang)
}

func GetBarangs(c *gin.Context) {
	var barangs []structs.Barang
	config.DB.Preload("Toko").Find(&barangs)
	c.JSON(http.StatusOK, barangs)
}

func GetBarangByID(c *gin.Context) {
	id := c.Param("id")
	var barang structs.Barang

	if err := config.DB.First(&barang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "barang not found"})
		return
	}

	c.JSON(http.StatusOK, barang)
}

func UpdateBarang(c *gin.Context) {
	var barang structs.Barang
	id := c.Param("id")

	if err := config.DB.First(&barang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "barang Not found"})
		return
	}

	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&barang)
	c.JSON(http.StatusOK, barang)
}

func DeleteBarang(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&structs.Barang{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Barang Deleted"})
}
