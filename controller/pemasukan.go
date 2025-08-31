package controller

import (
	"net/http"
	"time"

	"time.tea/config"
	"time.tea/structs"

	"github.com/gin-gonic/gin"
)

func CreatePemasukan(c *gin.Context) {
	var entry structs.Pemasukan

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entry.Total = entry.Terjual * entry.HargaJual

	switch entry.Jenis {
	case "Teh Tarik", "teh tarik":
		switch entry.Ukuran {
		case "kecil":
			entry.Profit = (entry.HargaJual - 2000) * entry.Terjual
		case "besar":
			entry.Profit = (entry.HargaJual - 2500) * entry.Terjual
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "ukuran tidak sesuai"})
		}
	case "Lemon Tea", "lemon tea":
		switch entry.Ukuran {
		case "kecil":
			entry.Profit = (entry.HargaJual - 2500) * entry.Terjual
		case "besar":
			entry.Profit = (entry.HargaJual - 3500) * entry.Terjual
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "ukuran tidak sesuai"})
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "minuman tidak tersedia"})
	}

	entry.Modal = entry.Total - entry.Profit

	if err := config.DB.Create(&entry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, entry)
}

func GetPemasukan(c *gin.Context) {
	var entries []structs.Pemasukan
	config.DB.Preload("Customer").Find(&entries)
	c.JSON(http.StatusOK, entries)
}

func GetPemasukanByID(c *gin.Context) {
	id := c.Param("id")
	var entry structs.Pemasukan

	if err := config.DB.First(&entry, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "entry not found"})
		return
	}

	c.JSON(http.StatusOK, entry)
}

func UpdatePemasukan(c *gin.Context) {
	var entry structs.Pemasukan
	id := c.Param("id")

	if err := config.DB.First(&entry, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "entry Not found"})
		return
	}

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&entry)
	c.JSON(http.StatusOK, entry)
}

func DeletePemasukan(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&structs.Pemasukan{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pemasukan Deleted"})
}

func GetProfitByTanggal(c *gin.Context) {
	tanggalStr := c.Query("tanggal") // expect format YYYY-MM-DD
	if tanggalStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tanggal parameter is required"})
		return
	}

	tanggal, err := time.Parse("2006-01-02", tanggalStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, use YYYY-MM-DD"})
		return
	}

	var totalProfit int64
	if err := config.DB.Model(&structs.Pemasukan{}).
		Where("DATE(tanggal) = ?", tanggal).
		Select("SUM(profit)").Scan(&totalProfit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tanggal":      tanggalStr,
		"total_profit": totalProfit,
	})
}
