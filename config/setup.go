package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"golang-pkllaporan/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/pkllaporan_1"))
	if err != nil {
		fmt.Println("Gagal koneksi database:", err)
		return
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Gagal melakukan AutoMigrate:", err)
		return
	}

	DB = db
}