package main

import (
	"os"

	"time.tea/config"
	"time.tea/controller"
	"time.tea/middleware"
	"time.tea/structs"

	"github.com/gin-gonic/gin"
)

var err error

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&structs.User{}, &structs.Toko{}, &structs.Customer{}, &structs.Barang{}, &structs.Pengeluaran{}, &structs.Pemasukan{})

	router := gin.Default()

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	auth := router.Group("/api", middleware.AuthMiddleware())
	{
		auth.POST("/pengeluaran", controller.CreatePengeluaran)
		auth.GET("/pengeluaran", controller.GetPengeluaran)
		auth.GET("/pengeluaran:id", controller.GetPengeluaranByID)
		auth.GET("/pengeluaran/total/:tanggal", controller.GetTotalByTanggal)
		auth.PUT("/pengeluaran/:id", controller.UpdatePengeluaran)
		auth.DELETE("/pengeluaran/:id", controller.DeletePengeluaran)

		auth.POST("/pemasukan", controller.CreatePemasukan)
		auth.GET("/pemasukan", controller.GetPemasukan)
		auth.GET("/pemasukan:id", controller.GetPemasukanByID)
		auth.GET("/pemasukan/total/:tanggal", controller.GetProfitByTanggal)
		auth.PUT("/pemasukan/:id", controller.UpdatePemasukan)
		auth.DELETE("/pemasukan/:id", controller.DeletePemasukan)

		auth.POST("/toko", controller.CreateToko)
		auth.GET("/toko", controller.GetTokos)
		auth.GET("/toko/:id", controller.GetTokoByID)
		auth.PUT("/toko/:id", controller.UpdateToko)
		auth.DELETE("/toko/:id", controller.DeleteToko)

		auth.POST("/customer", controller.CreateCust)
		auth.GET("/customer", controller.GetCustomers)
		auth.GET("/customer/:id", controller.GetCustByID)
		auth.PUT("/customer/:id", controller.UpdateCust)
		auth.DELETE("/customer/:id", controller.DeleteCust)

		auth.POST("/barang", controller.CreateBarang)
		auth.GET("/barang", controller.GetBarangs)
		auth.GET("/barang/:id", controller.GetBarangByID)
		auth.PUT("/barang/:id", controller.UpdateBarang)
		auth.DELETE("/barang/:id", controller.DeleteBarang)

		auth.GET("/user", controller.GetUsers)
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	router.Run(":" + PORT)
}
