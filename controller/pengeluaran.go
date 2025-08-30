package controller

import (
	"net/http"
	"time"

	"time.tea/config"
	"time.tea/structs"

	"github.com/gin-gonic/gin"
)

func CreatePengeluaran(c *gin.Context) {
	var entry structs.Pengeluaran

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Preload("Barang").Find(&entry)
	entry.Harga = entry.Barang.Harga

	entry.Total = entry.Jumlah * entry.Harga

	userID := c.MustGet("user_id").(uint)
	entry.ID_Pendata = userID

	if err := config.DB.Create(&entry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, entry)
}

func GetPengeluaran(c *gin.Context) {
	var entries []structs.Pengeluaran
	config.DB.Find(&entries)
	c.JSON(http.StatusOK, entries)
}

func GetPengeluaranByID(c *gin.Context) {
	id := c.Param("id")
	var entry structs.Pengeluaran

	if err := config.DB.First(&entry, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "entry not found"})
		return
	}

	c.JSON(http.StatusOK, entry)
}

func UpdatePengeluaran(c *gin.Context) {
	var entry structs.Pengeluaran
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

func DeletePengeluaran(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&structs.Pengeluaran{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengeluaran Deleted"})
}

func GetTotalByTanggal(c *gin.Context) {
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

	var total int64
	if err := config.DB.Model(&structs.Pengeluaran{}).
		Where("DATE(tanggal) = ?", tanggal).
		Select("SUM(total)").Scan(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tanggal":           tanggalStr,
		"total_pengeluaran": total,
	})
}
