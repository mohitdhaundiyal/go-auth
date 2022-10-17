package config

import (
	"log"

	"github.com/mohitdhaundiyal/go-auth/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

var urlDSN = "root:22031998@tcp(localhost:3306)/go_test?parseTime=true"

func DatabaseMigration() {
	DB, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	DB.AutoMigrate(&models.User{})
}
