package database

import (
	"go-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Public Function 은 앞에 대문자로 기록 / 소문자는 프라이빗
func Connect() {
	database, err := gorm.Open(mysql.Open("root:test1234@/go_admin"), &gorm.Config{})

	if err != nil {
		panic("Could not connet to the database")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{})

}
