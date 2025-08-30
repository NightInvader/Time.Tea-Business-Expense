package structs

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Nama     string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}

type Toko struct {
	ID     uint `gorm:"primaryKey"`
	Nama   string
	Lokasi string
	Barang []Barang `gorm:"foreignKey:ID_Toko"`
}

type Customer struct {
	ID   uint `gorm:"primaryKey"`
	Nama string
}

type Barang struct {
	ID          uint `gorm:"primaryKey"`
	Nama        string
	Harga       int
	ID_Toko     uint
	Toko        []Toko        `gorm:"foreignKey:ID_Toko"`
	Pengeluaran []Pengeluaran `gorm:"foreignKey:ID_Barang"`
}

type Pengeluaran struct {
	ID         uint `gorm:"primaryKey"`
	Tanggal    time.Time
	ID_Barang  uint
	Jumlah     int
	Harga      int
	Total      int
	ID_Toko    uint
	ID_Pendata uint
	Barang     Barang `gorm:"foreignKey:ID_Barang"`
	Toko       Toko   `gorm:"foreignKey:ID_Toko"`
	User       User   `gorm:"foreignKey:ID_Pendata" json:"user_id"`
}

type Pemasukan struct {
	ID          uint `gorm:"primaryKey"`
	Tanggal     time.Time
	Jenis       string
	Ukuran      string
	Terjual     int
	ID_Customer uint
	HargaJual   int
	Modal       int
	Profit      int
	Total       int
	Customer    Customer `gorm:"foreignKey:ID_Customer"`
}
