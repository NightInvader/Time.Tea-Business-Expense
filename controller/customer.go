package controller

import (
	"net/http"

	"time.tea/config"
	"time.tea/structs"

	"github.com/gin-gonic/gin"
)

func CreateCust(c *gin.Context) {
	var customer structs.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func GetCustomers(c *gin.Context) {
	var customers []structs.Customer
	config.DB.Find(&customers)
	c.JSON(http.StatusOK, customers)
}

func GetCustByID(c *gin.Context) {
	id := c.Param("id")
	var customer structs.Customer

	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func UpdateCust(c *gin.Context) {
	var customer structs.Customer
	id := c.Param("id")

	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer Not found"})
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&customer)
	c.JSON(http.StatusOK, customer)
}

func DeleteCust(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&structs.Customer{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer Deleted"})
}
