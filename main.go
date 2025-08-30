package main

import (
	"os"

	"time.tea/config"
	"time.tea/controller"
	"time.tea/structs"

	"github.com/gin-gonic/gin"
)

var err error

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&structs.User{}, &structs.Toko{}, &structs.Customer{}, &structs.Barang{}, &structs.Pengeluaran{}, &structs.Pemasukan{})

	router := gin.Default()

	router.GET("/toko", controller.GetTokos)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	router.Run(":" + PORT)
}
